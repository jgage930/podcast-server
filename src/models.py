from sqlalchemy import Boolean, Column, ForeignKey, Integer, String

from .database import Base

""" Global sqlalchemy db models"""

class Feed(Base):
    __tablename__ = 'feeds'

    id = Column(Integer, primary_key=True)
    name = Column(String)
    url = Column(String)
