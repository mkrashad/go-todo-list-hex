package server

import (
	"context"
	"gorm.io/gorm"

	"github.com/mkrashad/go-todo/user/internal"
	"github.com/mkrashad/go-todo/user/pb"
)

type server struct {
	pb.UnimplementedUserServiceServer
	service internal.Service
}

func NewServer(service internal.Service) pb.UserServiceServer {
	return &server{
		service: service,
	}
}

func (s *server) GetByUserNameAndPassword(ctx context.Context, req *pb.GetByUserNameAndPasswordRequest) (*pb.GetByUserNameAndPasswordResponse, error) {
	u, err := s.service.GetByUserNameAndPasword(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	user := toPbUser(u)
	return &pb.GetByUserNameAndPasswordResponse{User: user}, nil
}

func (s *server) GetAllUsers(ctx context.Context, _ *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	u := s.service.GetAllUsers()
	users := make([]*pb.User, 0, len(u))
	for _, user := range u {
		users = append(users, toPbUser(user))
	}

	return &pb.GetAllUsersResponse{
		Users: users,
	}, nil

}

func (s *server) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	t, err := s.service.GetUserById(uint64(req.Id))
	if err != nil {
		return nil, err
	}

	User := toPbUser(t)

	return &pb.GetUserByIdResponse{User: User}, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	User := internal.User{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
	}
	t, err := s.service.CreateUser(User)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{User: toPbUser(t)}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	User := internal.User{
		Model:     gorm.Model{ID: uint(req.Id)},
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
	}
	t, err := s.service.UpdateUserById(uint64(req.Id), User)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{User: toPbUser(t)}, nil
}

func (s *server) DeleteUserById(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.service.DeleteUserById(uint64(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{}, nil
}

func toPbUser(u internal.User) *pb.User {
	return &pb.User{
		Firstname: u.FirstName,
		Lastname:  u.LastName,
		Email:     u.Email,
		Username: u.Username,
		Password: u.Password,
	}
}
