from problems import split_str, find_next_square, find_unique_number #type:ignore
import problems

while True :
    user_input = input("inject Your input . . . ")
    if(user_input == "close"):
        print("sorry to see You go . . ")
        break
    
    #& problems functions 
    print(split_str(user_input))
    print(find_next_square(int(user_input)))
    print(problems.find_unique_number(list(user_input)))
