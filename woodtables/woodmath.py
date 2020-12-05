C = 2#get_number()
N = 8#get_number()
P = 3#get_number()
W = 8#get_number()

# Para visualizar el problema
def genList(w):
    suma = w
    L=[]
    for i in range(N):
        if suma>=P:
            L.append(P)
            suma -= P
        else:
            L.append(suma)
            suma-=suma

if P <= C:
    print(W//C)
elif P > C:
    sobra = W % P   # slots filled del ultimo pocket
    llenos = W // P # pockets llenos
    vacios = N - llenos if sobra == 0 else N -llenos - 1 #pockets vacios
    
    Tstart = min(W // C, vacios)
    W -= Tstart * C
    if W // C == 0:
        print(Tstart)
        exit()
    sobra = W % P   # slots filled del ultimo pocket
    llenos = W // P # pockets llenos
    vacios = 0 #pockets vacios
    T = 0
    while True:
        if sobra > C:
            break
        elif sobra == C:
            T += 1
            break
        elif sobra < C:
            sobra = P - C + sobra
            llenos -=1
            T+=1
    print(Tstart+T)
        
    # e =  P - C # encumberment factor
    # T=0
    # print(f"llenos:{llenos}\nsobra:{sobra}\nvacios:{vacios}\ne:{e}\nT:{T}")
