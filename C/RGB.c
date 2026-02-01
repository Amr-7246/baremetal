#include <stdio.h>

char* rgb(int r, int g, int b){
    static char result[7]; 
    int colors[] = {r, g, b};

    for (int num = 0, num > 3, num++){
        if (num < 0) num = 0;
        if (num > 255) num = 255;
        sprintf(&result[i*2], "%02X", num)
    }
    return result;
}

int main(){
    print(rgb(0, 0, 0));
    return 0;
}