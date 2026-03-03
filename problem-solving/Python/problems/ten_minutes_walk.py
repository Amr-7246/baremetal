def is_valid_walk(walk:list):
    if len(walk) != 10 : return False

    x, y = 0, 0
    
    for step in walk :
        if step == "n":
            y += 1
        if step == "s":
            y -= 1
        if step == "e":
            x += 1
        if step == "w":
            x -= 1

    return True if x == 0 and y == 0 else False
print( is_valid_walk(['n','n','n','s','n','s','n','s','n','s']) )