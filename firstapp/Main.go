package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// outside function must use var
var z int = 43

// if you need conversion fron int to string use strconv package
// var blok
var (
	nama  string  = "candra"
	age   int     = 30
	value float32 = 100000000
)

func printslice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func printmap(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func reverse(n int) int {
	rev := 0
	for n != 0 {
		rev = rev*10 + n%10
		fmt.Println()
		n /= 10
		fmt.Printf("rev %v n %v", rev, n)

		// rev 1 n = 432
		// rev 10 + 2 = 12 n 43
		// rev 120 + 3 = 123 n 4
		// rev 1230 + 4 = 1234 n 0
	}
	return rev
}

type car struct {
	brand string
	model string
	Year  int
}

func printcar(c car) {
	fmt.Println("brancd :", c.brand)
	fmt.Println("model :", c.model)
	fmt.Println("year :", c.Year)
}

func getCar(c *car) car {
	return car{
		brand: c.brand,
		model: c.model,
		Year:  c.Year,
	}
}

func (c *car) updateCar() car {
	c.brand = "honda"
	c.model = "hatchback"
	c.Year = 200
	return *c
}

func checkCansplitArray(inputArr []int) int {
	total := 0

	for _, value := range inputArr {
		total += value
	}

	leftSum := 0
	for i := 0; i < len(inputArr); i++ {
		leftSum += inputArr[i]
		// always new right sum
		rightSum := total - leftSum
		if leftSum == rightSum {
			return 1
		}
	}
	return 0
}

func isPaliandrome(numberInput int) int {
	// sorting from minimum to maximum
	//sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	original := numberInput
	reversed := 0

	if numberInput < 0 {
		return 0
	}
	for numberInput != 0 {
		remainder := numberInput % 10
		reversed = reversed*10 + remainder
		numberInput /= 10
	}
	if original == reversed {
		return 1
	} else {
		return 0
	}
}

type shape interface {
	area() float32
	perimeter() float32
}

type Rectangle struct {
	width, length float32
}

type Circle struct {
	radius float32
}

func (r Rectangle) area() float32 {
	return r.length * r.width
}

func (c Circle) area() float32 {
	return math.Phi * c.radius * c.radius
}

// wiring up declared
func calculatesArea(s shape) float32 {
	return s.area()
}

// interface must be implemented
func (c Circle) perimeter() float32 {
	return 2 * math.Phi * c.radius
}

func (r Rectangle) perimeter() float32 {
	return r.length*2 + r.width*2
}

type geometri interface {
}

func main() {
	car := car{
		brand: "toyota",
		model: "Kijang",
		Year:  14,
	}
	fmt.Println(getCar(car))
	printcar(car)
	car.updateCar()
	printcar(car)
	// intialization
	rect := Rectangle{width: 30, length: 30}
	circle := Circle{radius: 20}
	fmt.Println("calculate area rectangle", calculatesArea(rect))
	fmt.Println("calculate area circle", calculatesArea(circle))

	//struct_example()
	//constants()
	fmt.Println("reverse of 4321 is", reverse(4321))
	fmt.Println(checkCansplitArray([]int{1, 3, 3, 4, 3}))
	fmt.Println(isPaliandrome(123))
	array()
	//slice()
	//make_slice()
	//matric()
	//use_of_iota()
	// primitive()
	// variable()
	//map_example()
	//call_api()
	//rest_api()
	//control_flow()
	//get_method()
	//call_api()
	//pointer()
	// test_pointer_again()
	// for i := 0; i < 5; i++ {
	// 	sayMessage("haloo Goooo", i)
	// }

	fmt.Println("sum is ", test_pass_param())
	// anonymous function
	for i := 0; i < 5; i++ {
		func() {
			msg := "Hello Go"
			fmt.Println(msg, i)
		}() // ()--> it mean executed directly
	}

	// another example function asign to variable
	f := func() {
		fmt.Println("ini didalam variable f")
	}

	f() // --> how to call

	//example_anym()
	//example_method()
	//rest_api()

	var w Writer = ConsoleWriter{}
	w.Write([]byte("hello GO !"))

	MyInt := IntCounter(0)
	var inc Incrementer = &MyInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	add_date()

	//anym_struct()
	//example_embeded_struct()
	//other_struc()
	//simple_if_stmnt()
	//case_statement()
	//anotther_case_switch()
	//another_switch()
	//adv_switch()
	//simple_loop()
	//number_guess()
	//example_else()
}

func add_date() {
	// Get the current date and time
	currentTime := time.Now()

	// Add one month to the current date
	oneMonthLater := currentTime.AddDate(0, 0, 90)

	// Format and print the results
	fmt.Println("Current Date:", currentTime.Format("2006-01-02"))
	fmt.Println("One Month Later:", oneMonthLater.Format("2006-01-02"))
}

func Sqrt(x float64) float64 {
	z := float64(1)
	epsilon := 0.00000000001
	for diff := z*z - x; diff > epsilon || diff < -epsilon; {
		z -= (z*z - x) / (2 * z)
		diff = z*z - x
		fmt.Println(x, z)
	}
	return z
}

//func main() {
//	fmt.Println("hasil akhir ",Sqrt(8))
//}

// ///// pointer /////////////
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	fmt.Println("nilainy v.X yang kedua ", v.X)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	fmt.Println("di inisialisasi ulang")
	fmt.Println("v awal :", v)
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("nilainy v.X ", v.X)
}

// / kalau method diatas ditulis dalam bentukl function inputnya harus ada &
func Scale_ref(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("nilainy v.X ", v.X)
}

///// jika * dihapus nilai v.X dan v.Y tidak berubahh disemua tempat //////

func test_pointer_again() {
	v := Vertex{3, 4}
	fmt.Println("nilai v:", v)
	fmt.Println("V ABS BEFORE SCALE : ", v.Abs())
	v.Scale(10) /// CARA memanggil nya lebih natural
	fmt.Println("V ABS AFTER SCALE : ", v.Abs())
	//// cara manggil nya...yang method
	Scale_ref(&v, 10)
}

///////////////////

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Writer interface {
	Write([]byte) (int, error)
}

// Type definition
type data int

// Defining a method with
// non-struct type receiver
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

type greeter struct {
	greeting string
	name     string
}

// example of method
func example_method() {
	g := greeter{
		greeting: "hello",
		name:     "go",
	}
	g.greet()
	fmt.Println("The new name ", g.name)

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	value1 := data(1)
	value2 := data(2)
	result := value1.multiply(value2)
	fmt.Println("hasil perkalian ", result)

}

// func (g greeter) greet() { // --> METHOD DECLARATION //(g greeter) context that , that function is executing in
// 	fmt.Println(g.greeting, g.name)
// 	g.name = "" // this is will not modify g
// }

func (g *greeter) greet() { // --> METHOD DECLARATION //(g greeter) context that , that function is executing in
	fmt.Println(g.greeting, g.name)
	g.name = "candra" // this is will modify value G.. because the context is pointer
}

func example_anym() {
	var devide func(float64, float64) (float64, error) // define variable
	devide = func(f1, f2 float64) (float64, error) {
		if f2 == 0.0 {
			return 0.0, fmt.Errorf("Cannot devide by Zero")
		} else {
			return f1 / f2, nil
		}
	}
	d, err := devide(5.0, 3.0) // --> call the function
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hasil")
	fmt.Println(d)
}

func test_pass_param() int {
	greeting := "hello"
	name := "Stacy"
	sayGreeting(greeting, name)
	fmt.Println(name)
	s := sum(1, 2, 3, 4, 5) // there are 5 parameter that passed in
	fmt.Println(s)

	c, err := devide(5, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("devide ", c)

	return s

}

func devide(a, b float64) (float64, error) {
	if b == 0.0 {
		// panic("cannot devide by zero") // if we use panic application will close and cannot continue
		return 0.0, fmt.Errorf("cannot devide by zero")
	}
	return a / b, nil
}

func sum(values ...int) int { // veriadic parameter, take all last argument that are passed in, and wrap them up into a slice that has the name of "values"
	fmt.Println(values) // become slice [1,2,3,4,5]
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("the sum is ", result)
	return result
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("the value of the index is", idx)
}

func pointer() {
	var a int = 10
	//var b *int = &a //b not get coppied data from a, but b holding the memory location of a
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(&a, b) //&a print memory address of a
	// to dereference where pointer pointing at use this
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b) // value b now became 27 also, beacuse b holding memory location of a
	*b = 14
	fmt.Println(a, *b) // both value are changing
}

func rest_api() {
	http.HandleFunc("/halo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go"))
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err.Error())
	}
}

func panicker() {
	fmt.Println("about to panic")
	x, b := 1, 1

	ans := x / b
	fmt.Println(ans)

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error :", err)
		}
	}()
	panic("something bad happen")
	fmt.Println("done panicking")
}

func call_api() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()                  // to avoid forget close the resource, we can call first, it will executed later
	robots, err := ioutil.ReadAll(res.Body) // read response and print
	//res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("halooo %s", robots)
}

func control_flow() {
	fmt.Println("start")
	defer fmt.Println("middle") //defer will delay until last statement, and excute it
	fmt.Println("near end")
	fmt.Println("end")

	// defer hold value where it first time called
	a := "start"
	defer fmt.Println(a)
	a = "end"
	// it will print "start" eventhough value "a" has been overide

	//example panic

	panicker()

	fmt.Println("recover")
}

func simple_loop() {
	fmt.Println("======loop through slice=====")
	s := []int{1, 2, 3}
	for key, value := range s {
		fmt.Println(key, value)
	}

	str := "hallo candra"
	for key, value := range str {
		fmt.Println(key, string(value))
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
		// var i has been modified during looping // bad coding because modify the counter
	}
	fmt.Println("=========")
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("=========")

	// or increment process inside loop
	y := 0
	for y < 5 {
		fmt.Println(y)
		y++
	}
	fmt.Println("=========")
	for i := 0; i < 5; i = i + 2 {
		fmt.Println("another case", i)
	}
	// another more complex loop
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	fmt.Println("=========")
	// implement break
	x := 0
	for {
		fmt.Println("break", x)
		x++
		if x == 5 {
			break
		}
	}
	fmt.Println("=========")
	// exmple only printt odd number using continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("======nested loop=====")
	for i := 1; i < 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
		}
	}

	fmt.Println("======break outer loop=====")
loop:
	for i := 1; i < 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break loop
			}
		}
	}

}

func simple_if_stmnt() {
	statePopulation := map[string]int{ //key string value int --> must be consistent
		"california": 10000,
		"texas":      10000,
		"New York":   10000,
		"Ohio":       10000,
		"Florida":    10000,
	}
	if value, ok := statePopulation["Florida"]; ok { //value, ok := statePopulation["Florida"]; --> inializer //;ok --> seek value it should Boolean
		fmt.Println(value)
	}

	fmt.Println(statePopulation)

	if true {
		fmt.Println("the test is True")
	}
}

func number_guess() {
	number := 50
	guess := -5                   //guess := 100 -> this is will go to first if
	if guess < 1 || guess > 100 { // or
		fmt.Println("the Guess must be between 1 and 100!")
	}
	if guess >= 1 && guess <= 100 { // and
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("equal, you got it")
		}
	}

	fmt.Println(!(number <= guess), "ada sesuatu disini", number >= guess, number != guess)
}

func example_else() {
	number := 50
	guess := 102   //guess := 100 -> this is will go to first if
	if guess < 1 { // or
		fmt.Println("the Guess must be greater than 1")
	} else if guess > 100 {
		fmt.Println("the Guess must be less than 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("equal, you got it")
		}
	}

}

func case_statement() {

	switch 5 {
	case 1, 5, 10: // test case must  be unique
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}

func anotther_case_switch() {

	switch i := 2 + 7; i {
	case 1, 5, 10: // test case must  be unique
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}

func another_switch() {
	i := 10
	switch {
	case i <= 10: // go will evaluate first case first,. if okay go will not run  through all case left
		fmt.Println("first one")
		fallthrough // if we want go evaluate next case
	case i <= 20:
		fmt.Println("second two")
	default:
		fmt.Println("not one or two")
	}
}

func adv_switch() {
	var i interface{} = [3]int{} // interface can assign to all data type
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}

}

type Animal struct {
	Name   string `required max:"100"` //validation
	Origin string
}

type Bird struct {
	Animal   //embeded struct
	SpeedKPH float32
	CanFly   bool
}

func other_struc() {
	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 30,
		CanFly:   false,
	}
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	fmt.Println(b.Origin)

}

type Doctor struct {
	number      int      //inisializer
	actorName   string   //inisializer
	compnanions []string //inisializer
}

func example_embeded_struct() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Animal.Name) // or
	fmt.Println(b.Name)
}

func struct_example() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Candra Ganteng",
		compnanions: []string{
			"farida", "zafran", "sarwiyah",
		},
	}

	fmt.Println("example looping inside dictionary")

	v := reflect.ValueOf(aDoctor)
	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).CanInterface()
		fmt.Println(values)
	}
	not_exist_key := []string{}
	// search key exist
	metaValue := reflect.ValueOf(aDoctor)
	for _, key := range []string{"address"} {
		field := metaValue.FieldByName(key)
		if field == (reflect.Value{}) {
			//log.Printf("Field %s not exist in struct", key)
			not_exist_key = append(not_exist_key, key)
		}
	}
	fmt.Println("list of not exist key")
	fmt.Println(not_exist_key)
	// values := reflect.ValueOf(aDoctor)
	// typesOf := values.Type()

	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Printf("Field: %s\tValue: %v\n", typesOf.Field(i).Name, values.Field(i).Interface())
	// }

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.compnanions[1])
	fmt.Println(aDoctor.actorName)
}

func anym_struct() {
	aDoctor := struct{ name string }{name: "John"}
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}

func map_example() {
	statePopulation := map[string]int{ //key string value int --> must be consistent
		"california": 10000,
		"texas":      10000,
		"New York":   10000,
		"Ohio":       10000,
	}

	print("print key value over map")
	for k, v := range statePopulation {
		fmt.Println(k, v)
	}

	print("ignored variable _")
	for _, v := range statePopulation {
		fmt.Println(v)
	}

	// map of array of integer
	m := map[[3]int]string{} // array is valid map key but slice is not

	newmap := make(map[int]string)
	newmap[1] = "one"
	newmap[2] = "two"

	fmt.Println(newmap[1])

	fmt.Println(statePopulation, "map baru :", m)
	fmt.Println(statePopulation["california"])

	//adding more item
	statePopulation["ngawi"] = 10000000
	fmt.Println(statePopulation)

	//delete key in the map
	delete(statePopulation, "ngawi") //(name of map, key that want to delete)
	fmt.Println("map after deletion :", statePopulation)

	// check key exist or not
	pop, ok := statePopulation["Ohio"]
	fmt.Println(pop, ok) //==> 10000 true ==> key exist

	pop, ok = statePopulation["Ohiooo"]
	fmt.Println(pop, ok) //==> 0 false ==> key does not exist

	//check how many element
	fmt.Println("length of map :", len(statePopulation))

}

// example define role
const (
	isadmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAsia
	canSeeAffrica
)

const (
	errorr = iota // 0
	cat           // 1
	dog           // 2
	snake         // 3
)

func array() {
	grades := [3]int{97, 85, 93} // 3 length of array, int type of array, { value inside array}
	fmt.Printf("Grades %v \n", grades)

	var student [5]string
	student[0] = "Lisa"
	student[1] = "Ahmed"
	student[2] = "Arnold"
	fmt.Printf("Student %v \n", student)
	fmt.Printf("Student #1 %v \n", student[1])
	fmt.Printf("number of student Student %v \n", len(student))

	a := [...]int{1, 2, 3, 4} //exact array
	c := []int{1, 2, 3, 4}    //underliying array = slices array
	d := c
	d[2] = 6 // original slice array will change accordingly
	b := a   // make copy of array b
	b[1] = 5 // reassign array, this will not changes original array (a )
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("underlying array / slice array")
	fmt.Println(c)
	fmt.Println(d)
}

func slice() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice from 4th element to end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice the 4th, 5th, 6th elements
	// slicing ( : ) not changes original value
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(a)
	// this change original value
	f := a
	f[1] = 10
	fmt.Println(a)
	fmt.Println(a[4:5])
}

func matric() {
	//var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println("below, example of matrix")
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)
}

func use_of_iota() {
	var roles byte = isadmin | canSeeAffrica | canSeeAsia // example define role
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin ? %v \n", isadmin&roles == isadmin)            // is admin exist in role => True
	fmt.Printf("Is HQ ? %v \n", isHeadquarters&roles == isHeadquarters) // is HQ exist in role => false
}

func constants() {
	// default value is 0
	var animalType int                       // 0
	fmt.Printf("%v\n", animalType == cat)    // return False
	fmt.Printf("%v\n", animalType == errorr) // True

	// just like variable constant can be shadowed ( overirde by naother value , but canot in same inside function)
	const x = iota // constante no need to be used
	// iota is counter on constant that has been declared
	const (
		a = iota
		b = iota
		c = iota
	)

	// if same blok we also can use this
	// const (
	// 	a = iota
	// 	b
	// 	c
	// )

	// reset iota becaue different blok
	const (
		d = iota
	)

	const myConst int = 42
	fmt.Printf("%v, %T \n", myConst, myConst)

	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T \n", d, d)

}

func make_slice() {
	//a := make([]int, 3, 100) // 3 is length of the slice // 100 is capacity // another example create slice
	var a []int
	fmt.Println(a)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("capacity: %v \n", cap(a))

	a = append(a, 1) // add 1 to existing array [0,0,0] // first argument is name of slice variable that want to append
	a = append(a, 2)
	fmt.Println(a)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("capacity: %v \n", cap(a))

	// remove midle alement
	x := []int{1, 2, 3, 4, 5}
	y := append(x[:2], x[3:]...) // be carefull this will change underliying slice x!!!
	fmt.Println(x)
	fmt.Println(y)
}

func primitive() {
	var n bool = true
	b := 1 == 1 // comparation
	m := 1 == 2 // comparation
	var a uint16 = 42

	x := 3
	y := 2

	c := 8 // 2 ^3

	z := 3.4
	z = 13.41233
	fmt.Printf("%v, %T \n", n, n)
	fmt.Printf("%v, %T \n", m, m)
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", a, a)
	fmt.Println(x * y) //devide
	fmt.Println(x / y) //devide
	fmt.Println(x % y) // remind / mod
	// note: every new defined variable, if not initialize will have falue 0 and False

	fmt.Println(c << 3) // 2^3 * 2^3 = 2 ^6 = 64
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0 = 1
	fmt.Println(z)

}

func variable() {
	var i int //initiation later
	i = 42
	var b int = 24
	//or
	//$v = value of variable
	//$T type of the varaible
	c := 42 //faster declare new variable
	x := "candra"
	y := "pinter"

	//convert string to byte
	w := []byte(x)

	//conversion string to int
	c = int(value)
	var z string
	z = strconv.Itoa(i)
	fmt.Print("candra world, baru i \n")
	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(x)
	fmt.Println(w)
	fmt.Println(age)
	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T \n", z, z)
	fmt.Printf(x + y) // string concatenation

	//get value of char in dex 4 and convert to string
	fmt.Printf("%v, %T \n", string(x[4]), x)
}

func get_method() {
	// Empty String Payload for GET METHOD
	strBody := ""

	// Parse To Json
	reqBodyData, _ := json.Marshal(strBody)
	stringData := string(reqBodyData)
	fmt.Println("\nstringData = ", stringData)

	//To make the data same as json.dump in python | Golang 1.11+
	//stringData = strings.ReplaceAll(stringData, "\":", "\": ")
	//stringData = strings.ReplaceAll(stringData, ",\"", ", \"")

	// to cater json.dump in Python for Go 1.11
	// Create replacer with pairs as arguments.
	r := strings.NewReplacer("\":", "\": ",
		",\"", ", \"")
	stringData = r.Replace(stringData)

	url := "https://griffin.bigdata.ovo.id/apis/prakerja/validate?phone_number=+6281330141337&hash_nik=b3058015703ff9bd4f65c0217225da33aae97c9c"

	// Generate HMAC
	username := "prakerja.developer"
	secretKey := "AmMXc23UZRrLwQ2VADzZIGJuXyKqxWxnEpJjcpImxc543"
	hmacAuth := generateHmac(username, secretKey, url, stringData)

	fmt.Println("hmacAuth = ", hmacAuth)

	// Construc Request & Header
	payload := bytes.NewBufferString(stringData)
	req, err := http.NewRequest(http.MethodGet, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", hmacAuth)

	// Call Respective API
	client := &http.Client{Timeout: time.Second * 10}
	res, err := client.Do(req)
	defer res.Body.Close()

	// Handle Response
	response, err := ioutil.ReadAll(res.Body)

	fmt.Println("\nResponse : ", string(response))
}

func generateHmac(user string, sKey string, url string, body string) string {
	//username := "ovopepipost"
	//secretKey := "bEvOL7Rh5bXKHInkOxiR2p7suhr1TsPuOEDoMRVgIC4YtSMGwrGQoJZuau5inj46"
	username := user
	secretKey := sKey

	timestamp := strconv.Itoa(int(time.Now().UTC().Unix()))

	data := fmt.Sprintf("%s%s%s%s%s", username, timestamp, secretKey, url, body)
	//data := username + timestamp + secretKey + url + body
	fmt.Println("hmac Payload = ", data)

	byteSecret := []byte(secretKey)

	hmacGenerated := hmac.New(sha256.New, byteSecret)
	// Write Data to it
	hmacGenerated.Write([]byte(data))

	signature := base64.StdEncoding.EncodeToString(hmacGenerated.Sum(nil))

	auth := fmt.Sprintf("hmac %s:%s:%s", username, timestamp, signature)

	return auth
}
