import requests 

task = {
    
    "id":1,
    "priority":{
        "lvl":"high"
    },
    "text":"something",
    "title":"do something"
}

response = requests.get("http://localhost:8080/tasks")

print(response.json())
