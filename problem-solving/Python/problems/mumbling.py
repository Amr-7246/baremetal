def mumbling(st:str):
    return "-".join([  value.upper() + (value.lower() * i) for i , value in enumerate(st) ])

#& virsion_II
def mumbling_I(st:str):
    return "-".join([(value * (i+1)).title() for i , value in enumerate(st) ])

print(mumbling_I("igfofreoat"))