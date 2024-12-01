import asyncio
import time

from fastapi import APIRouter, HTTPException
from timm.data import resolve_data_config
from timm.data.transforms_factory import create_transform
import torch


from image_embedding_sidecar.v1.contracts import (
    EmbedImageRequest,
    EmbedImageResponse,
    AvailableModels,
    AvailableModelsResponse,
)
from image_embedding_sidecar.v1.models import models_loader
from image_embedding_sidecar.utils import load_image

router = APIRouter(prefix="/v1")


@router.get("/available_models")
async def available_models() -> AvailableModelsResponse:
    return AvailableModelsResponse(available_models=[m for m in AvailableModels])


@router.post("/embed_images")
async def embed_image(request: EmbedImageRequest) -> EmbedImageResponse:
    func_start = time.perf_counter()
    model = models_loader.load_model(request.model)
    model.eval()
    if model is None:
        raise HTTPException(status_code=406, detail="model not available")

    config = resolve_data_config({}, model=model)
    transform = create_transform(**config)

    download_start = time.perf_counter()
    tasks = [load_image(f) for f in request.filenames.split(",")]

    imgs = await asyncio.gather(*tasks)

    tensors_start = time.perf_counter()
    tensors = [transform(img) for img in imgs]
    embedding_start = time.perf_counter()
    embedding = model(torch.stack(tensors))

    func_end = time.perf_counter()
    print(f"Total time: {func_end - func_start}")
    print(f"Load Model time: {download_start - func_start}")
    print(f"Download images time: {tensors_start - download_start}")
    print(f"Transform time: {embedding_start - tensors_start}")
    print(f"Embedding time: {func_end - embedding_start}")

    return EmbedImageResponse(embeddings=embedding.tolist())
