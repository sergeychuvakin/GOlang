package tutorial

import (
	"fmt"
	"log"
	"strings"
)

var name string
var email string
var city string

// var WarningLogger log.Logger

func main() {

	// WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("City: ")
	fmt.Scan(&city)

	isValidLen := len(name) > 2 && len(email) > 2
	isContainsDog := strings.Contains(email, "@")
	isValildCity := city == "Moscow" || city == "London"

	if isValidLen && isContainsDog && isValildCity {

		log.Println(isValidLen)
		log.Println(isContainsDog)
		log.Println(isValildCity)

		log.Printf("\nName is: %v\nEmail is: %v\n", name, email)

	} else {
		// WarningLogger.Println("s")
		log.Print("Ooops, something went wrong.")
	}

}
