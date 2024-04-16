import numpy as np

def reduce(a_ : np.array):
    a = a_.copy()
    row_mins = np.min(a, axis=1)
    a = a.T - row_mins
    a = a.T
    col_mins = np.min(a, axis=0)
    a = a - col_mins
    L = sum(col_mins) + sum(row_mins)
    return a, L

def run(a: np.array, L=0, sr = {}, sc = {}, s = []):
    print()
    if len(a) == 1:
        return s, L
    zero_dugs = {}
    for i in range(len(a)):
        for j in range(len(a[0])):
            if a[i,j] == 0:
                a[i,j] = float("inf")
                di = min(a[i])
                dj = min(a.T[j])
                zero_dugs[(i,j)] = di + dj
                a[i,j] = 0
    maxis = max(zero_dugs.values())
    dugs = {k: v for k,v in zero_dugs.items() if v == maxis}
    best = list(dugs.keys())[0]
    neg = L + maxis
    a_ = np.delete(a,best[0], axis= 0)
    a_ = np.delete(a_,best[1], axis= 1)
    a_, L_ = reduce(a_)
    pos = L_+ L
    print(a)
    if neg < pos:
        a__ = a.copy()
        a__[best[0],best[1]]=float("inf")
        
        return run(a__,neg, sr, sc, s)
    elif pos <= neg:
        for i in range(best[0],len(sr)):
            sr[i]+=1
        for i in range(best[1],len(sc)):
            sc[i]+=1
        print(best, sr, sc)
        s +=  [(sr[best[0]], sc[best[1]])]
        return run(a_,pos,sr,sc,s)
    
def RUN(a):
    a, L = reduce(a)
    s, L_ =  run(a, L, {i : i for i in np.arange(len(a))},{i : i for i in np.arange(len(a))}, [])
    s += [(s[-1][1],s[0][0])] 
    return s, L_