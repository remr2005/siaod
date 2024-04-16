import commiviyagor
import bb
import gen_alg
import numpy as np

def f(x):
    if 3*x[0]+4*x[1] > 12 or 3*x[0]+2*x[1] > 6:
        return 0
    else:
        return x[0] + 3*x[1]

def main():
    # print("Задача комивояжора")
    # print(commiviyagor.RUN(
    #     np.array([[float("inf"),  2, 3],
    #           [4, float("inf"), 6],
    #           [7, 6, float("inf")]])
    # ))
    print("генетический алгоритм")
    print(gen_alg.run(f,2,100,1000,(0,4)))
    # print("Перебор")
    # print(bb.perebor(f,(0,4)))

    
if __name__ == "__main__":
    main()