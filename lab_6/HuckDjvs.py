import numpy as np

def Hoock(f,
          x,
          h=0.1,
          h_min=0.0000001,
          n=100000):
    i = 0
    buf = []
    for j in range(len(x)):
        buf.append(1.0)
    h = h * np.array(buf)
    
    while i < n:
        i += 1
        for j in range(len(x)):
            x_1 =x.copy()
            x_2 = x.copy()
            x_1[j] += h[j]
            print(x_2[j])
            x_2[j] -= h[j]
            print(x_2[j])
            
            fx = f(x)
            f1 = f(x_1)
            f2 = f(x_2)
            
            if f1 < fx:
                x = x_1
            elif f2 < fx:
                x = x_2
            else:
                h[j] /= 10
                if max(h) <= h_min:
                    return x
    return x