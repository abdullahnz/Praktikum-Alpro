a,b,c = input().split()
d,e,f = input().split()
x = int(input())

if x == 0 or x == 360:
    print(a,b,c)
    print(d,e,f)
elif x == 90:
    print(d,a)
    print(e,b)
    print(f,c)
elif x == 180:
    print(f,e,d)
    print(c,b,a)
elif x == 270:
    print(c,f)
    print(b,e)
    print(a,d)
