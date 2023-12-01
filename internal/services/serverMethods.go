package services

import (
	"context"
	"errors"
	"time"

	"github.com/dolginaa/goProject/internal/database/repository"
	pb "github.com/dolginaa/goProject/pkg/schedule"
)

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	if req == nil {
		return nil, errors.New("received nil request")
	}

	err := s.Repo.Insert(ctx, &repository.Activity{
		Description: req.Description,
		Deadline:    req.Deadline,
		Priority:    int(req.Priority),
		DateTime:    time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{}, nil
}

func (s *Server) GetDaySchedule(ctx context.Context, req *pb.GetDayScheduleRequest) (*pb.GetDayScheduleResponse, error) {
	rows, err := s.Repo.SelectMany(ctx, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, err
	}

	var result []*pb.Task
	for _, activity := range *rows {
		if err != nil {
			return nil, err
		}
		result = append(result, &pb.Task{Description: activity.Description, Priority: int32(activity.Priority), Deadline: activity.Deadline})
	}

	return &pb.GetDayScheduleResponse{Task: result}, nil
}

func (s *Server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	row, err := s.Repo.SelectOne(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetTaskResponse{Task: &pb.Task{Description: row.Description, Priority: int32(row.Priority), Deadline: row.Deadline}}, nil
}
