#include <iostream> 
#include <iomanip>  
#include <sstream>  
#include <string>   
#include <vector>   

std::string rgb(int r, int g, int b) {
    std::vector<int> colors = {r, g, b};
    std::stringstream ss;

    for (int num : colors) {
        if (num < 0) num = 0;
        if (num > 255) num = 255;

        ss << std::hex << std::uppercase << std::setw(2) << std::setfill('0') << num;
    }

    return ss.str();
}

int main() {
    std::cout << rgb(0, 66, 165) << std::endl;
    return 0;
}
