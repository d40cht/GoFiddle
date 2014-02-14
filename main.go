package main

import "fmt"

func main() {
  var message string = "Hello world from Go."
  message2 := "Now with type inference."
	const message3 = "And constants."
	fmt.Println(message, message2, message3)

  var zeros [3] int
  x := [3] int { 6, 7, 8 }

  // TODO: Slices, although the syntax is horrid
  slice1 := [] int { 1, 3, 5 }
	slice2 := make([]float64, 3)

  for i := 0; i < 3; i++ {
		fmt.Println("Counter is:", i, x[i], zeros[i], slice1[i], slice2[i])
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
}
