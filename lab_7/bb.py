def perebor(f,bor):
    x= []
    for i in range(bor[0],bor[1]+1):
        for j in range(bor[0],bor[1]+1):
            x.append((i,j))
    m_x = 0
    maxs = -float("inf")
    for i in range(len(x)):
        if f(x[i]) > maxs:
            maxs = f(x[i])
            m_x = i
    return x[m_x]