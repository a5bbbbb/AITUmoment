package services

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"
)

type UserService struct {
	userRepo                 UserRepository
	groupRepo                GroupRepository
	eduProgramRepo           EduProgramRepository
	emailVerifProducer       EmailVerificationProducer
	emailVerificationManager EmailVerificationManager
	redisCache               RedisCache
}

func NewUserService(
	userRepo UserRepository,
	groupRepo GroupRepository,
	eduProgramRepo EduProgramRepository,
	emailVerifProducer EmailVerificationProducer,
	verificationLinkGenerator EmailVerificationManager,
	redisCache RedisCache,
) *UserService {
	logger.GetLogger().Trace("Creating user service")
	return &UserService{
		userRepo:                 userRepo,
		groupRepo:                groupRepo,
		eduProgramRepo:           eduProgramRepo,
		emailVerifProducer:       emailVerifProducer,
		emailVerificationManager: verificationLinkGenerator,
		redisCache:               redisCache,
	}
}

func (us *UserService) Authorize(email string, passwd string) (*models.User, error) {
	userID, err := us.userRepo.GetUserIdByCred(email, passwd)

	if err != nil {
		logger.GetLogger().Errorf("error during authentication for user with email=%v err: %v", email, err.Error())
		return nil, err
	}

	return us.userRepo.GetUser(int64(*userID))

}

func (us *UserService) UpdateUser(user *models.User) (*models.User, error) {

	var err error
	id, err := us.userRepo.UpdateUser(user)

	if err != nil {
		logger.GetLogger().Errorf("Error during direct update %v", err.Error())
		return nil, err
	}

	logger.GetLogger().Infof("Here is userID %v", id)
	updatedUser, err := us.userRepo.GetUser(id)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting new user %v", err.Error())
		return nil, err
	}

	err = us.redisCache.Set(context.TODO(), updatedUser)
	if err != nil {
		logger.GetLogger().Error("UserService:UpdateUser could not set user cache")
	}

	return updatedUser, nil
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	var err error
	id, err := us.userRepo.CreateUser(user)

	if err != nil {
		logger.GetLogger().Error("error during creation of user: ", err.Error())
		return nil, err
	}

	verification, err := us.emailVerificationManager.Generate(*user)
	if err != nil {
		logger.GetLogger().Errorln("error generating verification:", err.Error())
		return nil, err
	}

	err = us.emailVerifProducer.Push(context.TODO(), *verification)
	if err != nil {
		logger.GetLogger().Errorln("error pushing verification to message queue:", err.Error())
	}

	createdUser, err := us.userRepo.GetUser(id)

	if err != nil {
		logger.GetLogger().Error("cannot get user with id ", id, " : ", err.Error())
		return nil, err
	}

	err = us.redisCache.Set(context.TODO(), createdUser)
	if err != nil {
		logger.GetLogger().Error("UserService:UpdateUser could not set user cache")
	}

	return createdUser, nil
}

func (us *UserService) VerifyEmail(token string) (*models.User, error) {
	email, err := us.emailVerificationManager.GetEmail(token)
	if err != nil {
		logger.GetLogger().Error("cannot get email from invalid verification token: ", err.Error())
		return &models.User{}, err
	}

	id, err := us.userRepo.UpdateUserVerified(email, true)
	if err != nil || id == nil {
		logger.GetLogger().Error("cannot update user with email", email, " : ", err.Error())
		return &models.User{}, err
	}

	updatedUser, err := us.userRepo.GetUser(int64(*id))
	if err != nil {
		logger.GetLogger().Error("cannot get user with id ", id, " : ", err.Error())
		return nil, err
	}

	err = us.redisCache.Set(context.TODO(), updatedUser)
	if err != nil {
		logger.GetLogger().Error("UserService:VerifyEmail could not set user cache")
	}

	return updatedUser, nil
}

func (us *UserService) GetFullUserInfo(userID int) (*models.User, *[]models.EduProgram, *models.Group, error) {
	user, err := us.redisCache.Get(context.TODO(), userID)

	if err != nil {
		logger.GetLogger().Info("UserService:GetFullUser cache miss (((")

		user, err = us.userRepo.GetUser(int64(userID))

	} else {
		logger.GetLogger().Info("UserService:GetFullUser cache hit  !!!")
	}

	if err != nil {
		logger.GetLogger().Error("cannot get user with id ", userID, " : ", err.Error())
		return nil, nil, nil, err
	}

	programs, err := us.eduProgramRepo.GetEduPrograms()

	if err != nil {
		logger.GetLogger().Error("cannot get edu programs", err.Error())
		return nil, nil, nil, err
	}

	group, err := us.groupRepo.GetGroupByID(int(user.Group))

	if err != nil {
		logger.GetLogger().Error("cannot get group with id ", user.Group, " : ", err.Error())
		return nil, nil, nil, err
	}

	err = us.redisCache.Set(context.TODO(), user)
	if err != nil {
		logger.GetLogger().Error("UserService:UpdateUser could not set user cache")
	}

	return user, programs, group, nil

}
