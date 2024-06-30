package main

import "context"

type Service interface {
	GetTask(context.Context) (*Task, error)
}

type TaskService struct{} 


func NewTaskService() Service {
    return &TaskService{}
}

func (s *TaskService) GetTask(ctx context.Context) (*Task, error)  {
    return &Task{
        Text: "Do something",
        Title: "Something",
        Priority: Priority{
            Lvl: "High",
        },
    }, nil
}
