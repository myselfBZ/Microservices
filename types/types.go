package types

type Task struct{
    ID int
    Text string `json:"text"`
    Title string  `json:"title"`
    Priority Priority   `json:"priority"`
}

type Priority struct{
    Lvl string 
}

