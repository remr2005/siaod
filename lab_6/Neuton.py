import numpy as np
from scipy.optimize import approx_fprime
from scipy.linalg import inv, det
import derivation as der

def Neuton(f, x, e = 0.001):
    jac = np.array([der.Jacobi(f,x)])
    hes = np.array([der.Hessean(f,x)])
    del_x = np.linalg.lstsq(hes, -jac, rcond=None)[0]
    x_ = x + del_x
    while np.abs(np.max(x_ - x)) > e:
        jac = np.array(der.Jacobi(f,x))
        hes = np.array(der.Hessean(f,x))
        del_x = np.linalg.lstsq(hes, -jac, rcond=None)[0]
        x_ = x + del_x
    return x_
