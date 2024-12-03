import logging
import structlog

from image_embedding_sidecar.config import Config

level = "DEBUG" if Config.ENV == "dev" else "INFO"
LOG_LEVEL = getattr(logging, level)

structlog.configure(
    wrapper_class=structlog.make_filtering_bound_logger(LOG_LEVEL),
    processors=[
        structlog.processors.add_log_level,
        structlog.processors.TimeStamper(fmt="iso"),
        structlog.processors.JSONRenderer(),
    ],
)

logger = structlog.get_logger()
