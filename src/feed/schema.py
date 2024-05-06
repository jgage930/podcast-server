from pydantic import BaseModel, AnyUrl, Field
from datetime import datetime


class FeedCreate(BaseModel):
    name: str
    url: AnyUrl


class Feed(FeedCreate):
    id: int
    last_updated: datetime = Field(default_factory=datetime.now)

    class Config:
        orm_mode = True
        from_attributes = True

