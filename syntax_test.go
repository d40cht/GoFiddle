package main

import "fmt"
import "testing"

func makeHalfAgain(in float64) float64 {
	return in * 1.5
}

func minMax(first float64, second float64) (float64, float64) {
	if first < second {
		return first, second
	}
	return second, first
}

type TestStruct struct {
	x int
	y float32
	z float64
}

type Shape interface {
	area() float64
}

type Square struct {
	width float64
}

type Rectangle struct {
	width float64
	height float64
}

// Makes square implement the shape interface
func (sq *Square) area() float64 {
	return sq.width * sq.width
}

// Makes rectangle implement the shape interface
func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}

func printArea(shp Shape) {
  fmt.Println("Shape area:", shp.area())
}

func shapeInterfaceTest() {
  var s Square = Square {4.0}
  r := Rectangle {3.0, 2.0}

	// Why, oh why, am I passing a pointer here when the declaration
	// of printArea looks like it takes a Shape by value?
	printArea(&s)
	printArea(&r)
}

// A method on TestStruct
func (ts *TestStruct)printTestStruct() {
	fmt.Println("TestStruct:", ts.x, ts.y, ts.z)
}

func TestSyntax(t *testing.T) {
  // Argument to defer must be a function call. Doesn't need to be a closure,
	// although here it is.
  defer func() {
		fmt.Println("This line should appear last of all.")
	}()

  var message string = "Hello world from Go."
  message2 := "Now with type inference."
	const message3 = "And constants."
	fmt.Println(message, message2, message3)

  var zeros [3] int
  x := [3] int { 6, 7, 8 }

  // TODO: Slices, although the syntax is horrid
  slice1 := [] int { 1, 3, 5 }
	slice2 := make([]float64, 3)
	slice3 := slice1[0:3]

  for i := 0; i < 3; i++ {
		fmt.Println("Counter is:", i, x[i], zeros[i], slice1[i], slice2[i], slice3[i])
	}

  for _, val := range slice3 {
    fmt.Println("Range loop:", val, makeHalfAgain(float64(val)))
	}

  testMap := make(map[string]int)
  testMap["foo"] = 12
  testMap["bar"] = 16
	if len(testMap) != 2 {
		t.Error("Test map has unexpected size:", len(testMap))
  }
	delete(testMap, "foo")
	if len(testMap) != 1 {
		t.Error("Test map has unexpected size:", len(testMap))
	}
	
	count, ok := testMap["bar"]
  if !ok {
		t.Error("Could not find 'bar' in testMap")
	}
	if count != 16 {
		t.Error("Key 'bar' had unexpected value association")
	}

  prefilledMap := map[string]string { "foo" : "bar", "baz" : "qux" }
  if len(prefilledMap) != 2 {
    t.Error("Unexpected prefilled map length")
	}

  min1, max1 := minMax(10.0, 3.0) 
	min2, max2 := minMax(3.0, 10.0)
  if !(min1 == 3.0 && min2 == 3.0 && max1 == 10.0 && max2 == 10.0) {
		t.Error("minMax function returned incorrect values")
	}

  // Bound by reference in the closure
  var xbound = 3
  closureTest := func(v int) int {
		return v + xbound
  }

	if closureTest(2) != 5 {
		t.Error("Error in closure test")
	}
  xbound = 5
  if closureTest(2) != 7 {
		t.Error("Unexpected behaviour of closure-captured variables")
	}

  // TODO: Pointers, dereference, new etc.



	ts := TestStruct { 1, 3.0, 7.0 }	
  fmt.Println( ts.x, ts.y, ts.z )
  ts.printTestStruct()


	// Next: concurrency. Goroutines, channels etc.
}
