from array import array
import abc
import random
from functools import lru_cache

def random_gaussian(mu: float, sigma: float) -> float:
    return random.gauss(mu, sigma)


class Matrix:
    data: array


class Vector:
    data: array


class Model(abc.ABC):
    @abc.abstractmethod
    def execute(self, x: Vector) -> Vector:
        raise NotImplementedError


class Leaner:
    training_set: list
    validation_set: list
    test_set: list
    model: Model
    
    def fit(self):
        pass
