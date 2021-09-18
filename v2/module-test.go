// To test go module in other repository.
//
//Yes we go on Beyond with v2: v2 changelog
//
//Adding GoPrint to old function
package modulecheck

import "fmt"

// ModPrint is used to print "You're succeed", and now Updated
// v2 Include in-house GoPrint
func ModPrint() {
	fmt.Println("You're succeed")
	fmt.Println("Now Update")
	GoPrint()
}

// GoPrint is used to print We're gophers
// 
//		with the addition about to publish
func GoPrint()  {
	fmt.Println("We're Gophers and about to publish")
}
