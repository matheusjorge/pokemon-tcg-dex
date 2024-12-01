from enum import Enum
from typing import List

from pydantic import BaseModel


class AvailableModels(str, Enum):
    mobilenetv3_large_100 = "mobilenetv3_large_100"
    levit_256 = "levit_256.fb_dist_in1k"


class AvailableModelsResponse(BaseModel):
    available_models: List[AvailableModels]


class EmbedImageRequest(BaseModel):
    filenames: str
    model: AvailableModels = AvailableModels.levit_256


class EmbedImageResponse(BaseModel):
    embeddings: List[List[float]]


class ComputeSimilarityRequest(BaseModel):
    base: str
    target: str
    model: AvailableModels = AvailableModels.levit_256


class ComputeSimilarityResponse(BaseModel):
    similarity: float
