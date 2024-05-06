from fastapi import APIRouter

from ..database import get_db
from .. import models
from .schema import FeedCreate, Feed


feed_router = APIRouter(prefix='/feed')


@feed_router.get('')
def get_all_feeds():
    pass


@feed_router.get('/{id}')
def get_feed_by_id(id: int):
    pass


@feed_router.post('')
def create_feed(payload: FeedCreate):
    pass


@feed_router.delete('/{id}')
def delete_feed(id: int):
    pass
