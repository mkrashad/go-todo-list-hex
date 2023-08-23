package server

import (
	"context"

	task "github.com/mkrashad/go-todo/task/internal"
	"github.com/mkrashad/go-todo/task/pb"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedTaskServiceServer
	service task.Service
}

func NewServer(service task.Service) pb.TaskServiceServer {
	return &server{
		service: service,
	}
}

func (s *server) GetAllTasks(ctx context.Context, _ *pb.GetAllTasksRequest) (*pb.GetAllTasksResponse, error) {
	t := s.service.GetAllTasks()
	tasks := make([]*pb.Task, 0, len(t))
	for _, task := range t {
		tasks = append(tasks, toPbTask(task))
	}

	return &pb.GetAllTasksResponse{
		Tasks: tasks,
	}, nil

}

func (s *server) GetTaskById(ctx context.Context, req *pb.GetTaskByIdRequest) (*pb.GetTaskByIdResponse, error) {
	t, err := s.service.GetTaskById(uint64(req.Id))
	if err != nil {
		return nil, err
	}

	task := toPbTask(t)

	return &pb.GetTaskByIdResponse{Task: task}, nil
}

func (s *server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	task := task.Task{
		TaskName:  req.TaskName,
		Completed: &req.Completed,
		UserID:    uint64(req.UserId),
	}
	t, err := s.service.CreateTask(task)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTaskResponse{Task: toPbTask(t)}, nil
}

func (s *server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	task := task.Task{
		Model:     gorm.Model{ID: uint(req.Id)},
		TaskName:  req.TaskName,
		Completed: &req.Completed,
		UserID:    uint64(req.UserId),
	}
	t, err := s.service.UpdateTaskById(uint64(req.Id), task)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateTaskResponse{Task: toPbTask(t)}, nil
}

func (s *server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := s.service.DeleteTaskById(uint64(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTaskResponse{}, nil
}

func toPbTask(t task.Task) *pb.Task {
	return &pb.Task{
		Id:        int64(t.ID),
		TaskName:  t.TaskName,
		Completed: *t.Completed,
		UserId:    int64(t.UserID),
	}
}
