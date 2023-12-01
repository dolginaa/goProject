package services

import (
	"github.com/dolginaa/goProject/internal/database/repository"
	pb "github.com/dolginaa/goProject/pkg/schedule"
)

type Server struct {
	Repo repository.ActivitiesRepo

	pb.UnimplementedScheduleServer
}
