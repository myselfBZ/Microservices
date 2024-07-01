from fastapi import FastAPI
from pydantic import BaseModel
import redis
import json

from pydantic import BaseModel

app = FastAPI()

r = redis.Redis(host="localhost", port=6379, db=0)



class Task(BaseModel):
    id:int
    text:str 
    title:str 
    class Config:
        arbitrary_types_allowed = True 

def serializer(task:Task):
    serializedTask = {
        "id":task.id,
        "text":task.text,
        "title":task.title
    }
    return json.dumps(serializedTask)

def redis_create_task(task:Task):
    serialized = serializer(task)
    r.lpush("todos", serialized)




@app.post("/tasks")
async def create_task(task:Task):
    try: 
        redis_create_task(task)
        return {"message":"created"}
    except Exception as e:
        return {"err":e}
