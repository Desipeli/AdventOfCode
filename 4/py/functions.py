import os
import time
from dataclasses import dataclass

class Colors:
    WHITE = "37"
    RED = "31"
    GREEN = "32"
    YELLOW = "33"

class BgColors:
    BLACK = "40"
    YELLOW = "43"

@dataclass
class OutputChar():
    color: str = Colors.WHITE
    bgcolor: str = BgColors.BLACK
    char: str = ""

def visualize(output, outputfile, sleep=0.1):
    with open(outputfile, "w", encoding="utf-8") as f:
        for row in output:
            f.write("\n")
            for oc in row:
                f.write(f"\033[1;{oc.color};{oc.bgcolor}m {oc.char}")
        time.sleep(sleep)

def find_xmas_p1(input: list, row_i: int, col_i: int, dir: tuple, word: str, output: list, outputfile: str) -> int:
    output_char = output[row_i][col_i]
    output[row_i][col_i] = OutputChar(color=output[row_i][col_i].color, bgcolor=BgColors.YELLOW,char=output_char.char)
    if row_i > 2:
        visualize(output, outputfile)
    word += input[row_i][col_i]
    if word == "XMAS": 
        output[row_i][col_i] = OutputChar(color=Colors.GREEN, bgcolor=BgColors.BLACK ,char=output_char.char)
        return 1
    if not "XMAS".startswith(word):
        output[row_i][col_i] = OutputChar(color=output[row_i][col_i].color, bgcolor=BgColors.BLACK, char=output_char.char)
        return 0
    possible, next_row_i, next_col_i = next_coord_p1(input, row_i, col_i, dir)
    if not possible:
        output[row_i][col_i] = OutputChar(color=output[row_i][col_i].color, bgcolor=BgColors.BLACK, char=output_char.char)
        return 0
    result = find_xmas_p1(input, next_row_i, next_col_i, dir, word, output, outputfile)
    if result == 1:
        output[row_i][col_i] = OutputChar(color=Colors.GREEN, bgcolor=BgColors.BLACK, char=output_char.char)
    else:
        output[row_i][col_i] = OutputChar(color=output[row_i][col_i].color, bgcolor=BgColors.BLACK, char=output_char.char)
    return result

def next_coord_p1(input: list, row_i: int, col_i: int, dir: tuple) -> tuple[bool, int, int]:
    new_row_i = row_i + dir[0]
    new_col_i = col_i + dir[1]
    if new_row_i < 0 or new_row_i >= len(input) or new_col_i < 0 or new_col_i >= len(input[0]):
        return False, 0, 0
    return True, new_row_i, new_col_i

def find_xmas_p2(input: list, row_i, col_i) -> int:
    if not (input[row_i-1][col_i-1] == "S" and input[row_i+1][col_i+1] == "M"
        or input[row_i-1][col_i-1] == "M" and input[row_i+1][col_i+1] == "S") :
        return 0
    if not (input[row_i+1][col_i-1] == "S" and input[row_i-1][col_i+1] == "M"
        or input[row_i+1][col_i-1] == "M" and input[row_i-1][col_i+1] == "S"):
        return 0
    return 1