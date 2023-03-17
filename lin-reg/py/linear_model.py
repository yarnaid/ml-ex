from shared import Model
from shared import Matrix, Vector
import random


class LinearModel(Model):
    coefs: Matrix
    b: Vector
    noise_fn: callable

    def __init__(self) -> None:
        super().__init__()
