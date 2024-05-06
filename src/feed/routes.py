from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from datetime import datetime
from typing import Optional
from pydantic import AnyUrl

from ..database import get_db
from .. import models
from .schema import FeedCreate, Feed


def get_feed_by_id(id: int, db: Session) -> Optional[Feed]:
    db_feed = db\
                .query(models.Feed)\
                .filter(models.Feed.id == id)\
                .first()

    return db_feed


def valid_feed_url(url: AnyUrl, db: Session) -> bool:
    db_feed = db\
                .query(models.Feed)\
                .filter(models.Feed.url == str(url))\
                .first()

    if db_feed:
        return False

    return True


feed_router = APIRouter(prefix='/feed')


@feed_router.get('', response_model=list[Feed])
def all_feeds(db: Session = Depends(get_db)):
    db_feeds = db.query(models.Feed).all()

    return [Feed.from_orm(feed) for feed in db_feeds]


@feed_router.get('/{id}', response_model=Feed)
def feed_by_id(id: int, db: Session = Depends(get_db)):
    db_feed = get_feed_by_id(id, db)

    if not db_feed:
        raise HTTPException(status_code=404, detail='feed not found')

    return Feed.from_orm(db_feed)


@feed_router.post('', response_model=FeedCreate)
def create_feed(payload: FeedCreate, db: Session = Depends(get_db)):
    if not valid_feed_url(payload.url, db):
        raise HTTPException(
            status_code=409, 
            detail=f'a feed with url {payload.url} already exists.'
        )

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
