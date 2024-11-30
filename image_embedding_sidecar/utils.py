import os
import urllib.request
from PIL import Image


async def download_image(url: str, filename: str):
    req = urllib.request.Request(url)
    req.add_header("User-Agent", "Mozilla/5.0")
    with urllib.request.urlopen(req) as response:
        with open(filename, "wb") as file:
            file.write(response.read())


async def load_image(filename: str):
    path = filename
    image_tmp_dir = "tmp/images"
    if filename.startswith("http"):
        if not os.path.exists(image_tmp_dir):
            os.makedirs(image_tmp_dir, exist_ok=True)
        path = f"{image_tmp_dir}/{filename.rsplit('/')[-1]}"
        await download_image(filename, path)

    return Image.open(path).convert("RGB")
