import sys

inputFile = sys.argv[1]
input = ""

with open(inputFile, "r") as f:
    input = f.read()

list1 = []
list2 = []
total_distance = 0

for row in input.split("\n"):
    splitted = row.split("   ")
    list1.append(int(splitted[0]))
    list2.append(int(splitted[-1]))

list1.sort()
list2.sort()

for i in range(len(list1)):
    total_distance += abs(list1[i]-list2[i])

print(total_distance)