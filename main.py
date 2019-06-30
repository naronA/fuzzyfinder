import numpy as np


def diagonal(n1, n2, pt):
    if (n1 == n2):
        return pt['MATCH']
    else:
        return pt['MISMATCH']


def pointers(di, ho, ve):
    pointer = max(di, ho, ve)
    if(di == pointer):
        return 'D'
    elif(ho == pointer):
        return 'H'
    else:
        return 'V'


def needleman_wunsch(s1, s2, match=2, mismatch=-1, gap=-2):
    penalty = {
        'MATCH': match,
        'MISMATCH': mismatch,
        'GAP': gap
    }
    n = len(s1) + 1
    m = len(s2) + 1
    al_mat = np.zeros((m, n), dtype=int)
    p_mat = np.zeros((m, n), dtype=str)

    for i in range(m):
        al_mat[i][0] = penalty['GAP'] * i
        p_mat[i][0] = 'V'

    for j in range(n):
        al_mat[0][j] = penalty['GAP'] * j
        p_mat[0][j] = 'H'

    p_mat[0][0] = 0
    for i in range(1, m):
        for j in range(1, n):
            di = al_mat[i-1][j-1] + diagonal(s1[j-1], s2[i-1], penalty)
            ho = al_mat[i][j-1] + penalty['GAP']
            ve = al_mat[i-1][j] + penalty['GAP']
            al_mat[i][j] = max(di, ho, ve)
            p_mat[i][j] = pointers(di, ho, ve)

    print(np.matrix(al_mat))
    print(np.matrix(p_mat))


if __name__ == '__main__':
    needleman_wunsch('ATTGC', 'ATGC')
