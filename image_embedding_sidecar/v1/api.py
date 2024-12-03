import asyncio
import json
import time

from fastapi import APIRouter, HTTPException, Response
from timm.data import resolve_data_config
from timm.data.transforms_factory import create_transform
import torch


from image_embedding_sidecar.v1.contracts import (
    ComputeSimilarityRequest,
    ComputeSimilarityResponse,
    EmbedImageRequest,
    EmbedImageResponse,
    AvailableModels,
    AvailableModelsResponse,
)
from image_embedding_sidecar.v1.models import models_loader
from image_embedding_sidecar.utils import load_image
from image_embedding_sidecar.logger import logger

router = APIRouter(prefix="/v1")


@router.get("/available_models")
async def available_models() -> AvailableModelsResponse:
    return AvailableModelsResponse(available_models=[m for m in AvailableModels])


@router.post("/images/embeddings")
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
    logger.debug(
        "Finished embeddings execution",
        total_time=func_end - func_start,
        load_model_time=download_start - func_start,
        download_time=tensors_start - download_start,
        transform_time=embedding_start - tensors_start,
        embedding_time=func_end - embedding_start,
    )

    return Response(
        content=json.dumps({"embeddings": embedding.tolist()}),
        media_type="application/json",
    )


@router.post("/images/similarity")
async def compute_similarity(
    request: ComputeSimilarityRequest,
) -> ComputeSimilarityResponse:
    func_start = time.perf_counter()
    model = models_loader.load_model(request.model)
    model.eval()
    if model is None:
        raise HTTPException(status_code=406, detail="model not available")

    config = resolve_data_config({}, model=model)
    transform = create_transform(**config)

    download_start = time.perf_counter()
    tasks = [load_image(f) for f in [request.base, request.target]]

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
    # Compute dot product
    dot_product = torch.dot(embedding[0], embedding[1])

    # Compute magnitudes (norms)
    norm1 = torch.norm(embedding[0])
    norm2 = torch.norm(embedding[1])
    print(embedding)

    # Compute cosine similarity
    cos_sim = dot_product / (norm1 * norm2)
    return ComputeSimilarityResponse(similarity=cos_sim.item())
