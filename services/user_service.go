package services

import (
	"aitu-moment/db/repository"
	"aitu-moment/logger"
	"aitu-moment/models"
)

type UserService struct{
    repo *repository.UserRepository
}

func NewUserService()(*UserService){
    return &UserService{repo: repository.NewUserRepository()}
}


func (service *UserService) Authorize(email string, passwd string) ( *models.User, error ) {
    userID, err := service.repo.GetUserIdByCred(email, passwd)

    if err != nil {
        return nil,err
    }

    return service.repo.GetUser(int64(*userID))

}

func(s *UserService) UpdateUser(user *models.User)(*models.User,error){

    var err error
    id, err := s.repo.UpdateUser(user)

    if err != nil {
        logger.GetLogger().Errorf("Error during direct update %v", err.Error())
        return nil,err
    }

    logger.GetLogger().Infof("Here is userID %v", id)
    updatedUser,err := s.repo.GetUser(id)

    if err != nil {
        logger.GetLogger().Errorf("Error during getting new user %v", err.Error())
        return nil,err
    }

    return updatedUser,nil
}

func (s *UserService) VerifyUser(email string) (error){
    
    return s.repo.VerifyUser(email)

}


func(service *UserService) CreateUser(user *models.User)(*models.User,error){
    var err error
    id, err := service.repo.CreateUser(user)

    if err != nil {
        return nil,err
    }

    createdUser,err := service.repo.GetUser(id)

    if err != nil {
        return nil,err
    }

    return createdUser,nil
}


func(s *UserService) GetFullUserInfo(userID int) (*models.User,*[]models.EduProgram,*models.Group,error){
    user,err := s.repo.GetUser(int64(userID))

    if err != nil {
        return nil,nil,nil,err
    }

    programs,err := NewEduService().GetPrograms()

    if err != nil {
        return nil,nil,nil,err
    }

    group,err := NewGroupService().GetGroup(int(user.Group))

    if err != nil {
        return nil,nil,nil,err
    }

    return user,&programs,group, nil


}




