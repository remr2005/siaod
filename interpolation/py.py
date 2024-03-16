import numpy as np
import matplotlib.pyplot as plt
from scipy import integrate
import math

def fourier_series_coefficients(f, interval, n_terms):
    a, b = interval
    coefficients = []
    a0 = (1/(b-a)) * integrate.quad(f, a, b)[0]
    coefficients.append(a0)
    
    for n in range(1, n_terms+1):
        an = (1/(b-a)) * integrate.quad(lambda x: f(x)*np.cos(2*np.pi*n*x/(b-a)), a, b)[0]
        bn = (1/(b-a)) * integrate.quad(lambda x: f(x)*np.sin(2*np.pi*n*x/(b-a)), a, b)[0]
        coefficients.append(an)
        coefficients.append(bn)
    
    return coefficients

def fourier_series(x, coefficients, interval):
    a, b = interval
    result = coefficients[0]/2
    for n in range(1, len(coefficients)//2):
        result += coefficients[2*n-1]*np.cos(2*np.pi*n*x/(b-a)) + coefficients[2*n]*np.sin(2*np.pi*n*x/(b-a))
    
    return result

f = lambda x: (2/math.pi)*math.asin(math.sin(x))
# f = lambda x: x
interval = (-np.pi, np.pi)
n_terms = 270
coefficients = fourier_series_coefficients(f, interval, n_terms)

# x_values = np.linspace(interval[0], interval[1], 1000)
x_values = np.array([0,math.pi/2,math.pi])
y_values = [fourier_series(x, coefficients, interval) for x in x_values]

print(y_values)
plt.plot(x_values, y_values)
plt.show()