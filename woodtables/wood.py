# C = get_number()
# N = get_number()
# P = get_number()
# W = get_number()

C = 2#get_number()
N = 3#get_number()
P = 3#get_number()
W = 8#get_number()

if P <= C:
    print(W//C)
elif P > C:
    #pockets = [3 3 2]
    sobra = W % P
    llenos = W // P
    T = 0
    loops = 0
    while True:
        loops+=1
        print(f"llenos:{llenos}, sobra:{sobra}, T:{T}")
        if C > sobra:
            T +=1
            sobra = P - (C-sobra)
            llenos -= max(C // P, 1)
        elif C<=sobra:
            T+=1
            sobra -= C
        if llenos <= 0 or loops >4:
            break
    print(T)