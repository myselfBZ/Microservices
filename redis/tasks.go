package redis

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/myselfBZ/mic-services/services"
	"github.com/myselfBZ/mic-services/types"
)


type TaskService struct{
    client *redis.Client
} 




func NewTaskService(client *redis.Client) services.Service {
    return &TaskService{
        client: client,
    }
}

func (s *TaskService) GetTasks(ctx context.Context) ([]types.Task, error)  {
    var tasks []types.Task
    var stop int64 = -1
    var start int64 = 0
    jsonTasks, err := s.client.LRange(ctx, "todos", start, stop).Result()
    if err != nil {
        return nil, err 
    }
    for _, jsonTask := range jsonTasks{
        var task types.Task 
        err := json.Unmarshal([]byte(jsonTask), &task)
        if err != nil{
            return nil, err 
        }
        tasks = append(tasks, task)

    }
    return tasks, nil 
}

func (s *TaskService) GetTask(ctx context.Context, id int) (*types.Task, error){
    var stop int64 = -1
    var start int64 = 0
    jsonTasks, err := s.client.LRange(ctx, "todos", start, stop).Result()
    if err != nil {
        return nil, err 
    }
    for _, jsonTask := range jsonTasks{
        var task types.Task 
        err := json.Unmarshal([]byte(jsonTask), &task)
        if err != nil{
            return nil, err 
        }
        if task.ID == id{
            return &task, nil 
        } 
    }
    return nil, errors.New("Task not found")
}






