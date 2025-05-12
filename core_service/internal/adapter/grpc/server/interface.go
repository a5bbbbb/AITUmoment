package server

import (
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/grpc/server/frontend"
)

type UserService interface {
	frontend.UserService
}

type ThreadService interface {
	frontend.ThreadService
}

type GroupService interface {
	frontend.GroupService
}

type EduProgramService interface {
	frontend.EduProgramService
}
