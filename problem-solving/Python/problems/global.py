import string
import re

def convert_to_title(string:str): 
    arr = []
    result = " ".join(arr)
    for i in string.split(" ") :
        if i.isdigit() : arr.append(i)
        for idx, value in enumerate(i) :
            if value.isalpha() :
                if idx > 0:
                    arr.append(i[:idx] + value.upper() + i[idx+1:].lower() )
                    break
                
                arr.append(i[idx].upper() + i[1:].lower() )
                break
            
    arr_I =  " ".join([i.capitalize() for i in string.split(" ") if len(i) > 1  ])
    return arr_I
    
print(convert_to_title("""When Die. 15 Then You Will Realize' sGHhould equal 'when '#%when   I Die. Then You Will Realize"""))