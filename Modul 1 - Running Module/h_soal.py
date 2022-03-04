def hitungVolume(r, t):
    pi = 3.14
    return r*r*pi*t

r,t = input().split()
print(hitungVolume(int(r), int(t)))
