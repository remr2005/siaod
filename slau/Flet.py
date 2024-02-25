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

if __name__ == "__main__":
    Flet(np.array([[5,0,1],[1,3,-1],[-3,2,10]]),np.array([[11],[4],[6]]))
    print(1)