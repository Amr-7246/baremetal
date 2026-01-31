import re

"""
Outputs. . . 
- iter(), zip()
- re lib AKA regular expresion lib , and its own function like findall()
- ljust AKA left justify 
"""


#& virsion_I
def split_str(str:str) -> str:
    word = str
    result = [] 
    
    while len(word) > 0 :
        
        if len(word) == 1 :
            result.append(f"{word}_")
            break
        
        splited_word = word[:2]
        word= word[2:]
        result.append(splited_word)
    
    return f"Here is You splited string {result}"

#& virsion_II
def split_str_2(str:str) -> list:
    if (len(str) % 2 != 0) :
        str += "_"
    return [str[i:i+2] for i in range(0, len(str), 2)]

#& virsion_III
def _split_str_3(s:str) -> list:
    return re.findall('.{2}', s + "_")
    #~ OR: return re.findall('..', s + "_")

#& virsion_IV
def _split_str_4(s:str) -> list:
    return [ s[i:i+2].ljust(2, '_' ) for i in range(0, len(s), 2) ]

#^ virsion_V, The Gloden virsion
def _split_str_5(s:str) -> list:
    return list(map("".join , zip( *[iter(s + '_')] * 2 ) ) ) 
