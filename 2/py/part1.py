import sys
from functions import is_safe

input_file = sys.argv[1]

input = []
safe_reports = 0


with open(input_file, "r") as f:
    input = f.read()

for report in input.split("\n"):
    levels = report.split(" ")
    if is_safe(levels):
        safe_reports += 1

print(safe_reports)