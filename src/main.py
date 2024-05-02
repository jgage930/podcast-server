from fastapi import FastAPI

from .database import Base, engine, get_db


# create db
Base.metadata.create_all(bind=engine)

app = FastAPI()


@app.get('/')
def index():
    return {
        'msg': 'podcast-server v0.0.1'
    }

