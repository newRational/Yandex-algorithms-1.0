n = int(input())
d = dict()
l = []

for i in range(n-1):
    m = int(input())
    for j in range(m):
        s = input()
        if s not in d:
            d[s]=1
        else:
            d[s] += 1

m = int(input())
for j in range(m):
    s = input()
    if s not in d:
        d[s]=1
    elif d[s] == n-1:
        l.append(s)

print(len(l))
for v in l:
    print(v)
print(len(d))
for k in d:
    print(k)