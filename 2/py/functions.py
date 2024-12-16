
def is_safe(levels: list) -> bool:
    desc = None
    for i, level in enumerate(levels[1:]):
        diff = int(level) - int(levels[i])
        if diff > 0:
            if desc == True:
                return False
            desc = False
        else:
            if desc == False:
                return False
            desc = True
        if diff < -3 or diff > 3 or diff == 0:
            return False
    return True



def is_safe2(levels: list):
    desc = None
    for i, level in enumerate(levels[1:]):
        diff = int(level) - int(levels[i])
        if diff > 0:
            if desc == True:
                return False, i
            desc = False
        else:
            if desc == False:
                return False, i
            desc = True
        if diff < -3 or diff > 3 or diff == 0:
            return False, i
    return True, 0