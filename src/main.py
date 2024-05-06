from fastapi import FastAPI

from .database import Base, engine, get_db
from .feed.routes import feed_router


# create db
Base.metadata.create_all(bind=engine)

app = FastAPI()

app.include_router(feed_router)


@app.get('/')
def index():
    return {
        'msg': 'podcast-server v0.0.1'
    }

