package service

import (
	"fmt"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) TimeUntil() string {
	t := fmt.Sprintf("%d", int(time.Until(time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Now().Location())).Hours())/24)
	return t
}
