# KISS = Keep It Simple, Stupid

It's a software design principle. The idea: most systems work best when you keep them simple, not complex. Complexity is the enemy, it makes code harder to read, harder to debug, harder to change later.

## Why it matters

When you write code, you're not just writing it for the computer. You're writing it for other humans too, your teammates, and even future-you after 6 months when you forgot why you wrote it. If the code is too clever or too complicated, people waste time just trying to understand it instead of fixing bugs or adding features.

## What "simple" means here

* Simple doesn't mean "dumb" or "low quality." It means: no unnecessary complexity. Solve the problem in the most straightforward way possible.
* If there are two ways to solve something, one simple and one clever/complex, and both work fine, pick the simple one.
* Simple code is easier to test, easier to debug, easier to onboard new people to.

## Guidelines for writing code that satisfies the KISS principle

* Do not use technologies your colleagues may not understand when implementing code.
* Do not reinvent the wheel; make good use of existing libraries.
* Do not over-optimize.

## Example

Here is an example of a simple calculator program designed with the KISS principle:

**C++**

```cpp
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
```

**Go**

```go
package main

import "fmt"

// Calculator defines a simple calculator structure
type Calculator struct{}

// Add method adds two numbers
func (c Calculator) Add(a, b int) int {
    return a + b
}

// Subtract method subtracts two numbers
func (c Calculator) Subtract(a, b int) int {
    return a - b
}

func main() {
    calculator := Calculator{}

    // Calculate 5 + 3
    result1 := calculator.Add(5, 3)
    fmt.Println("5 + 3 =", result1)

    // Calculate 8 - 2
    result2 := calculator.Subtract(8, 2)
    fmt.Println("8 - 2 =", result2)
}
```

In the above example, we defined a simple calculator structure `Calculator`, containing the `Add` and `Subtract` methods to perform addition and subtraction. With simple design and implementation, this calculator program is clear, easy to understand, and meets the requirements of the KISS principle.

## Quick way to remember it

Ask yourself, "Can a new developer read this and understand it in 5 minutes?" If yes, you're following KISS. If they need a whiteboard and 2 hours, you're overcomplicating it.