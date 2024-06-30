package main

import (
	"context"
	"log"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}

}

func (s *LoggingService) GetTask(ctx context.Context) (task *Task, err error){
    defer func(start time.Time) {
        log.Printf("task=%v err=%s took=%v", task, err, time.Since(start))
    }(time.Now())
    return s.next.GetTask(ctx)
}
