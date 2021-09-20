// To test go module in other repository.
//
// Yes we go on Beyond with v2: v2 changelog
//
// Nothing much, just test
package modulecheck

import "fmt"

// ModPrint is used to print "You're succeed", and now Updated
func ModPrint() {
	fmt.Println("You're succeed")
	fmt.Println("Now Update")
}

// GoPrint is used to print We're gophers
// 
//		with the addition about to publish
func GoPrint()  {
	fmt.Println("We're Gophers and about to publish")
}
