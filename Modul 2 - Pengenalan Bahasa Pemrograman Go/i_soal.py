n = int(input())
jumlah = 0

for i in range(n):
    t1,t2,t3 = input().split()
    if int(t1) + int(t2) + int(t3) >= 2 :
        jumlah = jumlah + 1

print(jumlah)
