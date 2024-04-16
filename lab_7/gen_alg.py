import random
import numpy as np

class osob:
    dnk = []
    bor = (-float("inf"), float("inf"))
    def __init__(self, bor,  dnk = []) -> None:
        self.dnk = dnk
        self.bor = bor
    def randomize(self, n):
        self.dnk = []
        for j in range(n):
            self.dnk.append(random.randint(self.bor[0],self.bor[1]))
    def mutate(self):
        n1 = 0
        n2 = 0
        while n1 == n2:
            n1 = random.randint(0,len(self.dnk)-1)
            n2 = random.randint(0,len(self.dnk)-1)
        buf = self.dnk[n1]
        self.dnk[n1] =self.dnk[n2]
        self.dnk[n2] = buf
    def reproduce(self,osobs):
        n = random.randint(1,len(self.dnk)-1)
        return [osob(self.bor, self.dnk[:n]+osobs.dnk[n:]), osob(self.bor, osobs.dnk[:n]+self.dnk[n:]) ]
        
def random_osobi(bor, n, arg_n):
    osobi = []
    for i in range(n):
        o = osob(bor,[])
        o.randomize(arg_n)
        osobi.append(o)
    return osobi

def iteration_alg(f, n_osob, osobi):
    f_osobi = [f(i.dnk) for i in osobi]
    f_sum = sum(f_osobi)
    p = [f_osobi[i]/f_sum for i in range(n_osob)]
    new_osobi = []
    for i in range(n_osob//4):
        n1 = np.random.choice(len(p), p=p)
        n2 = np.random.choice(len(p), p=p)
        n3 = np.random.choice(len(p), p=p)
        n4 = np.random.choice(len(p), p=p)
        new_osobi += osobi[n1].reproduce(osobi[n2])
        osobi[n3].mutate()
        osobi[n4].mutate()
        new_osobi += [osobi[n3],osobi[n4]]
    return new_osobi

def run(f,arg_n,n_ep = 100000, n_osob = 1000, bor = (-float("inf"), float("inf")), e = 0.0001):
    osobi = random_osobi(bor, n_osob, arg_n)
    for i in range(n_ep):
        print(i)
        osobi = iteration_alg(f,n_osob,osobi)
    f_osobi = [f(i.dnk) for i in osobi]
    f_sum = sum(f_osobi)
    p = [f_osobi[i]/f_sum for i in range(n_osob)]
    slov = {p[i]:osobi[i] for i in range(n_osob)}
    p = sorted(p)
    return slov[p[-1]].dnk