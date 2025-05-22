package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"
	srv "github.com/a5bbbbb/AITUmoment/core_service/internal/services"
)

// ────────────────────────── MOCKS ──────────────────────────

// RedisCache mock

type redisCacheMock struct{ mock.Mock }

func (m *redisCacheMock) Get(ctx context.Context, id int) (*models.User, error) {
	args := m.Called(ctx, id)
	if u, ok := args.Get(0).(*models.User); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *redisCacheMock) Set(ctx context.Context, u *models.User) error {
	return m.Called(ctx, u).Error(0)
}

func (m *redisCacheMock) SetMany(ctx context.Context, users []models.User) error {
	return nil // not used in this test
}

func (m *redisCacheMock) Delete(ctx context.Context, userID int) error {
	return nil // not used in this test
}

// UserRepository mock (implements every method of real interface so the compiler is happy)

type userRepoMock struct{ mock.Mock }

func (m *userRepoMock) GetUserIdByCred(email string, passwd string) (*int, error) {
	return nil, errors.New("not implemented in test")
}

func (m *userRepoMock) GetUser(id int64) (*models.User, error) {
	args := m.Called(id)
	if u, ok := args.Get(0).(*models.User); ok {
		return u, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *userRepoMock) UpdateUser(user *models.User) (int64, error) { return 0, nil }
func (m *userRepoMock) CreateUser(user *models.User) (int64, error) { return 0, nil }
func (m *userRepoMock) UpdateUserVerified(email string, verified bool) (*int, error) {
	return nil, nil
}
func (m *userRepoMock) GetAllUsers() ([]models.User, error) { return nil, nil }

// EduProgramRepository mock

type eduRepoMock struct{ mock.Mock }

func (m *eduRepoMock) GetEduPrograms() (*[]models.EduProgram, error) {
	args := m.Called()
	if ep, ok := args.Get(0).(*[]models.EduProgram); ok {
		return ep, args.Error(1)
	}
	return nil, args.Error(1)
}

// GroupRepository mock

type groupRepoMock struct{ mock.Mock }

func (m *groupRepoMock) GetGroupByID(id int) (*models.Group, error) {
	args := m.Called(id)
	if g, ok := args.Get(0).(*models.Group); ok {
		return g, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *groupRepoMock) GetGroups(eduProg uint8) (*[]models.Group, error) { return nil, nil }

// Stubs for the deps not exercised in this test

type noopProducer struct{}

func (noopProducer) Push(ctx context.Context, _ models.EmailVerification) error { return nil }

type noopManager struct{}

func (noopManager) Generate(u models.User) (*models.EmailVerification, error) { return nil, nil }
func (noopManager) GetEmail(_ string) (string, error)                         { return "", errors.New("not implemented") }

// ────────────────────────── TESTS ──────────────────────────

func TestUserService_GetFullUserInfo_CacheMissThenHit(t *testing.T) {
	// — Arrange (common objects) —
	userID := 1
	eduProgramID := 1
	groupID := 1

	user := &models.User{
		Id:                 userID,
		Name:               "alice",
		EducationalProgram: uint8(eduProgramID),
		ProgramName:        "Computer Science",
		PublicName:         "Alice",
		Email:              "alice@example.com",
		Group:              uint8(groupID),
		Verified:           true,
	}

	eduPrograms := &[]models.EduProgram{{Id: eduProgramID, Name: "Computer Science"}}
	group := &models.Group{Id: groupID, GroupName: "CS-25-1"}

	// Mocks
	redisM := new(redisCacheMock)
	userRepoM := new(userRepoMock)
	eduRepoM := new(eduRepoMock)
	groupRepoM := new(groupRepoMock)

	// ─── 1️⃣ CACHE MISS PATH ───
	redisM.On("Get", mock.Anything, userID).Return((*models.User)(nil), errors.New("miss"))
	userRepoM.On("GetUser", int64(userID)).Return(user, nil)
	eduRepoM.On("GetEduPrograms").Return(eduPrograms, nil)
	groupRepoM.On("GetGroupByID", int(groupID)).Return(group, nil)
	redisM.On("Set", mock.Anything, user).Return(nil)

	service := srv.NewUserService(
		userRepoM,
		groupRepoM,
		eduRepoM,
		noopProducer{},
		noopManager{},
		redisM,
	)

	// Act
	gotUser, gotPrograms, gotGroup, err := service.GetFullUserInfo(userID)
	// Assert
	require.NoError(t, err)
	require.Equal(t, user.PublicName, gotUser.PublicName)
	require.Equal(t, *eduPrograms, *gotPrograms)
	require.Equal(t, group.Id, gotGroup.Id)

	// Verify expectations for first path
	redisM.AssertExpectations(t)
	userRepoM.AssertExpectations(t)
	eduRepoM.AssertExpectations(t)
	groupRepoM.AssertExpectations(t)

	// ─── 2️⃣ CACHE HIT PATH ───
	// Reset mock call history but keep underlying data
	redisM.ExpectedCalls = nil
	redisM.Calls = nil
	userRepoM.ExpectedCalls = nil
	userRepoM.Calls = nil

	redisM.On("Get", mock.Anything, userID).Return(user, nil) // hit
	eduRepoM.On("GetEduPrograms").Return(eduPrograms, nil)
	groupRepoM.On("GetGroupByID", int(groupID)).Return(group, nil)
	redisM.On("Set", mock.Anything, user).Return(nil)

	gotUser2, _, _, err := service.GetFullUserInfo(userID)
	require.NoError(t, err)
	require.Equal(t, gotUser.PublicName, gotUser2.PublicName)

	redisM.AssertExpectations(t)
	eduRepoM.AssertExpectations(t)
	groupRepoM.AssertExpectations(t)
}
