package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)


type TaskHandler struct{
    serv Service

}

func NewTaskHandler(serv Service) *TaskHandler{
    return &TaskHandler{
        serv: serv,
    }
}


func main(){
    srvc := NewTaskService()
    srvc = NewLoggingService(srvc)
    mux := http.NewServeMux()
    h := NewTaskHandler(srvc) 
    mux.HandleFunc("/", h.HandleTaskGet)
    RunServer(":8080", mux)
}


func RunServer(port string, mux *http.ServeMux){
    log.Fatal(http.ListenAndServe(port, mux))

}


func (h *TaskHandler) HandleTaskGet(w http.ResponseWriter, r *http.Request){
    task, _ := h.serv.GetTask(context.TODO())
    json.NewEncoder(w).Encode(task)
}

