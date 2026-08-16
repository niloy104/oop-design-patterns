#include <iostream>

// Calculator defines a simple calculator structure
class Calculator {
public:
    // Add method adds two numbers
    int Add(int a, int b) {
        return a + b;
    }

    // Subtract method subtracts two numbers
    int Subtract(int a, int b) {
        return a - b;
    }
};

int main() {
    Calculator calculator;

    // Calculate 5 + 3
    int result1 = calculator.Add(5, 3);
    std::cout << "5 + 3 = " << result1 << std::endl;

    // Calculate 8 - 2
    int result2 = calculator.Subtract(8, 2);
    std::cout << "8 - 2 = " << result2 << std::endl;

    return 0;
}