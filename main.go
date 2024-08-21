package main

import "fmt"

func main() {

	
	var planet [8] string  
	planet[0]= "Mercury"
	planet[1]= "Venus"
	planet[2]= "Earth"
	planet[3]= "Mars"
	planet[4]= "Jupitar"
	planet[5]= "Saturn"
	planet[6]= "Uranus"
	planet[7]= "Neptune"

	fmt.Printf("We have 8 planets %v\n", planet)
	fmt.Printf("The smallest planet is %s\n", planet[0])
	fmt.Printf("The brightest planet is %s\n", planet[1])
	fmt.Printf("The Biggest planet is %s\n", planet[4])
	fmt.Printf("The red planet is %s\n", planet[3])
	fmt.Printf("The planet with life is %s\n", planet[2])

	fmt.Printf("We have %d planets in the universe", len(planet))
}