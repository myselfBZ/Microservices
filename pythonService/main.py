from fastapi import FastAPI
from pydantic import BaseModel
from redis.redis import *
from pythonService.models import *
app = FastAPI()

from pydantic import BaseModel





@app.post("/tasks")
async def create_task(task:Task):
    try: 
        redis_create_task(task)
        return {"message":"created"}
    except Exception as e:
        return {"err":e}
