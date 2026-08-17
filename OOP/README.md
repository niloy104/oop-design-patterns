# Object-Oriented Programming Notes (C++ and Go)

This note explains OOP in simple words. Each topic has one real-life example, and the code below it uses that same example. So you don't have to jump between a story and unrelated code, they match.

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

**Simple idea:** A class is a plan. An object is the real thing made from that plan.

**Story:** Think of a house plan on paper. The plan says how many rooms, how many doors, how many windows. You cannot sleep inside a paper plan. Someone has to build the house first. That built house is the object. You can build many houses from the same plan. Each house is separate. Each has its own furniture. But they all follow the same plan.

A class holds two things: data (fields) and behavior (methods). Making the plan does not use much memory. Building a real house (an object) is what uses memory.

**Why:** without a plan, every house needs its own new design. That is slow and messy.
**When:** any time you model something real that has data and behavior, like a `User` or an `Order`.

Go has no `class` keyword. It uses `struct` for the plan, and normal functions attached to the struct for behavior.

**C++**
```cpp
#include <iostream>
#include <string>
using namespace std;

class House {                              // the plan
public:
    int rooms;
    int doors;

    House(int r, int d) : rooms(r), doors(d) {}   // this runs when we build a real house

    void describe() {
        cout << "House with " << rooms << " rooms and " << doors << " doors" << endl;
    }
};

int main() {
    House h1(3, 2);     // h1 is a real house built from the plan
    House h2(5, 4);     // h2 is another, separate house, same plan, different numbers
    h1.describe();
    h2.describe();
}
```

**Go**
```go
package main

import "fmt"

type House struct { // the plan
    Rooms int
    Doors int
}

func NewHouse(rooms, doors int) House { // builds a real house from the plan
    return House{Rooms: rooms, Doors: doors}
}

func (h House) Describe() {
    fmt.Println("House with", h.Rooms, "rooms and", h.Doors, "doors")
}

func main() {
    h1 := NewHouse(3, 2) // a real house
    h2 := NewHouse(5, 4) // another, separate house
    h1.Describe()
    h2.Describe()
}
```

---

## 2. The Four Pillars

The four pillars of OOP are Encapsulation, Abstraction, Inheritance, and Polymorphism.

### 2.1 Encapsulation

**Simple idea:** Keep data private. Let other code touch it only through a method you allow.

**Story:** Think of a bank account. You cannot open your account file and change your balance to a million dollars by hand. You can only ask the bank to `deposit` or `withdraw`, and the bank checks the rules for you.

**Why:** this stops your data from getting a wrong value, like a negative balance.
**When:** any time a field can go wrong if touched directly.

**C++**
```cpp
class Account {
private:
    double balance;                          // hidden, nobody outside can touch this directly
public:
    Account(double initial) : balance(initial) {}
    double getBalance() const { return balance; }   // safe, read-only way to check balance
    void deposit(double amount) {
        if (amount > 0) balance += amount;   // the rule is checked here, inside the class
    }
};
```

**Go**
```go
type Account struct {
    balance float64 // lowercase name, so it is private to this package
}

func NewAccount(initial float64) *Account { return &Account{balance: initial} }

func (a *Account) Balance() float64 { return a.balance }
func (a *Account) Deposit(amount float64) {
    if amount > 0 { // rule checked here, caller cannot skip it
        a.balance += amount
    }
}
```

### 2.2 Abstraction

**Simple idea:** Show only what someone needs. Hide the messy details of how it works.

**Story:** Think of a coffee machine. You put in water and coffee beans. You press one button. You get coffee. You do not see the heater, the pump, or the timer inside. All that is hidden from you.

**Why:** this lets you change the inside later (a better pump, a new heater) without changing how people use the button.
**When:** any time more than one version of "the same idea" might exist now or later.

C++ does this with an abstract class (a class with methods that have no body). Go does this with an `interface`. In Go, a type does not need to say "I am a CoffeeMachine." It just needs the right method, and it counts automatically.

**C++**
```cpp
class CoffeeMachine {                          // the button, the contract
public:
    virtual string makeCoffee() = 0;          // no body here, just a promise
    virtual ~CoffeeMachine() {}
};

class EspressoMachine : public CoffeeMachine { // the real machine, with real steps hidden inside
public:
    string makeCoffee() override {
        return "Heating water, grinding beans, pouring espresso"; // the messy part, hidden from the user
    }
};
```

**Go**
```go
type CoffeeMachine interface { // the button, the contract
    MakeCoffee() string
}

type EspressoMachine struct{} // the real machine

func (EspressoMachine) MakeCoffee() string {
    return "Heating water, grinding beans, pouring espresso" // messy part hidden here
}
```

### 2.3 Inheritance

**Simple idea:** A child class reuses what a parent class already has, and can add more. This is an "is-a" relationship.

**Story:** A parent passes traits down to a child. A `Vehicle` can move. A `Bike` is a kind of `Vehicle`, so it can move too, without us writing that code again. A `Bike` also has its own thing, like ringing a bell.

**Why:** avoids writing the same field or method twice for related things.
**When:** only when the relationship is a real "is-a," like a `Bike` really is a `Vehicle`.

C++ uses `class Child : public Parent`. Go has no class inheritance. It uses struct embedding, which feels similar.

**C++**
```cpp
class Vehicle {
public:
    void move() { cout << "Vehicle is moving" << endl; }  // shared behavior
};

class Bike : public Vehicle {     // Bike IS-A Vehicle
public:
    void ringBell() { cout << "Bike rings the bell" << endl; }  // Bike's own extra behavior
};

int main() {
    Bike b;
    b.move();      // came from Vehicle, Bike did not write this again
    b.ringBell();  // Bike's own method
}
```

**Go**
```go
type Vehicle struct{}

func (Vehicle) Move() { fmt.Println("Vehicle is moving") } // shared behavior

type Bike struct {
    Vehicle // embedded, Bike gets Vehicle's method for free
}

func (Bike) RingBell() { fmt.Println("Bike rings the bell") } // Bike's own extra behavior

func main() {
    b := Bike{}
    b.Move()     // came from the embedded Vehicle
    b.RingBell() // Bike's own method
}
```

### 2.4 Polymorphism

**Simple idea:** One method name, many behaviors, depending on the real object.

**Story:** A `Dog` and a `Cat` both `speak()`. But a dog barks and a cat meows. Same method name, different sound, because the object is different.

**Why:** you can write one piece of code that works for many kinds of objects, without checking "if this is a dog, if this is a cat."
**When:** any time related objects need their own version of the same action.

Polymorphism has two types: compile-time and runtime. Full detail is in sections 6, 7, and 8. Here is the quick version:

**C++**
```cpp
class Animal { public: virtual void speak() { cout << "..." << endl; } };
class Dog : public Animal { public: void speak() override { cout << "Bark" << endl; } };
class Cat : public Animal { public: void speak() override { cout << "Meow" << endl; } };

void makeItSpeak(Animal* a) { a->speak(); }  // does not care if it is a Dog or a Cat

int main() {
    Dog d; Cat c;
    makeItSpeak(&d);  // Bark
    makeItSpeak(&c);  // Meow
}
```

**Go**
```go
type Animal interface{ Speak() } // any type with a Speak() method counts as an Animal

type Dog struct{}
func (Dog) Speak() { fmt.Println("Bark") }

type Cat struct{}
func (Cat) Speak() { fmt.Println("Meow") }

func makeItSpeak(a Animal) { a.Speak() } // works with any Animal, Dog or Cat or anything else

func main() {
    makeItSpeak(Dog{}) // Bark
    makeItSpeak(Cat{}) // Meow
}
```

---

## 3. Access Modifiers

**Simple idea:** Access modifiers decide who is allowed to see or use a field or method.

**Story:** Think of an employee's record card. The name is on a public badge, anyone can see it. The role is only shown to managers. The salary is private, only that person and payroll can see it.

- **Public**: name, anyone can see it.
- **Protected**: role, only this class and its children can see it.
- **Private**: salary, only this class can see it.

Go has no keywords for this. It uses one simple rule: capital letter at the start means visible outside the package, small letter means hidden.

| | C++ | Go |
|---|---|---|
| Public | `public:` | Capital letter, like `Name` |
| Protected | `protected:` | No matching feature |
| Private | `private:` | Small letter, like `salary` |
| Shared by all objects | `static` | Package-level `var` |

**C++**
```cpp
class Employee {
public:
    string name;          // like the badge, anyone can see it
protected:
    string role;           // like info shown only to managers
private:
    double salary;          // like info seen only by payroll
public:
    static int companyCode; // one shared value for every Employee, not per-object
};
```

**Go**
```go
var CompanyCode int // shared value, closest thing Go has to "static"

type Employee struct {
    Name   string  // capital N, like the public badge
    role   string  // small r, hidden outside this package
    salary float64 // small s, hidden outside this package
}
```

---

## 4. Types of Inheritance

**Simple idea:** Inheritance can happen in five shapes: single, multilevel, hierarchical, multiple, and hybrid. Go does not support any of these the classic way.

**Story:** Think of a family tree, that covers all five shapes.

1. **Single**: one parent, one child. `Bike` from `Vehicle`.
2. **Multilevel**: a chain, grandparent to parent to child.
3. **Hierarchical**: one parent, many children. `Car` and `Truck` both from `Vehicle`.
4. **Multiple**: one child, many parents. C++ allows this. Java and Go do not, they use interfaces instead.
5. **Hybrid**: a mix of the shapes above.

**C++ examples**
```cpp
// 1. Single: one parent, one child
class Vehicle {};
class Bike : public Vehicle {};

// 2. Multilevel: grandparent, parent, child
class A {};
class B : public A {};   // parent
class C : public B {};   // child, also gets A's stuff through B

// 3. Hierarchical: two children, one shared parent
class Vehicle2 {};
class Car : public Vehicle2 {};
class Truck : public Vehicle2 {};

// 4. Multiple: one child, two parents
class Flyable { public: void fly() {} };
class Swimmable { public: void swim() {} };
class FlyingBoat : public Flyable, public Swimmable {}; // gets both parents' stuff

// 5. Hybrid: shape 2 mixed with shape 4
class Base {};
class Mid : public Base {};   // multilevel part
class Other {};
class Combo : public Mid, public Other {}; // multiple part added on top
```

**Go** has no `extends`. Instead it uses struct embedding, which can copy the feel of shapes 1, 2, and 4. But under the hood it is composition, not real inheritance.
```go
type Vehicle struct{}
func (Vehicle) Move() { fmt.Println("moving") }

type Bike struct { Vehicle } // shape 1 feel

type A struct{}
type B struct{ A } // shape 2 feel
type C struct{ B }

type Flyable struct{}
func (Flyable) Fly() { fmt.Println("flying") }
type Swimmable struct{}
func (Swimmable) Swim() { fmt.Println("swimming") }

type FlyingBoat struct { // shape 4 feel, two embedded structs
    Flyable
    Swimmable
}
```

---

## 5. The Diamond Problem

**Simple idea:** If a class gets the same method through two different parents, and both parents came from one shared grandparent, the compiler gets confused about which copy to use.

**Story:** Imagine you have two managers, and both of them report to the same CEO. The CEO sends out one company rule. Each manager passes it to you, but with small edits. Now you have two different versions of the same rule. Which one do you follow? That confusion is the diamond problem.

The shape looks like a diamond:
```
        CEO
       /    \
  Manager1  Manager2
       \    /
        You
```

**Why it matters:** the compiler cannot pick a side on its own, so it either refuses to compile, or picks unpredictably.
**When it happens:** only in languages that allow multiple inheritance, mainly C++.
**How it is fixed:** C++ uses `virtual` inheritance, which tells the compiler to keep only one shared copy of the CEO's rule. Go and Java skip the whole problem, they do not allow multiple inheritance of classes at all.

**C++, the problem**
```cpp
class CEO { public: void rule() { cout << "CEO::rule" << endl; } };
class Manager1 : public CEO {};   // gets its own copy of CEO
class Manager2 : public CEO {};   // gets a separate copy of CEO
class You : public Manager1, public Manager2 {}; // now has TWO copies of CEO's rule

int main() {
    You obj;
    // obj.rule(); // ERROR: ambiguous, compiler cannot tell Manager1's copy from Manager2's copy
}
```

**C++, the fix (virtual inheritance)**
```cpp
class CEO { public: void rule() { cout << "CEO::rule" << endl; } };
class Manager1 : virtual public CEO {};   // "virtual" means: share CEO, do not copy it
class Manager2 : virtual public CEO {};   // same here
class You : public Manager1, public Manager2 {}; // now only ONE shared copy of CEO exists

int main() {
    You obj;
    obj.rule(); // Works, no confusion, only one copy of the rule
}
```

**Go**, this problem does not happen, because Go has no multiple inheritance. The closest thing is embedding two structs that both have the same method name. Go refuses to guess, and forces you to be exact.
```go
type CEO struct{}
func (CEO) Rule() { fmt.Println("CEO.Rule") }

type Manager1 struct{ CEO } // own copy of CEO
type Manager2 struct{ CEO } // separate copy of CEO

type You struct {
    Manager1
    Manager2
}

func main() {
    y := You{}
    // y.Rule()             // Compile error: ambiguous, which path, Manager1 or Manager2?
    y.Manager1.CEO.Rule()   // must spell out the exact path
}
```

---

## 6. Method Overloading (Compile-time Polymorphism)

**Simple idea:** Same method name, different inputs. The compiler decides which one to run before the program even starts.

**Story:** Think of a calculator with an `add` button. It can add two whole numbers, or two numbers with decimals, or even three numbers at once. Same button name, but which version runs depends on what you typed in.

**Why:** lets you use one simple name for the same idea, instead of `addTwoInts`, `addTwoDoubles`, `addThreeInts`.
**When:** the action is really the same idea, just with different input types or counts.

**C++**
```cpp
class Calculator {
public:
    int add(int a, int b) { return a + b; }               // version 1: two whole numbers
    double add(double a, double b) { return a + b; }       // version 2: two decimal numbers
    int add(int a, int b, int c) { return a + b + c; }     // version 3: three whole numbers
};

int main() {
    Calculator calc;
    cout << calc.add(2, 3) << endl;        // matches version 1
    cout << calc.add(2.5, 3.5) << endl;    // matches version 2
    cout << calc.add(1, 2, 3) << endl;     // matches version 3
}
```

**Go does not allow this at all.** Two functions in the same package cannot share a name, no matter how different their inputs are. This is a real difference to remember. Here is how Go works around it:

```go
// This would NOT compile in Go:
// func Add(a, b int) int { return a + b }
// func Add(a, b float64) float64 { return a + b }   // ERROR: Add already declared

// Option 1: give each version a clear, separate name
func AddInts(a, b int) int          { return a + b }
func AddFloats(a, b float64) float64 { return a + b }

// Option 2: use generics (Go 1.18+) plus variadic args, to cover different counts of the same type
func AddAll[T int | float64](nums ...T) T { // T can be int OR float64, Go figures it out from the input
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(AddInts(2, 3))       // 5
    fmt.Println(AddFloats(2.5, 3.5)) // 6
    fmt.Println(AddAll(1, 2, 3))     // works with ints
    fmt.Println(AddAll(1.1, 2.2))    // works with floats too
}
```

---

## 7. Method Overriding (Runtime Polymorphism)

**Simple idea:** A child class replaces a parent's method with its own version. The program picks the right version while it is running, not before.

**Story:** Every `Animal` has a general `sound()`. A `Dog` overrides it to bark. A `Cat` overrides it to meow. You don't choose which one runs when you write the code, the program checks the real object while running, and picks the right one.

**Why:** lets one loop, like "go through every animal and call sound," automatically get the correct sound for each one, without writing many if-checks.
**When:** related objects share the same method name, but each needs its own version.

C++ needs the word `virtual` on the parent method, or overriding will not work correctly. Go has no `override` keyword and no class inheritance, it uses interfaces instead, and every type just writes its own method.

**C++**
```cpp
class Animal {
public:
    virtual void sound() { cout << "Animal makes a sound" << endl; } // "virtual" is required here
};
class Dog : public Animal {
public:
    void sound() override { cout << "Dog barks" << endl; }   // replaces Animal's version
};
class Cat : public Animal {
public:
    void sound() override { cout << "Cat meows" << endl; }   // replaces Animal's version
};

int main() {
    Animal* a1 = new Dog();   // pointer type says Animal, real object is a Dog
    Animal* a2 = new Cat();   // pointer type says Animal, real object is a Cat
    a1->sound(); // "Dog barks", picked at RUNTIME by checking the real object
    a2->sound(); // "Cat meows", same check
    delete a1; delete a2;
}
```
Note: if `sound()` in `Animal` is missing the word `virtual`, C++ will always call `Animal::sound()` no matter what the real object is. This is a very common bug.

**Go**
```go
type Animal interface {
    Sound() // any type with a Sound() method counts as an Animal
}

type Dog struct{}
func (Dog) Sound() { fmt.Println("Dog barks") } // Dog's own version

type Cat struct{}
func (Cat) Sound() { fmt.Println("Cat meows") } // Cat's own version

func main() {
    animals := []Animal{Dog{}, Cat{}}  // a list holding two different real types
    for _, a := range animals {
        a.Sound() // correct version picked at runtime, based on the real type
    }
}
```

---

## 8. Compile-time vs Runtime Polymorphism

**Simple idea:** Compile-time polymorphism is method overloading, and it is decided while the code is being built. Runtime polymorphism is method overriding, and it is decided while the program is actually running.

This section reuses the Calculator example from section 6, and the Animal example from section 7, so you can see both side by side.

| | Compile-time (Static) | Runtime (Dynamic) |
|---|---|---|
| Achieved by | Method overloading | Method overriding, interfaces |
| Decided when | While the code is compiled, by looking at argument types | While the program is running, by looking at the real object |
| Also called | Early binding | Late binding |
| C++ mechanism | Same name, different parameters | `virtual` plus `override` |
| Go mechanism | Not supported | Interfaces, this is Go's only kind of polymorphism |

**C++, both together**
```cpp
Calculator calc;
cout << calc.add(2, 3) << endl;      // COMPILE-TIME: compiler picks add(int,int) just from the input types

Animal* a = new Dog();
a->sound();                          // RUNTIME: decided only when the program actually runs, based on real object
delete a;
```

**Go, only has runtime, since overloading does not exist**
```go
var animals []Animal = []Animal{Dog{}, Cat{}}
for _, a := range animals {
    a.Sound() // decided at RUNTIME every single time, based on the real type in the list
}
```

---

## 9. Interfaces vs Concrete Types

**Simple idea:** An interface is a promise of what a type can do. A concrete type is the real thing that actually does it.

**Story:** A job post says "must be able to drive." That is the interface, just a requirement. The real person you hire, with their own driving style, is the concrete type. The job post is written this way on purpose, so tomorrow you can hire someone else who can also drive, without changing the job post.

**Why use an interface:** your code does not depend on one specific person or thing. You can swap the driver later without changing anything else.
**When to use a concrete type instead:** when there will only ever be one real version, and an interface would just add extra steps for no benefit.

C++ gets close to "interface" using an abstract class with methods that have no body. In Go, a type does not need to say "I can drive." If it has a `Drive()` method, it already counts, automatically.

**C++**
```cpp
class Driver {                              // the job post, the promise
public:
    virtual void drive() = 0;              // no body, just a requirement
    virtual ~Driver() {}
};

class Person : public Driver {             // the real person hired
public:
    void drive() override {
        cout << "Person is driving the car" << endl;
    }
};

void startTrip(Driver* d) {                // depends on the JOB POST, not one specific person
    d->drive();
}

int main() {
    Person p;
    startTrip(&p); // works with Person today, could work with anyone else who can drive tomorrow
}
```

**Go**, this is the key point to remember: no keyword needed to say "I match this interface."
```go
type Driver interface { // the job post
    Drive()
}

type Person struct{} // the real hire, no "implements Driver" written anywhere

func (Person) Drive() {
    fmt.Println("Person is driving the car")
}

func startTrip(d Driver) { // depends on the job post, not one specific person
    d.Drive()
}

func main() {
    startTrip(Person{}) // Person just fits, because it has the right method
}
```
A common Go saying: "accept interfaces, return concrete types." Function inputs should usually be interfaces, so the code stays flexible. Function outputs should usually be real, concrete types, so the caller gets full access to what was built.

---

## 10. Coupling vs Cohesion

**Simple idea:** Coupling is how much one piece of code depends on another piece's private details. Cohesion is how focused one piece of code is on doing just one job. Good code has low coupling and high cohesion.

**Story for coupling:** Imagine two coworkers who can only finish their work by walking to each other's desk and going through each other's private notes. If one changes their filing system, the other one's work breaks too. That is tight coupling. Now imagine they only pass a simple form back and forth, and never touch each other's private notes. If one changes their internal system, the other one does not even notice. That is loose coupling.

**Story for cohesion:** Imagine one person doing accounting, HR, marketing, and customer support, all at once. Too many unrelated jobs, that is low cohesion. Now imagine a person who only does accounting. One clear job, that is high cohesion.

**Why:** low coupling and high cohesion means code is easier to test and change, and a fix in one place does not break another place.
**How to get it:** depend on a shared contract (interface) instead of someone else's private details, and split a class the moment it starts doing two unrelated jobs.

The code below reuses the coworker idea directly: `OrderService` is like the coworker who insists on going through `MySQLDatabase`'s private details, versus a version that only uses the shared form (the `Database` interface).

**C++, tightly coupled vs loosely coupled**
```cpp
// Tightly coupled: like a coworker who goes straight into MySQLDatabase's private notes
class MySQLDatabase {
public:
    void saveOrder(string data) { cout << "Saved to MySQL: " << data << endl; }
};
class OrderService {
    MySQLDatabase db; // hardcoded, this class can never swap the database without editing itself
public:
    void placeOrder(string data) { db.saveOrder(data); }
};

// Loosely coupled: like a coworker who only uses a shared form (the Database contract)
class Database {
public:
    virtual void save(string data) = 0;   // the shared form
    virtual ~Database() {}
};
class MySQLDatabase2 : public Database {
public:
    void save(string data) override { cout << "Saved to MySQL: " << data << endl; }
};
class OrderService2 {
    Database* db; // only knows about the shared form, not the real database's private details
public:
    OrderService2(Database* d) : db(d) {}    // the real database is handed in from outside
    void placeOrder(string data) { db->save(data); }
};
```

**Go, same idea**
```go
type Database interface {
    Save(data string) // the shared form
}

type MySQLDatabase struct{}
func (MySQLDatabase) Save(data string) { fmt.Println("Saved to MySQL:", data) }

type OrderService struct {
    db Database // only knows the shared form, not MySQLDatabase's private details
}

func NewOrderService(db Database) OrderService { return OrderService{db: db} } // handed in from outside
func (s OrderService) PlaceOrder(data string) { s.db.Save(data) }

func main() {
    service := NewOrderService(MySQLDatabase{}) // swap this for a test database later, nothing else changes
    service.PlaceOrder("order-123")
}
```

---

## 11. Composition over Inheritance

**Simple idea:** Prefer building a class by putting smaller objects inside it, a "has-a" relationship, instead of extending a parent class, an "is-a" relationship, unless "is-a" is really true.

**Story:** Think about a car. The engine is not a type of car, so the car does not "inherit" from the engine. Instead, the car has an engine inside it, as a separate, swappable part. Want a diesel engine instead of petrol? Just swap the part. You do not need a whole new "kind of car."

**Why:** long inheritance chains break easily. A change high up can quietly break many classes below it. Composition keeps parts independent and easy to swap.
**When to still use inheritance:** only when the relationship is truly "is-a" and will stay that way, like a `Car` really is a `Vehicle`.

**C++, composition (has-a)**
```cpp
class Engine {
public:
    int horsepower;
    Engine(int hp) : horsepower(hp) {}
    void start() { cout << "Engine started with " << horsepower << " HP" << endl; }
};

class Car {
    Engine engine; // Car HAS-A Engine, not "Car extends Engine"
public:
    Car(int hp) : engine(hp) {}   // Car builds its own Engine as a part
    void drive() {
        engine.start();           // Car uses the Engine's behavior through this part
        cout << "Car is driving..." << endl;
    }
};

int main() {
    Car car(120);
    car.drive();
}
```

**Go** leans hard into this idea. A well-known Go saying is "favor composition over inheritance." Since Go has no inheritance at all, embedding gets used for both inheritance-like reuse (section 4) and true has-a composition, depending on how you write it.
```go
type Engine struct{ Horsepower int }
func (e Engine) Start() { fmt.Println("Engine started with", e.Horsepower, "HP") }

type Car struct {
    engine Engine // named field, not embedded, makes "Car HAS-A Engine" clear
}

func NewCar(hp int) Car { return Car{engine: Engine{Horsepower: hp}} } // Car builds its own Engine part
func (c Car) Drive() {
    c.engine.Start()          // Car uses the Engine, but Car is not "an" Engine
    fmt.Println("Car is driving...")
}

func main() {
    car := NewCar(120)
    car.Drive()
}
```
Tip: a named field like `engine Engine` means "has-a," clearly. Plain embedding, like just writing `Engine` with no field name, means behavior reuse, closer to inheritance. Choosing between these on purpose is a good sign you understand Go, instead of just copying Java habits into it.

---

## 12. Association vs Aggregation vs Composition

This topic was not in your original list, but it fits right after composition, so it is added here. It is the family of "has-a" relationships, from weakest to strongest.

| Relationship | Strength | Story | Key trait |
|---|---|---|---|
| Association | Weakest, "knows-about" | A developer and a GitHub repo | Both exist on their own, neither owns the other |
| Aggregation | Medium, "has-a," parts can live on without the whole | A team and its microservices | Team is closed down, the microservices still exist, just get reassigned |
| Composition | Strongest, "has-a," parts die with the whole | An order and its line items | Order is deleted, its line items are deleted too, they mean nothing on their own |

Quick test: if you destroy the whole, do the parts still make sense on their own? Yes means association or aggregation. No means composition.

**C++, all three side by side**
```cpp
// Association: Developer and Repository both exist on their own, just point at each other
class Repository { public: string name; };
class Developer {
public:
    string username;
    vector<Repository*> repos; // just holds pointers, does not own or create these repos
};

// Aggregation: Team "has" Microservices, but they are made outside and can outlive the team
class Microservice { public: string name; };
class Team {
public:
    vector<Microservice*> services; // handed in from outside, Team did not create these
    void addService(Microservice* m) { services.push_back(m); }
};

// Composition: Order OWNS its LineItems, they are created and destroyed together with the Order
class LineItem { public: string product; int qty; };
class Order {
    vector<LineItem> items; // Order builds its own line items inside itself
public:
    void addItem(string product, int qty) {
        items.push_back(LineItem{product, qty}); // built inside Order, cannot exist without it
    }
    // when Order is destroyed, "items" and all its LineItems are destroyed automatically too
};
```

**Go, all three side by side**
```go
// Association: both sides exist on their own, just hold references to each other
type Repository struct{ Name string }
type Developer struct {
    Username string
    Repos    []*Repository // just references, Developer does not own these
}

// Aggregation: Team holds Microservices, but they are made outside and passed in
type Microservice struct{ Name string }
type Team struct {
    Services []*Microservice // assigned from outside, could move to another team later
}
func (t *Team) AddService(m *Microservice) { t.Services = append(t.Services, m) }

// Composition: Order creates and owns its LineItems, they have no life outside the Order
type LineItem struct {
    Product string
    Qty     int
}
type Order struct {
    items []LineItem // created and destroyed together with the Order
}
func (o *Order) AddItem(product string, qty int) {
    o.items = append(o.items, LineItem{Product: product, Qty: qty}) // built inside Order, owned by it
}
```

---

## 13. Cheat-sheet Table

One table to glance at before you recap:

| Concept | Core idea | C++ | Go |
|---|---|---|---|
| Class/Object | Plan vs real thing | `class` | `struct` plus receiver methods |
| Encapsulation | Hide data | `private`/`protected` | lowercase (unexported) fields |
| Abstraction | Hide the messy part | Abstract class (pure virtual) | `interface` |
| Inheritance | Reuse via is-a | `class B : public A` | No real inheritance, struct embedding |
| Polymorphism | One call, many behaviors | Overloading plus `virtual`/override | Interfaces only, no overloading |
| Access modifiers | Control visibility | `public`/`protected`/`private` | Capital letter vs lowercase |
| Diamond problem | Confusing multi-inheritance | Fixed with `virtual` inheritance | Cannot happen, no multiple class inheritance |
| Method overloading | Same name, different inputs | Supported | Not supported, use different names or generics |
| Method overriding | Child replaces parent's method | `virtual` plus `override` | Interface implementations |
| Interfaces vs concrete | Promise vs the real thing | Abstract class vs concrete class | `interface` vs `struct`, no keyword needed |
| Coupling/Cohesion | Independence vs focus | Depend on abstract base class | Depend on `interface` |
| Composition over inheritance | Has-a over is-a | Object as a member field | Named field (has-a) vs embedding (is-a-like) |

---

## 13. What's Next?

Now we have covered everything from OOP. Next, let’s move to FAQ.md to get familiar with common interview questions, real-life scenarios, and understand the differences between various OOP concepts.

---

## Sources
This note is built from these four links. All the topics and explanations come from them:
1. freeCodeCamp: Object-Oriented Programming Concepts, https://medium.com/free-code-camp/object-oriented-programming-concepts-21bb035f7260
2. OOP in Go, https://medium.com/@amren1254/oops-in-go-36b40a8d1b4d
3. AlgoMaster: 12 OOP Concepts Every Developer Should Know, https://blog.algomaster.io/p/12-oop-concepts-every-developer-should-know
