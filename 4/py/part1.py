import sys
from functions import find_xmas_p1, Colors, OutputChar, BgColors, visualize

directions = [
    # (row, col)
    (0,1),
    (0,-1),
    (1,1),
    (1,0),
    (1,-1),
    (-1,0),
    (-1,1),
    (-1,-1)
]


def main(input_string, outputfile):
    input = input_string.split("\n")
    output = []
    for row in input:
        new_list = []
        for char in row:
            new_list.append(
                OutputChar(char=char)
            )
        output.append(new_list)
    total_xmas = 0

    for row_i, row in enumerate(input):
        for col_i, char in enumerate(row):
            if char != "X":
                output[row_i][col_i] = OutputChar(color=output[row_i][col_i].color, bgcolor=BgColors.YELLOW, char=output[row_i][col_i].char)
                visualize(output, outputfile, sleep=0.05)
                output[row_i][col_i] = OutputChar(color=output[row_i][col_i].color, bgcolor=BgColors.BLACK, char=output[row_i][col_i].char)
                visualize(output, outputfile, sleep=0.05)
                continue
            for dir in directions:
                res = find_xmas_p1(input, row_i, col_i, dir, "", output, outputfile)
                if res:
                    total_xmas += res
                    output[row_i][col_i] = OutputChar(color=Colors.GREEN, bgcolor=BgColors.BLACK, char=output[row_i][col_i].char)

    print(output)
    return total_xmas


if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("provide inputfile and outputfile")

    inputfile = sys.argv[1]
    outputfile = sys.argv[2]
    input = ""

    with open(inputfile, "r") as f:
        input = f.read()
    
    print(main(input, outputfile))