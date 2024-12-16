import sys
import re

if len(sys.argv) < 2:
    print("provide input file")

inputfile = sys.argv[1]
input = ""

with open(inputfile, "r") as f:
    input = f.read()

commands = re.findall("mul\(\d+,\d+\)|do\(\)|don't\(\)", input)
total = 0
enabled = True

for command in commands:
    match command:
        case "do()":
            enabled = True
        case "don't()":
            enabled = False
        case _:
            if not enabled: continue
            first, second = re.findall("(\d+,\d+)", command)[0].split(",")
            try:
                total += int(first)*int(second)
            except Exception as e:
                print(e)

print(total)