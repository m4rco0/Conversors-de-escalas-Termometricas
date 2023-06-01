// You can edit this code!
// Click here and start typing.
package main

import "fmt"

const ebuK float64 = 373.2

// C = K -237
func main() {
	tempK := ebuK
	tempC := int(tempK - 273)
	fmt.Printf("temperatura em K = %g temperatura em C = %v", tempK, tempC)

}

