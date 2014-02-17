package main

//import "fmt"
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

func checkArea(t *testing.T, shp Shape, area float64) {
  if shp.area() != area {
		t.Error("Shape has invalid area:", shp.area(), area)
	}
}

func TestShapeInterface(t *testing.T) {
  var s Square = Square {4.0}
  r := Rectangle {3.0, 2.0}

	// Why, oh why, am I passing a pointer here when the declaration
	// of printArea looks like it takes a Shape by value?
  checkArea(t, &s, 16.0)
	checkArea(t, &r, 6.0)
}

// A method on TestStruct, showing the need for explicit casts for
// mixed-type arithmetic operations in Go.
func (ts *TestStruct) structArith() float64 {
	return float64(ts.x) * float64(ts.y) + ts.z
}

func TestSyntax(t *testing.T) {
  // Argument to defer must be a function call. Doesn't need to be a closure,
	// although here it is.
  /*defer func() {
		fmt.Println("This line should appear last of all.")
	}()*/

  var value1 int = 1
	value2 := 2
	const value3 = 3
	if (value1 + value2 + value3) != 6 {
		t.Error("Simple variable declaration test fail.")
	}

  var zeros [3] int
  x := [3] int { 6, 4, 2 }

  // TODO: Slices, although the syntax is horrid
  slice1 := [] int { 1, 3, 5 }
	slice2 := make([]float64, 3)
	slice3 := slice1[0:3]

  for i := 0; i < 3; i++ {
		if !(zeros[i] == 0.0 && slice2[i] == 0.0 && (x[i] + slice1[i]) == 7 && slice1[i] == slice3[i]) {
			t.Error("Slices/arrays have unexpected values.")
		}
	}

	accumulator := 0.0
  for _, val := range slice3 {
		accumulator += 2.0*makeHalfAgain(float64(val))
	}
	if accumulator != 27.0 {
		t.Error("Range loop did not calculate expected value.")
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
  if !(ts.x == 1 && ts.y == 3.0 && ts.z == 7.0) {
		t.Error("Incorrect struct assignment.")
	}
  if ts.structArith() != 10.0 {
		t.Error("Incorrect struct arithmetic function")
	}


	// Next: concurrency. Goroutines, channels etc.
}
