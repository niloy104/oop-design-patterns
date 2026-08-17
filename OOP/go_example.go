// =====================================================================================
//  OOP NOTES - ONE FULL GO PROGRAM
//  Every topic from the note is demonstrated here, in order, with heavy comments.
//  Run:  go run oop_notes.go
// =====================================================================================

package main

import "fmt"

// =====================================================================================
// 1. OBJECT AND CLASS
// Story: a house PLAN (struct) vs a real built HOUSE (object/value).
// Go has no "class" keyword, a struct plus methods is Go's version of a class.
// =====================================================================================
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

func demo1ObjectAndClass() {
	fmt.Println("\n=== 1. Object and Class ===")
	h1 := NewHouse(3, 2) // h1 is a real house built from the plan
	h2 := NewHouse(5, 4) // h2 is a separate house, same plan, different numbers
	h1.Describe()
	h2.Describe()
}

// =====================================================================================
// 2.1 ENCAPSULATION
// Story: a bank account. You cannot hand-edit your balance, only Deposit().
// Go has no private/public keywords, lowercase = hidden outside the package.
// =====================================================================================
type Account struct {
	balance float64 // lowercase, so it's private to this package
}

func NewAccount(initial float64) *Account { return &Account{balance: initial} }

func (a *Account) Balance() float64 { return a.balance } // safe, read-only way to check balance
func (a *Account) Deposit(amount float64) {
	if amount > 0 { // the rule lives here, inside the method
		a.balance += amount
	}
}

func demo2_1Encapsulation() {
	fmt.Println("\n=== 2.1 Encapsulation ===")
	acc := NewAccount(1000)
	acc.Deposit(500) // only allowed way to change balance
	fmt.Println("Balance:", acc.Balance())
}

// =====================================================================================
// 2.2 ABSTRACTION
// Story: a coffee machine. Press one button, the messy steps inside are hidden.
// Go uses an interface for this, no "implements" keyword needed.
// =====================================================================================
type CoffeeMachine interface { // the button, the contract
	MakeCoffee() string
}

type EspressoMachine struct{} // the real machine, messy steps hidden inside

func (EspressoMachine) MakeCoffee() string {
	return "Heating water, grinding beans, pouring espresso" // hidden from the user
}

func demo2_2Abstraction() {
	fmt.Println("\n=== 2.2 Abstraction ===")
	var machine CoffeeMachine = EspressoMachine{}
	fmt.Println(machine.MakeCoffee()) // caller only sees the result, not the steps
}

// =====================================================================================
// 2.3 INHERITANCE-LIKE REUSE (also used again for 4.1 "single inheritance")
// Story: Vehicle can move, Bike embeds Vehicle so it can move too, plus its own trick.
// Go has no real inheritance, struct embedding is the closest thing.
// =====================================================================================
type Vehicle struct{}

func (Vehicle) Move() { fmt.Println("Vehicle is moving") } // shared behavior

type Bike struct {
	Vehicle // embedded, Bike gets Vehicle's method "for free" (promoted)
}

func (Bike) RingBell() { fmt.Println("Bike rings the bell") } // Bike's own extra behavior

func demo2_3Inheritance() {
	fmt.Println("\n=== 2.3 Inheritance-like reuse (also = 4.1 Single) ===")
	b := Bike{}
	b.Move()     // came from the embedded Vehicle, Bike did not rewrite this
	b.RingBell() // Bike's own method
}

// =====================================================================================
// 2.4 POLYMORPHISM  (Dog / Cat is reused later in section 7 and 8 too)
// Story: Dog and Cat both Speak(), but each sounds different.
// =====================================================================================
type Animal interface{ Speak() } // any type with a Speak() method counts as an Animal

type Dog struct{}

func (Dog) Speak() { fmt.Println("Bark") }

type Cat struct{}

func (Cat) Speak() { fmt.Println("Meow") }

func makeItSpeak(a Animal) { a.Speak() } // does not care if it's a Dog or a Cat

func demo2_4Polymorphism() {
	fmt.Println("\n=== 2.4 Polymorphism ===")
	makeItSpeak(Dog{}) // Bark
	makeItSpeak(Cat{}) // Meow
}

// =====================================================================================
// 3. ACCESS MODIFIERS (Go's version: capital letter = exported, lowercase = hidden)
// Story: employee badge (exported name), payroll-only info (unexported salary).
// Go has no separate "protected" level, only exported vs package-private.
// =====================================================================================
var CompanyCode int = 7 // shared value, closest thing Go has to C++'s "static"

type Employee struct {
	Name   string  // capital N, like the public badge, visible outside the package
	role   string  // lowercase r, hidden outside this package
	salary float64 // lowercase s, hidden outside this package
}

func NewEmployee(name, role string, salary float64) Employee {
	return Employee{Name: name, role: role, salary: salary}
}

func (e Employee) Salary() float64 { return e.salary } // controlled access to the hidden field

func demo3AccessModifiers() {
	fmt.Println("\n=== 3. Access Modifiers ===")
	e := NewEmployee("Abul", "Engineer", 50000)
	fmt.Println(e.Name, "| code", CompanyCode, "| salary via method:", e.Salary())
	// e.salary would NOT compile from another package, it's unexported
}

// =====================================================================================
// 4. TYPES OF INHERITANCE (Go only has embedding, so these are all "inheritance-like")
// 4.1 Single = Vehicle/Bike above (reused, not repeated).
// =====================================================================================

// 4.2 Multilevel: grandparent -> parent -> child
type AGrandparent struct{}

func (AGrandparent) Greet() { fmt.Println("Hello from Grandparent A") }

type BParent struct{ AGrandparent } // parent
type CChild struct{ BParent }       // child, also gets A's stuff through B

// 4.3 Hierarchical: two children share one parent (Car renamed to Sedan to avoid a name
// clash with the Car used later in the Composition section)
type Vehicle2 struct{}
type Sedan struct{ Vehicle2 }
type Truck struct{ Vehicle2 }

// 4.4 Multiple: one struct embeds two other structs at once
type Flyable struct{}

func (Flyable) Fly() { fmt.Println("Flying") }

type Swimmable struct{}

func (Swimmable) Swim() { fmt.Println("Swimming") }

type FlyingBoat struct { // gets behavior from BOTH embedded structs
	Flyable
	Swimmable
}

// 4.5 Hybrid: shape 4.3 mixed with shape 4.4 (multilevel + multiple together)
type Base struct{}
type Mid struct{ Base } // multilevel part
type Other struct{}
type Combo struct { // multiple part added on top
	Mid
	Other
}

func demo4TypesOfInheritance() {
	fmt.Println("\n=== 4. Types of Inheritance (via struct embedding) ===")

	fmt.Print("4.1 Single: ")
	b := Bike{}
	b.Move() // reused from section 2.3

	fmt.Print("4.2 Multilevel: ")
	c := CChild{}
	c.Greet() // C has no Greet() of its own, it comes all the way from A

	fmt.Println("4.3 Hierarchical: Sedan and Truck both embed Vehicle2 (no output, just structure)")
	_ = Sedan{}
	_ = Truck{}

	fmt.Print("4.4 Multiple: ")
	fb := FlyingBoat{}
	fb.Fly()
	fb.Swim() // has behavior from BOTH embedded structs

	fmt.Println("4.5 Hybrid: Combo embeds through Mid and Other (no output, just structure)")
	_ = Combo{}
}

// =====================================================================================
// 5. THE DIAMOND PROBLEM
// Story: two managers, one shared CEO. Which manager's copy of the rule do you follow?
// Go has no multiple inheritance of classes, so this can only happen through embedding,
// and Go forces you to be explicit instead of guessing.
// =====================================================================================
type CEO struct{}

func (CEO) Rule() { fmt.Println("CEO.Rule") }

type Manager1 struct{ CEO } // its OWN copy of CEO
type Manager2 struct{ CEO } // a SEPARATE copy of CEO

type You struct {
	Manager1
	Manager2
}

func demo5DiamondProblem() {
	fmt.Println("\n=== 5. The Diamond Problem ===")
	y := You{}
	// y.Rule()             // Compile error: ambiguous selector, which path, Manager1 or Manager2?
	y.Manager1.CEO.Rule() // must spell out the exact path, Go refuses to guess
	fmt.Println("(This problem never happens in Go for classic multiple inheritance, only for ambiguous embedding like above.)")
}

// =====================================================================================
// 6. METHOD OVERLOADING (Compile-time Polymorphism)
// Story: a calculator's "add" button. Go does NOT allow two functions with the same
// name in one package, no matter how different their parameters are.
// =====================================================================================

// This would NOT compile in Go:
// func Add(a, b int) int { return a + b }
// func Add(a, b float64) float64 { return a + b }   // ERROR: Add already declared

func AddInts(a, b int) int           { return a + b } // Option 1: separate, clear names
func AddFloats(a, b float64) float64 { return a + b }

// Option 2: generics (Go 1.18+) plus variadic args, to cover different counts of the same type
func AddAll[T int | float64](nums ...T) T { // T can be int OR float64, Go figures it out from the input
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func demo6Overloading() {
	fmt.Println("\n=== 6. Method Overloading (Compile-time Polymorphism) ===")
	fmt.Println(AddInts(2, 3))       // 5
	fmt.Println(AddFloats(2.5, 3.5)) // 6
	fmt.Println(AddAll(1, 2, 3))     // works with ints
	fmt.Println(AddAll(1.1, 2.2))    // works with floats too
}

// =====================================================================================
// 7. METHOD OVERRIDING (Runtime Polymorphism)
// Reuses Dog / Cat from section 2.4 on purpose, overriding and polymorphism
// are really the same mechanism in Go, both come from interfaces.
// =====================================================================================
func demo7Overriding() {
	fmt.Println("\n=== 7. Method Overriding (Runtime Polymorphism) ===")
	animals := []Animal{Dog{}, Cat{}} // a list holding two different real types
	for _, a := range animals {
		a.Speak() // correct version picked at runtime, based on the real type
	}
}

// =====================================================================================
// 8. COMPILE-TIME vs RUNTIME POLYMORPHISM
// Go has no overloading, so it really only has the "runtime" column, always through
// interfaces. Reuses AddInts (section 6) and Animal (section 7) side by side.
// =====================================================================================
func demo8CompileVsRuntime() {
	fmt.Println("\n=== 8. Compile-time vs Runtime Polymorphism ===")
	fmt.Println("Compile-time-like (fixed function name, chosen by the programmer):", AddInts(2, 3))

	var a Animal = Dog{}
	fmt.Print("Runtime (decided while the program runs, from the real type): ")
	a.Speak()
}

// =====================================================================================
// 9. INTERFACES vs CONCRETE TYPES
// Story: a job post says "must be able to drive." That's the interface (the promise).
// The real person hired is the concrete type. No keyword needed to "match" it.
// =====================================================================================
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

func demo9InterfacesVsConcrete() {
	fmt.Println("\n=== 9. Interfaces vs Concrete Types ===")
	startTrip(Person{}) // Person just fits, because it has the right method
}

// =====================================================================================
// 10. COUPLING vs COHESION
// Story: a coworker who digs through another coworker's private notes (tight coupling)
// vs a coworker who only uses a shared, agreed form (loose coupling).
// =====================================================================================

// Tightly coupled version
type MySQLDatabase struct{}

func (MySQLDatabase) SaveOrder(data string) { fmt.Println("Saved to MySQL:", data) }

type OrderServiceTight struct {
	db MySQLDatabase // hardcoded, this struct can NEVER swap the database without editing itself
}

func (s OrderServiceTight) PlaceOrder(data string) { s.db.SaveOrder(data) }

// Loosely coupled version
type Database interface {
	Save(data string) // the shared, agreed form
}

type MySQLDatabase2 struct{}

func (MySQLDatabase2) Save(data string) { fmt.Println("Saved to MySQL:", data) }

type OrderServiceLoose struct {
	db Database // only knows the shared form, not the real database's private details
}

func NewOrderServiceLoose(db Database) OrderServiceLoose { return OrderServiceLoose{db: db} } // handed in from outside
func (s OrderServiceLoose) PlaceOrder(data string)       { s.db.Save(data) }

func demo10CouplingCohesion() {
	fmt.Println("\n=== 10. Coupling vs Cohesion ===")
	tight := OrderServiceTight{db: MySQLDatabase{}}
	tight.PlaceOrder("order-A (tightly coupled)")

	loose := NewOrderServiceLoose(MySQLDatabase2{}) // could hand in ANY Database implementation here instead
	loose.PlaceOrder("order-B (loosely coupled)")
}

// =====================================================================================
// 11. COMPOSITION OVER INHERITANCE
// Story: a car HAS an engine (a separate, swappable part), it does not "inherit" from one.
// =====================================================================================
type Engine struct{ Horsepower int }

func (e Engine) Start() { fmt.Println("Engine started with", e.Horsepower, "HP") }

type Car struct {
	engine Engine // named field, not embedded, makes "Car HAS-A Engine" explicit
}

func NewCar(hp int) Car { return Car{engine: Engine{Horsepower: hp}} } // Car builds its own Engine part
func (c Car) Drive() {
	c.engine.Start() // Car uses the Engine's behavior, but Car is not "an" Engine
	fmt.Println("Car is driving...")
}

func demo11Composition() {
	fmt.Println("\n=== 11. Composition over Inheritance ===")
	car := NewCar(120)
	car.Drive()
}

// =====================================================================================
// 12. ASSOCIATION vs AGGREGATION vs COMPOSITION
// Three "has-a" strengths: developer-repo (weakest), team-microservice (medium),
// order-lineitem (strongest, parts die with the whole).
// =====================================================================================

// Association: both sides exist independently, just point at each other
type Repository struct{ Name string }
type Developer struct {
	Username string
	Repos    []*Repository // just references, Developer does not own these
}

// Aggregation: Team "has" Microservices, but they are made outside and can outlive the team
type Microservice struct{ Name string }
type Team struct {
	Services []*Microservice // handed in from outside, Team did not create these
}

func (t *Team) AddService(m *Microservice) { t.Services = append(t.Services, m) }

// Composition: Order OWNS its LineItems, they are created and destroyed together with the Order
type LineItem struct {
	Product string
	Qty     int
}
type Order struct {
	items []LineItem // Order builds its own line items inside itself
}

func (o *Order) AddItem(product string, qty int) {
	o.items = append(o.items, LineItem{Product: product, Qty: qty}) // built inside Order, owned by it
}
func (o *Order) ListItems() {
	for _, item := range o.items {
		fmt.Println(" -", item.Product, "x", item.Qty)
	}
}

func demo12AssociationAggregationComposition() {
	fmt.Println("\n=== 12. Association vs Aggregation vs Composition ===")

	repo := Repository{Name: "OOP-context"}
	dev := Developer{Username: "niloy104", Repos: []*Repository{&repo}}
	fmt.Println("Association:", dev.Username, "knows about repo", dev.Repos[0].Name)

	svc := Microservice{Name: "auth-service"}
	team := Team{}
	team.AddService(&svc)
	fmt.Println("Aggregation: team has service", team.Services[0].Name,
		"(service would survive even if the team is disbanded)")

	order := Order{}
	order.AddItem("Keyboard", 1)
	order.AddItem("Mouse", 2)
	fmt.Println("Composition: order owns these line items directly:")
	order.ListItems()
}

// =====================================================================================
// MAIN - runs every topic's demo, in the same order as the note
// =====================================================================================
func main() {
	demo1ObjectAndClass()
	demo2_1Encapsulation()
	demo2_2Abstraction()
	demo2_3Inheritance()
	demo2_4Polymorphism()
	demo3AccessModifiers()
	demo4TypesOfInheritance()
	demo5DiamondProblem()
	demo6Overloading()
	demo7Overriding()
	demo8CompileVsRuntime()
	demo9InterfacesVsConcrete()
	demo10CouplingCohesion()
	demo11Composition()
	demo12AssociationAggregationComposition()
}