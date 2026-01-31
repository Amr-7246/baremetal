from collections import Counter

"""
Outputs. . . 
- collictions lib, Counter(list).items(element, count that it apear)
- set(), list.count()
"""

def find_unique_number(arr:list):
    repeated = ""
    
    if arr[0] == arr[1]:
        repeated = arr[0]
    if arr[0] == arr[2]:
        repeated = arr[0]
    else :
        repeated = arr[1]
    
    return next(filter( lambda i: i != repeated , arr))

def find_unique_number_I(arr:list):
    # a, b = set(arr)
    # return a if arr.count(a) == 1 else b 
    return [i for i in set(arr) if arr.count(i) == 1][0]

def find_unique_number_II(arr:list):
    return next(v for v,c in Counter(arr).items() if c == 1 )

# print(find_unique_number_II([12,12,5,12]))