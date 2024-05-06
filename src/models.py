from sqlalchemy import Column, Integer, String, DateTime

from .database import Base

""" Global sqlalchemy db models"""

class Feed(Base):
    __tablename__ = 'feeds'

    id = Column(Integer, primary_key=True)
    name = Column(String)
    url = Column(String)
    last_updated = Column(DateTime)
