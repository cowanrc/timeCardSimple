package timecardsvc

import (
	"context"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/timecard"
)

var _ timecard.Service = &Service{}

type Service struct {
	timecardRepo timecard.Repo
}

func New(
	timecardRepo timecard.Repo,
) *Service {
	return &Service{
		timecardRepo: timecardRepo,
	}
}

func (s *Service) CreateTimecard(ctx context.Context, employeeID id.ID) (*timecard.Timecard, error) {
	return nil, nil
}
