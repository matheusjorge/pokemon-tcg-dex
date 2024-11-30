from enum import Enum
from typing import List

from pydantic import BaseModel


class AvailableModels(str, Enum):
    mobilenetv3_large_100 = "mobilenetv3_large_100"


class AvailableModelsResponse(BaseModel):
    available_models: List[AvailableModels]


class EmbedImageRequest(BaseModel):
    filenames: str
    model: AvailableModels = AvailableModels.mobilenetv3_large_100


class EmbedImageResponse(BaseModel):
    embeddings: List[List[float]]
