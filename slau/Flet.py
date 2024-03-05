import numpy as np

def  Flet(a,b,e =0.0001):
    x = np.ones((len(a),1))
    r = b - np.dot(a,x)
    r_ = r
    z = r
    while np.dot(r.T,r)[0,0]/np.dot(b.T,b)[0,0] >= e:
        alpha = np.dot(r.T,r)[0,0]/np.dot(np.dot(a,z).T,z)[0,0]
        x = x + alpha * z
        r = r - alpha * np.dot(a,z)
        betta = np.dot(r.T,r)[0,0]/np.dot(r_.T,r_)[0,0]
        z = r + betta*z
    print("A*X")
    print(np.dot(a,x))
    print("Должно быть равно")
    print(b)
    return x
import numpy as np

def rafs(arr_func, e, bord):
    x = np.array([1, 1, 1, 1, 1, 1])
    W = np.ones((6, 6))

    def calc_W(a_):
        a = a_.copy()
        for i in range(len(a)):
            for j in range(len(a[i])):
                a[i][j] = derivation_first_deriv_array(arr_func[i], x, j)
        return a

    W = calc_W(W)

    F = np.array([func(x) for func in arr_func])

    def calc_del_X(F, W):
        matrixF = np.array(F).reshape(-1, 1)
        matrix = np.array(W)

        matrix = np.linalg.inv(matrix)
        matrixF = np.dot(matrix, matrixF)

        delta_x = matrixF.flatten()
        return delta_x

    del_x = calc_del_X(F, W)
    print(x, del_x)
    x = x + del_x
    print(x)
    sigma = np.min(x)
    print(sigma)
    k = 0
    print(1)
    while sigma > e:
        k += 1
        if k > bord:
            break
        W = calc_W(W)
        F = np.array([func(x) for func in arr_func])
        del_x = calc_del_X(F, W)
        x = x + del_x
        sigma = np.min(x)
    return x

# Функция для вычисления первой производной
def derivation_first_deriv_array(func, x, index):
    h = 1e-8
    x1 = x.copy()
    x2 = x.copy()
    x1[index] += h
    x2[index] -= h
    return (func(x1) - func(x2)) / (2 * h)

# Пример использования
def func1(x):
    return x[0] + x[1] + x[2] + x[3] + x[4] + x[5] - 20

def func2(x):
    return -x[1] + 2*x[2] + 3*x[3] + x[4] + 2*x[5] + 3

def func3(x):
    return x[0] + x[2] + 2*x[3] - x[4] + x[5] - 12

def func4(x):
    return -x[0] + 2*x[1] + x[3] + 3*x[4] - 2*x[5] - 5

def func5(x):
    return 4*x[0] + 2*x[1] + x[2] - x[3] + 3*x[4] - 8

def func6(x):
    return 3*x[0] + x[1] + 2*x[2] - x[4] + 4*x[5] - 10

import numpy as np

def jacobi_method(funcs, x, epsilon=1e-6, max_iterations=100):
    n = len(x)
    x_new = np.zeros(n)
    
    for _ in range(max_iterations):
        for i in range(n):
            x_new[i] = funcs[i](x)
        
        delta = np.linalg.norm(x_new - x)
        if delta < epsilon:
            return x_new
        
        x = x_new.copy()
    
    raise ValueError("Метод не сошелся за заданное количество итераций.")

# Массив функций
arr_func = [func1, func2, func3, func4, func5, func6]
matrix = np.array([
    [2, 1, 3, 4, 5, 6],
    [0, -1, 2, 3, 1, 2],
    [1, 0, 1, 2, -1, 1],
    [-1, 2, 0, 1, 3, -2],
    [4, 2, 1, 0, -1, 3],
    [3, 1, 2, -1, 0, 4]
], dtype=float)
# Вызов функции rafs
res2 = jacobi_method(arr_func,np.array([1, 1,1,1,1,1]),1e-5,1000000000000)
result = rafs(arr_func, 1e-5, 10000)
print("Результат:", result, np.dot(matrix, result),np.dot(matrix, res2))
# if __name__ == "__main__":
#     print(Flet(np.array([[4,1,-1],[2,3,0],[1,-1,5]]),np.array([[7],[7],[10]])))
#     print(1)