package main

import (
	"fmt"
)

func main() {
	var nome string
	var idade int

	fmt.Scan(&nome)
	fmt.Scan(&idade)

	if idade < 12 {
		fmt.Printf("%v eh crianca\n", nome)
	} else if idade < 18 {
		fmt.Printf("%v eh jovem\n", nome)
	} else if idade < 65 {
		fmt.Printf("%v eh adulto\n", nome)
	} else if idade < 1000 {
		fmt.Printf("%v eh idoso\n", nome)
	} else {
		fmt.Printf("%v eh mumia\n", nome)
	}
}
