from functools import cached_property
import timm


class ModelsLoader:
    def __init__(
        self,
    ):
        pass

    def load_model(self, model_name: str | None):
        match model_name:
            case "mobilenetv3_large_100":
                return self.mobilenetv3_large_100
            case "levit_256.fb_dist_in1k":
                return self.levit_256
            case _:
                return self.mobilenetv3_large_100

    @cached_property
    def mobilenetv3_large_100(self):
        return timm.create_model(
            "mobilenetv3_large_100", pretrained=True, num_classes=0
        )

    @cached_property
    def levit_256(self):
        return timm.create_model(
            "levit_256.fb_dist_in1k", pretrained=True, num_classes=0
        )


models_loader = ModelsLoader()
