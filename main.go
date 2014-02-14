package main

import "fmt"

func main() {
  var message string = "Hello world from Go."
  message2 := "Now with type inference."
	const message3 = "And constants."
	fmt.Println(message, message2, message3)

  var zeros [3] int
  x := [3] int { 6, 7, 8 }

  for i := 0; i < 3; i++ {
		fmt.Println("Counter is: ", i, x[i], zeros[i])
	}
}
