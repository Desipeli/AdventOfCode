import sys
from functions import is_safe, is_safe2

input_file = sys.argv[1]

input = []
safe_reports = 0

with open(input_file, "r") as f:
    input = f.read()

for report in input.split("\n"):
    levels = report.split(" ")
    safe, index = is_safe2(levels)
    if safe:
        safe_reports += 1

    else:
        for i in range(-1,2):
            if index == 0 and i == -1: continue
            if is_safe(levels[0:i+index]+levels[index+i+1:]):
                safe_reports += 1
                break

print(safe_reports)