import random
from shared import Model


class LinearGenerator:
    model: Model
    mu_noise: float = 0
    sigma_noise: float = 1.
