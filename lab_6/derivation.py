def first(f, x, n):
    h = 0.0001
    f_0 = f(x)
    x_ = x.copy()
    x_[n]+=h
    f_1=f(x_)
    return (f_1-f_0)/h

def second(f,x,n):
    h = 0.0001
    f_0 = first(f,x,n)
    x_ = x.copy()
    x_[n]+=h
    f_1=first(f,x_,n)
    return (f_1-f_0)/h

def Jacobi(f,x):
    return [first(f,x,i) for i in range(len(x))]

def Hessean(f,x):
    return [second(f,x,i) for i in range(len(x))]