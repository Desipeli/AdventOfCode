import sys
import re

if len(sys.argv) < 2:
    print("provide input file")

inputfile = sys.argv[1]
input = ""

with open(inputfile, "r") as f:
    input = f.read()

muls = re.findall("mul\(\d+,\d+\)", input)
total = 0


for mul in muls:
    first, second = re.findall("(\d+,\d+)", mul)[0].split(",")
    try:
        total += int(first)*int(second)
    except Exception as e:
        print(e)

print(total)
print(sum([ int(x[0])*int(x[1]) for x in [re.findall("(\d+,\d+)", mul)[0].split(",") for mul in re.findall("mul\(\d+,\d+\)", input)]]))