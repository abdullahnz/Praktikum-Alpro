nilai = int(input())
jumlah = 0
n = 0

while nilai != -1:
    n = n + 1
    jumlah = jumlah + nilai
    nilai = int(input())

if n == 0:
    rata2 = 0.0
else:
    rata2 = jumlah / n
    print(rata2)
