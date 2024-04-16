import HuckDjvs
import Neuton
import numpy as np

def f(x):
    return (x[0] - 3)**2 + (x[1] + 1)**2

def main():
    x = HuckDjvs.Hoock(f,np.array([1.0,2.0]))
    print(x)
    x = Neuton.Neuton(f,np.array([1.0,2.0]))
    print(x)


if __name__ == "__main__":
    main()