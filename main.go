package main

import (
	"fmt"
	str "strings" // Package Alias

	"github.com/fredrikstrandin/golib/numbers"
	"github.com/fredrikstrandin/golib/strings"
	"github.com/fredrikstrandin/golib/strings/greeting" // Importing a nested package
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("fredrikstrandin"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
