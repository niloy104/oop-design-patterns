# Object-Oriented Programming

This note explains the main OOP topics with real-life examples and code in both C++ and Go.


---

## Table of Contents
1. [Object and Class](#1-object-and-class)
2. [The Four Pillars](#2-the-four-pillars)
3. [Access Modifiers](#3-access-modifiers)
4. [Types of Inheritance](#4-types-of-inheritance)
5. [The Diamond Problem](#5-the-diamond-problem)
6. [Method Overloading (Compile-time Polymorphism)](#6-method-overloading-compile-time-polymorphism)
7. [Method Overriding (Runtime Polymorphism)](#7-method-overriding-runtime-polymorphism)
8. [Compile-time vs Runtime Polymorphism](#8-compile-time-vs-runtime-polymorphism)
9. [Interfaces vs Concrete Types](#9-interfaces-vs-concrete-types)
10. [Coupling vs Cohesion](#10-coupling-vs-cohesion)
11. [Composition over Inheritance](#11-composition-over-inheritance)
12. [Association vs Aggregation vs Composition](#12-association-vs-aggregation-vs-composition)
13. [Cheat-sheet Table](#13-cheat-sheet-table)

---

## 1. Object and Class

**Simple idea:** A class is the blueprint. An object is the real thing built from that blueprint.

Think of a class as the drawing of a house. The drawing shows how many rooms, doors, and windows there will be. You cannot sleep inside a drawing though. You need to actually build the house. That built house is the object. You can build many houses from the same drawing, and each one is a separate house with its own furniture, but all of them follow the same design.

In code, a class defines data (fields) and behavior (methods). The class itself barely uses memory. Memory is used only when you create an object from it.

Why and when to use it: without this idea, every object would need its own hand-written design, with no reuse and no consistency. You use classes anywhere you model a real thing that has state and behavior, like a `User`, an `Order`, or a `RateLimiter`.

Go has no `class` keyword. It uses a `struct` for data and attaches behavior to it with functions called receiver functions. That is basically Go's version of a class.

**C++**
```cpp
#include <iostream>
#include <string>
using namespace std;

class Employee {                              // blueprint (the "class")
public:
    string name;                              // field: holds the name
    double salary;                            // field: holds the salary

    Employee(string n, double s) : name(n), salary(s) {}  // constructor: runs when object is created

    void showDetails() {                      // method: behavior tied to this object
        cout << name << " earns " << salary << endl;
    }
};

int main() {
    Employee e1("Abul", 50000);               // e1 is an object (real instance built from the blueprint)
    Employee e2("Karim", 65000);              // e2 is a separate, independent object
    e1.showDetails();                         // uses e1's own data
    e2.showDetails();                         // uses e2's own data, does not affect e1
}
```

**Go**
```go
package main

import "fmt"

type Employee struct { // struct = Go's version of a class (data only, no methods here)
    Name   string
    Salary float64
}

func NewEmployee(name string, salary float64) Employee { // acts like a constructor
    return Employee{Name: name, Salary: salary}          // builds and returns a filled-in Employee
}

func (e Employee) ShowDetails() { // receiver method: "e" is the object this method belongs to
    fmt.Println(e.Name, "earns", e.Salary)
}

func main() {
    e1 := NewEmployee("Abul", 50000)  // e1 is an object (instance)
    e2 := NewEmployee("Karim", 65000) // e2 is a separate object
    e1.ShowDetails()
    e2.ShowDetails()
}
```

---

## 2. The Four Pillars

The four pillars of OOP are Encapsulation, Abstraction, Inheritance, and Polymorphism. Below is each one explained simply.

### 2.1 Encapsulation

**Simple idea:** Keep an object's data private. Let it be touched only through its own public methods.

Real-life example: a capsule medicine. The active ingredients are sealed inside. You do not touch them directly, you just take the capsule and it works in a controlled way. Another example is a bank account. You cannot just reach in and change your `balance` to a huge number. You can only call `deposit()` or `withdraw()`, and the bank checks the rules behind that.

Why and when: this stops other code from putting your object into a wrong state, like a negative balance. You use it for any field where a wrong value is possible. To do it, make fields private, and expose behavior through public methods.

**C++**
```cpp
class Account {
private:
    double balance;                          // hidden field, nobody outside can touch this directly
public:
    Account(double initial) : balance(initial) {}
    double getBalance() const { return balance; }   // read-only access, safe to expose
    void deposit(double amount) {
        if (amount > 0) balance += amount;   // validation happens here, inside the class, not outside
    }
};
```

**Go**
```go
type Account struct {
    balance float64 // lowercase = private to this package, no direct outside access
}

func NewAccount(initial float64) *Account { return &Account{balance: initial} }

func (a *Account) Balance() float64 { return a.balance } // controlled read access
func (a *Account) Deposit(amount float64) {
    if amount > 0 {              // rule enforced inside the method, caller cannot skip it
        a.balance += amount
    }
}
```

### 2.2 Abstraction

**Simple idea:** Show only what is needed. Hide how it works underneath.

Real-life example: a coffee machine. You put in beans and water, press one button, and get coffee. You do not need to know about the heating coil or the pump inside. Same with your phone. You tap an icon, the app opens, and you do not need to know what happens inside the operating system.

Encapsulation and abstraction sound similar, so here is the clean way to tell them apart: encapsulation hides data. Abstraction hides how the work actually gets done, usually through interfaces or abstract classes.

Why and when: this lets you change the internal logic later, like switching from Stripe to PayPal, without breaking the code that calls it. In C++, you get this with abstract classes and pure virtual functions. In Go, you get this with interfaces. A Go type does not need to say "I implement this interface," it just needs to have the right methods.

**C++**
```cpp
class CloudStorage {                          // abstract class = the "what," not the "how"
public:
    virtual string upload(string fileName) = 0;   // pure virtual, no body, just a contract
    virtual ~CloudStorage() {}                // virtual destructor, good practice for base classes
};

class S3Storage : public CloudStorage {       // concrete class = the "how"
public:
    string upload(string fileName) override {
        return "uploaded " + fileName + " to S3"; // real AWS SDK complexity would be hidden in here
    }
};
```

**Go**
```go
type CloudStorage interface { // the contract, just method names, no implementation
    Upload(fileName string) string
}

type S3Storage struct{} // concrete type, actually does the work

func (S3Storage) Upload(fileName string) string {
    return "uploaded " + fileName + " to S3" // caller never sees the AWS SDK details
}
```

### 2.3 Inheritance

**Simple idea:** A child class reuses and extends the fields and methods of a parent class. This is called an "is-a" relationship.

Real-life example: a parent passes down traits to a child. A `Vehicle` might define general things like `speed` and `move()`. A `Car` and a `Bike` are both vehicles, so instead of writing `speed` and `move()` twice, both inherit from `Vehicle` and add what is special to them.

Why and when: this avoids repeating shared fields and behavior across related types. Use it only when there is a real "is-a" relationship, like a `Dog` is an `Animal`. If it is not a true "is-a," do not force inheritance. See section 11 for more on this.

C++ uses `class Child : public Parent`. Go has no class inheritance at all. It uses struct embedding to get a similar effect.

**C++**
```cpp
class Vehicle {
public:
    void move() { cout << "Vehicle is moving" << endl; }  // shared/common behavior
};

class Bike : public Vehicle {     // Bike IS-A Vehicle (inherits everything from Vehicle)
public:
    void ringBell() { cout << "Bike rings the bell" << endl; }  // Bike's own extra behavior
};

int main() {
    Bike b;
    b.move();      // inherited from Vehicle, Bike did not have to redefine this
    b.ringBell();  // Bike's own method
}
```

**Go**
```go
type Vehicle struct{}

func (Vehicle) Move() { fmt.Println("Vehicle is moving") } // shared behavior

type Bike struct {
    Vehicle // embedded struct, not real inheritance, but fields/methods get "promoted" up to Bike
}

func (Bike) RingBell() { fmt.Println("Bike rings the bell") } // Bike's own extra behavior

func main() {
    b := Bike{}
    b.Move()     // promoted from the embedded Vehicle, feels like inheritance
    b.RingBell() // Bike's own method
}
```

### 2.4 Polymorphism

**Simple idea:** Poly means many, morph means form. The same method call behaves differently depending on the actual object.

Real-life example: think about the word "approve" in a company. `approve()` can mean approving leave, approving a raise, or approving an invoice. Same action name, different logic depending on what is being approved. Or simply, a `Dog.speak()` barks, a `Cat.speak()` meows. Same method name, different result, based on which object you are holding.

Polymorphism has two types, compile-time and runtime, and both get a full section later (6, 7, 8). Here is a quick preview:

**C++**
```cpp
class Animal { public: virtual void speak() { cout << "..." << endl; } };       // base behavior
class Dog : public Animal { public: void speak() override { cout << "Bark" << endl; } };  // Dog's own version
class Cat : public Animal { public: void speak() override { cout << "Meow" << endl; } };  // Cat's own version

void makeItSpeak(Animal* a) { a->speak(); }  // does not care if it is really a Dog or a Cat

int main() {
    Dog d; Cat c;
    makeItSpeak(&d);  // prints Bark, decided by the real object type
    makeItSpeak(&c);  // prints Meow, decided by the real object type
}
```

**Go**
```go
type Animal interface{ Speak() } // the contract every animal type must follow

type Dog struct{}
func (Dog) Speak() { fmt.Println("Bark") }  // Dog's version

type Cat struct{}
func (Cat) Speak() { fmt.Println("Meow") }  // Cat's version

func makeItSpeak(a Animal) { a.Speak() } // works with any type that satisfies Animal

func main() {
    makeItSpeak(Dog{}) // Bark
    makeItSpeak(Cat{}) // Meow
}
```

---

## 3. Access Modifiers

**Simple idea:** Access modifiers control who can see or touch a class member.

Think of rooms in an office building. Public areas are open to anyone, like the lobby. Protected areas are open only to staff and their direct reports, like internal meeting rooms. Private areas are open to only one person, like their own locked office. Go does not have named keywords for this. It uses a simple rule: capitalize the name to expose it outside the package, use lowercase to keep it hidden.

| | C++ | Go |
|---|---|---|
| Public | `public:` section | Capitalized identifier, like `Name` |
| Protected | `protected:` section, visible to subclasses | No direct match |
| Private | `private:` section | Lowercase identifier, like `name`, hidden outside the package |
| Class-wide/shared | `static` | Package-level `var` |

**C++**
```cpp
class Employee {
public:
    string name;          // Public- visible everywhere
protected:
    string role;           // Private- visible to this class and its subclasses only
private:
    double salary;          // Private- visible only inside this class
public:
    static int companyCode; // Private- shared across every Employee object, only one copy exists
};
```

**Go**
```go
var CompanyCode int // package-level, closest thing Go has to "static", shared by everyone in the package

type Employee struct {
    Name   string  // Public- exported: capital N, visible outside the package
    role   string  // Private- unexported: lowercase r, visible only inside this package
    salary float64 // Private- unexported: same rule as role
}
```

---

## 4. Types of Inheritance

**Simple idea:** Inheritance can take five shapes: single, multilevel, hierarchical, multiple, and hybrid. Go supports none of these in the classic sense.

Picture a family tree, that covers all five shapes:

1. **Single**: one parent, one child. `Bike extends Vehicle`.
2. **Multilevel**: a chain, grandparent to parent to child. `C extends B extends A`.
3. **Hierarchical**: one parent, many children. `Car` and `Bike` both extend `Vehicle`.
4. **Multiple**: one child, many parents. C++ allows this. Java and Go do not, they use interfaces instead, mainly to avoid the problem explained in the next section.
5. **Hybrid**: a mix of two or more shapes above.

**C++ examples**
```cpp
// 1. Single: one parent, one child
class Vehicle {};
class Bike : public Vehicle {};

// 2. Multilevel: a chain of three generations
class A {};
class B : public A {};   // B inherits from A
class C : public B {};   // C inherits from B, so C also gets A's stuff

// 3. Hierarchical: two children share the same parent
class Vehicle2 {};
class Car : public Vehicle2 {};
class Truck : public Vehicle2 {};

// 4. Multiple: one child, two separate parents
class Flyable { public: void fly() {} };
class Swimmable { public: void swim() {} };
class FlyingBoat : public Flyable, public Swimmable {}; // inherits from both parents at once

// 5. Hybrid: multilevel mixed with hierarchical
class Base {};
class Mid : public Base {};   // multilevel part
class Other {};
class Combo : public Mid, public Other {}; // multiple part on top of it
```

**Go** has no `extends` at all. Go chose not to have classic inheritance. Instead it gives you struct embedding, which can copy the feel of single, multilevel, and even "multiple" inheritance by embedding more than one struct. But under the hood this is composition, not a true parent and child type.
```go
type Vehicle struct{}
func (Vehicle) Move() { fmt.Println("moving") }

type Bike struct { Vehicle } // "single"-like: one embedded struct

type A struct{}
type B struct{ A } // "multilevel"-like: B carries A
type C struct{ B } // C carries B, which carries A

type Flyable struct{}
func (Flyable) Fly() { fmt.Println("flying") }
type Swimmable struct{}
func (Swimmable) Swim() { fmt.Println("swimming") }

type FlyingBoat struct { // "multiple"-like: embeds two structs at once
    Flyable
    Swimmable
}
```

---

## 5. The Diamond Problem

**Simple idea:** When a class inherits the same method through two different paths that both trace back to one shared ancestor, the compiler does not know which version to use.

Real-life example: imagine you have two bosses because of a matrix org structure, and both report to the same CEO. If the CEO sends a company policy, and both your bosses forward it to you with slightly different edits, which version do you follow? That confusion, two paths, one shared origin, conflicting versions arriving at the bottom, is the diamond problem.

The class shape literally looks like a diamond:
```
        A
       / \
      B   C
       \ /
        D
```

`D` inherits from both `B` and `C`, and both `B` and `C` inherit from `A`. If `A` has `abc()`, does `D` get `B`'s copy or `C`'s copy?

Why it matters: this ambiguous method leads to compile errors, or unpredictable behavior in languages that allow it. It happens only in languages that support multiple inheritance of implementation, mainly C++. It is solved with `virtual` inheritance in C++, which tells the compiler to keep only one shared copy of `A`. Java and Go avoid the whole problem by not allowing multiple inheritance of classes. They only allow implementing multiple interfaces, which have no state to conflict over.

**C++, the problem**
```cpp
class A { public: void abc() { cout << "A::abc" << endl; } };
class B : public A {};   // B gets its own copy of A
class C : public A {};   // C gets its own separate copy of A
class D : public B, public C {}; // D now has TWO copies of A, one via B and one via C

int main() {
    D obj;
    // obj.abc(); // ERROR: ambiguous, compiler cannot tell B::abc from C::abc
}
```

**C++, the fix (virtual inheritance)**
```cpp
class A { public: void abc() { cout << "A::abc" << endl; } };
class B : virtual public A {};   // "virtual" tells the compiler: share A, do not duplicate it
class C : virtual public A {};   // same here, share the single A
class D : public B, public C {}; // now D has only ONE shared copy of A

int main() {
    D obj;
    obj.abc(); // Works now, no ambiguity, only one copy of A exists
}
```

**Go**, this problem simply does not exist, because Go has no multiple inheritance of implementation. The closest thing is embedding two structs that both have a method with the same name. Go will not guess for you, it forces you to be clear.
```go
type A struct{}
func (A) Abc() { fmt.Println("A.Abc") }

type B struct{ A } // B embeds its own copy of A
type C struct{ A } // C embeds its own separate copy of A

type D struct {
    B
    C
}

func main() {
    d := D{}
    // d.Abc()      // Compile error: ambiguous selector d.Abc (which path, B.A or C.A?)
    d.B.A.Abc()      // must spell out the exact path, Go refuses to guess
}
```

---

## 6. Method Overloading (Compile-time Polymorphism)

**Simple idea:** Multiple methods with the same name but different parameter lists. The compiler decides which one to call before the program even runs.

Real-life example: think of the word "cut," as in "cut the cake," "cut the grass," "cut a deal." Same word, but the meaning depends on what you give it. That is overloading. Same method name, but the shape of the input decides which version runs, and that choice is made at compile time.

Why and when: it lets you offer several ways to call the same idea, like `add(int,int)` and `add(double,double)`, without inventing awkward names like `addInts` and `addDoubles`. Use it when an operation is really the same idea but needs to accept different types or numbers of arguments. In C++, you just define multiple methods with the same name but different parameters.

**C++**
```cpp
class Calculator {
public:
    int add(int a, int b) { return a + b; }               // version 1: two ints
    double add(double a, double b) { return a + b; }       // version 2: two doubles
    int add(int a, int b, int c) { return a + b + c; }     // version 3: three ints
};

int main() {
    Calculator calc;
    cout << calc.add(2, 3) << endl;        // matches version 1 (int, int)
    cout << calc.add(2.5, 3.5) << endl;    // matches version 2 (double, double)
    cout << calc.add(1, 2, 3) << endl;     // matches version 3 (int, int, int)
}
```

**Go does not support method or function overloading at all.** Two functions in the same package cannot share a name, no matter how different their parameters are. If you try, it is a compile error. Here is how Go developers work around it:

```go
// This would NOT compile in Go:
// func Add(a, b int) int { return a + b }
// func Add(a, b float64) float64 { return a + b }   // ERROR: Add redeclared in this block

// Option 1: give each version its own clear name
func AddInts(a, b int) int          { return a + b }
func AddFloats(a, b float64) float64 { return a + b }

// Option 2: use variadic args plus generics (Go 1.18+) to cover different counts of the same type
func AddAll[T int | float64](nums ...T) T { // T can be int OR float64, decided by the caller
    var total T
    for _, n := range nums { // loop through however many numbers were passed in
        total += n
    }
    return total
}

func main() {
    fmt.Println(AddInts(2, 3))       // 5
    fmt.Println(AddFloats(2.5, 3.5)) // 6.0
    fmt.Println(AddAll(1, 2, 3))     // works with ints, T becomes int
    fmt.Println(AddAll(1.1, 2.2))    // works with floats too, T becomes float64
}
```

---

## 7. Method Overriding (Runtime Polymorphism)

**Simple idea:** A child class replaces a parent's method with its own version. The correct version is picked at runtime, based on the actual object type.

Real-life example: every `Animal` has a general `makeSound()`, but a `Dog` overrides it to bark, a `Cat` overrides it to meow. You do not decide which one runs when you write the code. The program decides while it is running, by checking what object you actually have.

Why and when: this lets a general piece of code, like a loop over a list of `Animal`s calling `makeSound()`, automatically get the right specific behavior for each object, without a long if/else chain checking types. In C++, mark the base method `virtual`, and override it in the child with the same signature. In Go, there is no `override` keyword and no inheritance to override from. Instead you define an interface, and each concrete struct provides its own version of that method.

**C++**
```cpp
class Animal {
public:
    virtual void sound() { cout << "Animal makes a sound" << endl; } // "virtual" is the key word here
};
class Dog : public Animal {
public:
    void sound() override { cout << "Dog barks" << endl; }   // replaces Animal's version for Dog objects
};
class Cat : public Animal {
public:
    void sound() override { cout << "Cat meows" << endl; }   // replaces Animal's version for Cat objects
};

int main() {
    Animal* a1 = new Dog();   // pointer type is Animal, but the real object is a Dog
    Animal* a2 = new Cat();   // pointer type is Animal, but the real object is a Cat
    a1->sound(); // prints "Dog barks", decided at RUNTIME by checking the real object
    a2->sound(); // prints "Cat meows", same runtime check
    delete a1; delete a2;     // free the memory we created with "new"
}
```
Note: if `sound()` in `Animal` is not marked `virtual`, C++ will use static binding and always call `Animal::sound()` through an `Animal*` pointer. This is a common bug when someone forgets the `virtual` keyword.

**Go**
```go
type Animal interface {
    Sound() // contract: any Animal must have a Sound method
}

type Dog struct{}
func (Dog) Sound() { fmt.Println("Dog barks") } // Dog's own implementation

type Cat struct{}
func (Cat) Sound() { fmt.Println("Cat meows") } // Cat's own implementation

func main() {
    animals := []Animal{Dog{}, Cat{}}  // a slice holding two different concrete types
    for _, a := range animals {
        a.Sound() // correct version picked at runtime, based on whether "a" is really a Dog or a Cat
    }
}
```

---

## 8. Compile-time vs Runtime Polymorphism

**Simple idea:** Compile-time polymorphism is decided while the code is being compiled, through overloading. Runtime polymorphism is decided while the code is actually running, through overriding.

| | Compile-time (Static) | Runtime (Dynamic) |
|---|---|---|
| Achieved by | Method overloading | Method overriding, interfaces |
| Decision made | At compile time, by the compiler looking at argument types | At runtime, by looking at the actual object's type |
| Also called | Early binding | Late binding |
| C++ mechanism | Same name, different parameter lists | `virtual` functions plus inheritance |
| Go mechanism | Not supported, Go has no overloading | Interfaces, the only form of polymorphism Go has |
| Flexibility | Fixed at compile time, less flexible | Can add new types later without touching existing code, more flexible |

**C++, both in one file**
```cpp
class Demo {
public:
    void show(int a) { cout << "int: " << a << endl; }         // resolved at COMPILE time
    void show(string s) { cout << "string: " << s << endl; }   // resolved at COMPILE time
};

class Parent { public: virtual void display() { cout << "Parent" << endl; } };
class Child : public Parent { public: void display() override { cout << "Child" << endl; } };

int main() {
    Demo d;
    d.show(10);       // compiler picks show(int) just by looking at the argument type
    d.show("hello");  // compiler picks show(string) the same way

    Parent* p = new Child();  // pointer type is Parent, real object is Child
    p->display(); // resolved at RUNTIME, prints "Child" because that is the real object
    delete p;
}
```

**Go**: since overloading does not exist, Go only really has the runtime column, always achieved through interfaces:
```go
type Shape interface{ Area() float64 } // contract: anything with an Area() method counts as a Shape

type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return 3.14159 * c.Radius * c.Radius }

type Square struct{ Side float64 }
func (s Square) Area() float64 { return s.Side * s.Side }

func printArea(s Shape) { fmt.Println("Area:", s.Area()) } // s.Area() is resolved at runtime

func main() {
    printArea(Circle{Radius: 5}) // runs Circle's Area()
    printArea(Square{Side: 4})   // runs Square's Area()
}
```

---

## 9. Interfaces vs Concrete Types

**Simple idea:** An interface is a contract that says what a type must do. A concrete type is a real implementation that says how it does it.

Real-life example: a job posting says "must be able to drive, must have a license." That is the interface, the contract. The actual person you hire, with their own driving style, is the concrete type. The job requirement is written generically on purpose, so that tomorrow you could hire someone else who also fits the requirement, without rewriting the job posting.

Why use an interface instead of a concrete type as a parameter or field: your code becomes independent of one specific implementation. You can swap Stripe for PayPal, Redis for an in-memory cache, or a mock for a real database in tests, without touching the calling code.

When to use a concrete type instead: when there really is, and will only ever be, one implementation. An interface adds indirection with no real benefit in that case. Using an interface for everything, even when only one implementation will ever exist, is called premature abstraction and it is not a good habit.

How it works: C++ gets close to "interface" with an abstract class made of pure virtual functions. Go has real interfaces, and here is a key point: a Go type satisfies an interface simply by having the right methods. There is no `implements` keyword needed.

**C++, interface-like abstract class vs concrete class**
```cpp
class PaymentGateway {                        // the contract (interface-like abstract class)
public:
    virtual bool charge(double amount) = 0;   // pure virtual, no body, every subclass must implement it
    virtual ~PaymentGateway() {}
};

class StripeGateway : public PaymentGateway { // one concrete implementation of the contract
public:
    bool charge(double amount) override {
        cout << "Charging $" << amount << " via Stripe" << endl;
        return true;
    }
};

void checkout(PaymentGateway* gateway, double amount) { // depends on the CONTRACT, not a specific gateway
    gateway->charge(amount);
}

int main() {
    StripeGateway stripe;
    checkout(&stripe, 49.99); // works with StripeGateway today, PayPalGateway tomorrow, no changes needed here
}
```

**Go**, the implicit match is the key point to remember here:
```go
type PaymentGateway interface { // the contract
    Charge(amount float64) bool
}

type StripeGateway struct{} // concrete type, notice: no "implements PaymentGateway" written anywhere

func (StripeGateway) Charge(amount float64) bool {
    fmt.Println("Charging via Stripe:", amount)
    return true
}

func checkout(g PaymentGateway, amount float64) { // parameter type is the interface, not StripeGateway directly
    g.Charge(amount)
}

func main() {
    checkout(StripeGateway{}, 49.99) // StripeGateway just "fits" because it has the right method, nothing more needed
}
```
A common Go saying is worth remembering: "accept interfaces, return concrete types." This means function parameters should usually be interfaces, so the code stays flexible, but return values should usually be concrete structs, so the caller gets full access to what was actually built.

---

## 10. Coupling vs Cohesion

**Simple idea:** Coupling means how much one module depends on another module's internal details. Cohesion means how focused a single module's job is. Good design means low coupling and high cohesion.

Real-life example for coupling: imagine two coworkers who can only finish their own work by constantly walking to each other's desk and going through each other's private notes. If one of them changes their filing system, the other one's work breaks too. That is tight coupling, reaching into someone else's internals means you break easily when they change. Loose coupling is the opposite. They only talk through a shared, agreed form, like an interface, so internal changes on either side do not break the other.

Real-life example for cohesion: imagine one person doing accounting, HR, marketing, and customer support all at once. That is low cohesion, too many unrelated jobs in one role. A dedicated accountant who only does accounting is high cohesion, one clear job. Classes should work like that dedicated accountant.

Why and when: low coupling and high cohesion means code that is easy to test, easy to change, and a bug fix in one place does not break other places. Ask yourself, when designing a class, does this class do one clear job, and does this class know too much about how another class works internally. To reduce coupling, use interfaces or dependency injection instead of hardcoding a specific dependency. To increase cohesion, split a class the moment it starts doing two unrelated jobs, for example, separate "calculate order total" from "send order confirmation email."

**C++, tightly coupled vs loosely coupled**
```cpp
// Tightly coupled: OrderService directly creates and depends on MySQLDatabase's internals
class MySQLDatabase {
public:
    void saveOrder(string data) { cout << "Saved to MySQL: " << data << endl; }
};
class OrderService {
    MySQLDatabase db; // hardcoded dependency, this class can NEVER swap the database without editing itself
public:
    void placeOrder(string data) { db.saveOrder(data); }
};

// Loosely coupled: OrderService only knows about a contract (interface)
class Database {
public:
    virtual void save(string data) = 0;   // the contract, any database must implement save()
    virtual ~Database() {}
};
class MySQLDatabase2 : public Database {
public:
    void save(string data) override { cout << "Saved to MySQL: " << data << endl; }
};
class OrderService2 {
    Database* db; // depends only on the ABSTRACTION, does not know or care which database it really is
public:
    OrderService2(Database* d) : db(d) {}    // the actual database is "injected" from outside
    void placeOrder(string data) { db->save(data); }
};
```

**Go, same idea**
```go
type Database interface {
    Save(data string) // contract, any storage type must implement Save
}

type MySQLDatabase struct{}
func (MySQLDatabase) Save(data string) { fmt.Println("Saved to MySQL:", data) }

type OrderService struct {
    db Database // depends on the interface, not a specific database
}

func NewOrderService(db Database) OrderService { return OrderService{db: db} } // db is passed in from outside
func (s OrderService) PlaceOrder(data string) { s.db.Save(data) }

func main() {
    service := NewOrderService(MySQLDatabase{}) // swap this for a MockDatabase in tests, nothing else changes
    service.PlaceOrder("order-123")
}
```

---

## 11. Composition over Inheritance

**Simple idea:** Prefer building a class by combining smaller objects, a "has-a" relationship, instead of extending a parent class, an "is-a" relationship, unless the is-a relationship is genuinely true.

Real-life example: think about building a car. You do not "inherit" the engine, the engine is not a type of car. You have an engine inside the car, as a separate, swappable part. If tomorrow you want a diesel engine instead of petrol, you just plug in a different engine, you do not need a whole new "kind of car" class family. That is composition, building complex objects out of smaller, independent, swappable pieces.

Why and when: deep inheritance chains get brittle fast, a change in a grandparent class can quietly break several generations of children. Composition keeps pieces independent and swappable. Prefer inheritance only when the relationship is a true, stable "is-a," like a `Car` really is a `Vehicle` and always will be. If you are inheriting just to reuse some code, and the relationship is not really "is-a," use composition instead.

**C++, composition (has-a)**
```cpp
class Engine {
public:
    int horsepower;
    Engine(int hp) : horsepower(hp) {}
    void start() { cout << "Engine started with " << horsepower << " HP" << endl; }
};

class Car {
    Engine engine; // Car HAS-A Engine, this is composition, not "Car extends Engine"
public:
    Car(int hp) : engine(hp) {}   // Car builds its own Engine as a member field
    void drive() {
        engine.start();           // Car uses the Engine's behavior through composition
        cout << "Car is driving..." << endl;
    }
};

int main() {
    Car car(120);
    car.drive();
}
```

**Go**'s whole design leans on this idea. A well known Go saying is "favor composition over inheritance." Since Go has no inheritance at all, embedding is used for both is-a-like reuse (section 4) and true has-a composition, depending on how you use it.
```go
type Engine struct{ Horsepower int }
func (e Engine) Start() { fmt.Println("Engine started with", e.Horsepower, "HP") }

type Car struct {
    engine Engine // named field (not embedded), makes the "Car HAS-A Engine" relationship explicit
}

func NewCar(hp int) Car { return Car{engine: Engine{Horsepower: hp}} } // Car builds its own Engine
func (c Car) Drive() {
    c.engine.Start()          // Car uses the Engine's behavior, but is not "an" Engine
    fmt.Println("Car is driving...")
}

func main() {
    car := NewCar(120)
    car.Drive()
}
```
Tip: in Go, using a named field like `engine Engine` shows "has-a" composition clearly. Using embedding like just `Engine`, without a field name, shows behavior reuse where methods get promoted, similar to inheritance. Knowing this difference, and choosing on purpose, shows real understanding of Go instead of copying Java habits into it.

---

## 12. Association vs Aggregation vs Composition

This topic was not in your original list, but it usually comes up right after composition, so it is worth having here. It is really just composition's family of relationship types, from weakest to strongest ownership.

| Relationship | Strength | Real-life example | Key trait |
|---|---|---|---|
| Association | Weakest, "knows-about" | A developer and a GitHub repo | Both exist fully on their own, neither owns the other |
| Aggregation | Medium, "has-a," but parts can outlive the whole | A team and its microservices | Team is disbanded, microservices still exist, just get reassigned |
| Composition | Strongest, "has-a," parts die with the whole | An order and its line items | Order is deleted, line items are deleted with it, they have no meaning on their own |

Quick test: if I destroy the whole, do the parts still make sense on their own? Yes means association or aggregation. No means composition.

**C++, all three side by side**
```cpp
// Association: Developer and Repository both exist independently, just reference each other
class Repository { public: string name; };
class Developer {
public:
    string username;
    vector<Repository*> repos; // just holds references, does not own or create them
};

// Aggregation: Team "has" Microservices, but services are created outside and can outlive the team
class Microservice { public: string name; };
class Team {
public:
    vector<Microservice*> services; // passed in from outside, Team does not create them itself
    void addService(Microservice* m) { services.push_back(m); }
};

// Composition: Order OWNS its LineItems, they are created and destroyed together with the Order
class LineItem { public: string product; int qty; };
class Order {
    vector<LineItem> items; // Order creates its own line items internally
public:
    void addItem(string product, int qty) {
        items.push_back(LineItem{product, qty}); // created inside Order, cannot exist without it
    }
    // when an Order object is destroyed, its "items" vector and all LineItems go with it automatically
};
```

**Go, all three side by side**
```go
// Association: both sides exist independently, just hold references to each other
type Repository struct{ Name string }
type Developer struct {
    Username string
    Repos    []*Repository // just references, Developer does not own these
}

// Aggregation: Team holds Microservices, but they are created outside and passed in
type Microservice struct{ Name string }
type Team struct {
    Services []*Microservice // assigned from outside, can be reassigned to another team later
}
func (t *Team) AddService(m *Microservice) { t.Services = append(t.Services, m) }

// Composition: Order creates and owns its LineItems directly, they have no life outside the Order
type LineItem struct {
    Product string
    Qty     int
}
type Order struct {
    items []LineItem // created and destroyed together with the Order itself
}
func (o *Order) AddItem(product string, qty int) {
    o.items = append(o.items, LineItem{Product: product, Qty: qty}) // built inside Order, owned by it
}
```

---

## 13. Cheat-sheet Table

One table to glance at before you recap:

| Concept | Core idea | C++ mechanism | Go mechanism |
|---|---|---|---|
| Class/Object | Blueprint vs instance | `class` | `struct` plus receiver methods |
| Encapsulation | Hide data | `private`/`protected` | lowercase (unexported) fields |
| Abstraction | Hide complexity | Abstract class (pure virtual) | `interface` |
| Inheritance | Reuse via is-a | `class B : public A` | No true inheritance, struct embedding |
| Polymorphism | One call, many behaviors | Overloading plus `virtual`/override | Interfaces only, no overloading |
| Access modifiers | Control visibility | `public`/`protected`/`private` | Capitalized (exported) vs lowercase |
| Diamond problem | Ambiguous multi-inheritance | Fixed via `virtual` inheritance | Does not occur, no multiple class inheritance |
| Method overloading | Same name, different params | Supported natively | Not supported, use distinct names or generics |
| Method overriding | Child replaces parent behavior | `virtual` plus `override` | Interface implementations |
| Interfaces vs concrete | Contract vs implementation | Abstract class vs concrete class | Implicit `interface` vs `struct` |
| Coupling/Cohesion | Independence vs focus | Program to abstract base class | Program to `interface` |
| Composition over inheritance | Has-a over is-a | Member object field | Named field (has-a) vs embedding (is-a-like) |

---


## 13. What's Next?

Now we have covered everything from OOP. Next, let’s move to FAQ.md to get familiar with common interview questions, real-life scenarios, and understand the differences between various OOP concepts.