import numpy as np 
import scipy as sp
from scipy.optimize import linprog

def method(x,
        cel : str="max",
        c:tuple=[],
        A:tuple=[[]],
        b:tuple=[]):
    c = np.array(c)
    if cel == "max":
        c *= -1
    print(linprog(c, A_ub=A, b_ub = b, bounds=x, method="simplex"))