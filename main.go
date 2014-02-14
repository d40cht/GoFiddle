package main

import "fmt"


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

// A method on TestStruct
func (ts *TestStruct)printTestStruct() {
	fmt.Println("TestStruct:", ts.x, ts.y, ts.z)
}

func main() {
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
  delete(testMap, "foo")
  fmt.Println("Test map size:", len(testMap))

  if count, ok := testMap["bar"]; ok {
		fmt.Println("Found bar in your map sir: %d.", count)
	}

  prefilledMap := map[string]string { "foo" : "bar", "baz" : "qux" }
  fmt.Println("Length of prefilled map:", len(prefilledMap))

  min1, max1 := minMax(10.0, 3.0) 
	min2, max2 := minMax(3.0, 10.0)
	fmt.Println(min1, max1)
	fmt.Println(min2, max2)

  // Bound by reference in the closure
  var xbound = 3
  closureTest := func(v int) int {
		return v + xbound
  }

  fmt.Println("Fn 3, 2", closureTest(2))
  xbound = 5
  fmt.Println("Fn 5, 2", closureTest(2))


  // TODO: Pointers, dereference, new etc.



	ts := TestStruct { 1, 3.0, 7.0 }	
  fmt.Println( ts.x, ts.y, ts.z )
  ts.printTestStruct()
}
