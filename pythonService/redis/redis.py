import json
import redis 
from ..models import Task

r = redis.Redis(host="localhost", port=6379, db=0)


def serializer(task:Task):
    serializedTask = {
        "id":task.id,
        "priority":task.priority,
        "text":task.text,
        "title":task.title
    }
    return json.dumps(serializedTask)

def redis_create_task(task:Task):
    serialized = serializer(task)
    r.lpush("todos", serialized)
