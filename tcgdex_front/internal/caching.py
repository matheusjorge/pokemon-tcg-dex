from functools import lru_cache
import io

import requests


@lru_cache(maxsize=1000)
def cached_image(image_url: str):
    res = requests.get(image_url)
    img_byte_arr = io.BytesIO(res.content)
    img_byte_arr.seek(0)

    return img_byte_arr
