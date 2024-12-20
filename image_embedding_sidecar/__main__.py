from fastapi import FastAPI
import uvicorn

from image_embedding_sidecar.v1.api import router

app = FastAPI(name="ImageEmbeddingSidecar")
app.include_router(router)


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
