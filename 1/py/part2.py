import sys

inputFile = sys.argv[1]
input = ""

with open(inputFile, "r") as f:
    input = f.read()

left_side = {}
right_side = {}
similarity_score = 0

for row in input.split("\n"):
    splitted = row.split("   ")
    first_value = int(splitted[0])
    if first_value not in left_side:
        left_side[first_value] = 0
    left_side[first_value] += 1

    second_value = int(splitted[-1])
    if second_value not in right_side:
        right_side[second_value] = 0
    right_side[second_value] += 1

for key, value in left_side.items():
    if key in right_side:
        similarity_score += key*value*right_side[key]

print(similarity_score)