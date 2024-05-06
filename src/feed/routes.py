from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session
from datetime import datetime

from ..database import get_db
from .. import models
from .schema import FeedCreate, Feed


feed_router = APIRouter(prefix='/feed')


@feed_router.get('', response_model=list[Feed])
def get_all_feeds(db: Session = Depends(get_db)):
    db_feeds = db.query(models.Feed).all()

    return [Feed.from_orm(feed) for feed in db_feeds]


@feed_router.get('/{id}')
def get_feed_by_id(id: int, db: Session = Depends(get_db)):
    pass


@feed_router.post('', response_model=FeedCreate)
def create_feed(payload: FeedCreate, db: Session = Depends(get_db)):
    db_feed = models.Feed(
        name=payload.name,
        url=str(payload.url),
        last_updated=datetime.now()
    )
    db.add(db_feed)
    db.commit()

    return payload


@feed_router.delete('/{id}')
def delete_feed(id: int, db: Session = Depends(get_db)):
    pass
