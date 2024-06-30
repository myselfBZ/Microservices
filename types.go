package main 

type Task struct{
    Text string `json:"text"`
    Title string  `json:"title"`
    Priority Priority   `json:"priority"`
}

type Priority struct{
    Lvl string 
}

