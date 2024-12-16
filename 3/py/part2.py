import sys
import string

if len(sys.argv) < 2:
    print("provide input file")

inputfile = sys.argv[1]
input = ""

with open(inputfile, "r") as f:
    input = f.read()


def parseMul(s: string) -> int:
    if len(s)<5 or s[0] != "(": return 0
    first = ""
    second = ""
    comma_found = False
    for c in s[1:]:
        if c in string.digits:
            if not comma_found:
                first += c
            else:
                second += c
        elif c == "," and not comma_found:
            comma_found = True
        elif c == ")":
            try:
                return int(first)*int(second)
            except Exception as e:
                print(e)
                break
        else:
            break
    return 0

input = input.split("don't()")
enabled = True
total = 0

for i in input:
    dos = i.split("do()")
    if not enabled:
        if len(dos) == 0:
            continue
        dos = dos[1:]
    for d in dos:
        muls = d.split("mul")
        for m in muls:
            total += parseMul(m)

    enabled = False

print(total)