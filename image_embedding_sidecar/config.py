import os


class Config:
    ENV = os.getenv("ENV", "dev")
