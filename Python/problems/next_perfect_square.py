import math

def find_next_square(n) : 
    return int(-1 if (n ** 0.5 - int(n ** 0.5)) > 0 else (int(n** 0.5) + 1) ** 2 )

def find_next_square_1(n) : 
    return int(-1 if (math.sqrt(n) - int(math.sqrt(n))) > 0 else (int(math.sqrt(n)) + 1) ** 2 )
