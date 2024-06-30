from pydantic import BaseModel 

class Priority:
    lvl:str 

class Task(BaseModel):
    id:int
    text:str 
    title:str 
    priority:Priority
    class Config:
        arbitrary_types_allowed = True 
