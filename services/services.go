package services

import( 
    "github.com/myselfBZ/mic-services/types"
    "context"
)


type Service interface {
	GetTasks(context.Context) ([]types.Task, error)
    GetTask(context.Context) (*types.Task, error)
}

