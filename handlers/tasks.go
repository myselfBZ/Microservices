package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/myselfBZ/mic-services/redis"
	"github.com/myselfBZ/mic-services/services"
)


func NewTaskHandler(serv services.Service) *TaskHandler{
    return &TaskHandler{
        serv: serv,
    }
}

type TaskHandler struct{
    serv services.Service
    redis redis.TaskService
}



func (h *TaskHandler) HandleGetAll(w http.ResponseWriter, r *http.Request){
    tasks, _ := h.serv.GetTasks(context.TODO())
    json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) HandleTaskGet(w http.ResponseWriter, r *http.Request){
    id := r.PathValue("id")
    validatedId, err := strconv.Atoi(id)
    if err != nil{
        w.WriteHeader(http.StatusBadRequest)
        return 
    }
    task, err := h.redis.GetTask(context.TODO(), validatedId)
    if err != nil{
        w.WriteHeader(http.StatusNotFound)
        return 
    }
    json.NewEncoder(w).Encode(task)
}

