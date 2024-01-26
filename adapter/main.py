from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def get_mean():
    return {"hello":"World"}