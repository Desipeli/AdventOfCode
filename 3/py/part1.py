import sys
import string

if len(sys.argv) < 2:
    print("provide input file")

inputfile = sys.argv[1]
input = ""

with open(inputfile, "r") as f:
    input = f.read()

current_string = ""
first_value = ""
second_value = ""
total = 0

for char in input:
    
    if char == "m":
        current_string = "m"
    elif char == "u" and current_string == "m":
        current_string = "mu"
    elif char == "l" and current_string == "mu":
        current_string = "mul"
    elif char == "(" and current_string == "mul":
        current_string = "mul("
        first_value = ""
        second_value = ""
    elif char in string.digits:
        if current_string == "mul(":
            first_value += char
        elif current_string == "mul(,":
            second_value += char
    elif char == "," and current_string == "mul(":
        current_string = "mul(,"
    elif char == ")" and current_string == "mul(,":
        try:
            current_string = ""
            total += int(first_value)*int(second_value)
        except:
            pass
    else:
        current_string = ""

print(total)