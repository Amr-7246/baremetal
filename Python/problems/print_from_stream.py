import typing
import re
Input = [input().strip() for _ in range(int(input())) ] # [ '' , '' , .... ]
all_odd_nums = list(range(1 , 100 , 2)) # [1 , 3 , 5 . . .. ]
all_even_nums = [ i for i in range(100) if i % 2 == 0 ] # [0 , 2 , 4 . . .. ]
result = []

def print_form_stream(input) :
    if 1 >= len(input) >= 100  :
        return
    for I in input : 
        match = re.search(r'\d+' , I ) 
        if 1 >= int(match.group()) >= 10  : # type: ignore
            return
        if 'odd' in I :
            Num = int(I.split()[-1]) 
            odd_nums = '\n'.join( str(i) for i in [i for i in all_odd_nums[:Num] ] ) # 1 , 3 ...
            result.append(odd_nums)
        else :
            Num = int(I.split()[-1])
            even_nums = '\n'.join( str(i) for i in [i for i in all_even_nums[:Num] ] )
            result.append(even_nums)
    return result


print(*print_form_stream(Input) , sep='\n' ) # type: ignore
