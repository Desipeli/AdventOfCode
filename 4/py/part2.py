import sys
from functions import find_xmas_p2

directions = [
    # (y, x)
    (0,1),
    (0,-1),
    (1,1),
    (1,0),
    (1,-1),
    (-1,0),
    (-1,1),
    (-1,-1)
]


def main(input):
    input = input.split("\n")
    total_xmas = 0
    for row_i, row in enumerate(input):
        for col_i, char in enumerate(row):
            if row_i == 0 or row_i == len(input) -1 or col_i == 0 or col_i == len(row)-1 or char != "A":
                continue
            total_xmas += find_xmas_p2(input, row_i, col_i)

    
    return total_xmas


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("provide inputfile")

    inputfile = sys.argv[1]
    input = ""

    with open(inputfile, "r") as f:
        input = f.read()
    
    print(main(input))