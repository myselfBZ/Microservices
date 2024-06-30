package main

import (
	"log"
	"net/http"

	"github.com/myselfBZ/mic-services/handlers"
	"github.com/myselfBZ/mic-services/redis"
)





func main(){
    redisClient := redis.RedisInit()
    
    srvc := redis.NewTaskService(redisClient)
    
    srvc = NewLoggingService(srvc)
    mux := http.NewServeMux()
    h := handlers.NewTaskHandler(srvc) 
    mux.HandleFunc("/tasks/{id}", h.HandleTaskGet)
    mux.HandleFunc("/tasks", h.HandleGetAll)
    RunServer(":8080", mux)
}


func RunServer(port string, mux *http.ServeMux){
    log.Fatal(http.ListenAndServe(port, mux))

}

