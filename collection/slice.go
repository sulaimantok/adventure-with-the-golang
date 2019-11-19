package main

import "fmt"

func main() {
	mySlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(mySlice)
	// [2 3 5 7 11 13]

	fmt.Println(mySlice[1:4])
	// [3 5 7]

	// missing low index implies 0
	fmt.Println(mySlice[:3])
	// [2 3 5]

	// missing high index implies len(s)
	fmt.Println(mySlice[4:])
	// [11 13]


	names := [4]string{
		"Amin",
		"johan",
		"Amar",
		"faiz",
	}
	fmt.Println(names)

	a := names[0:2]    //slice a 
	b := names[1:3]      //slice b
	fmt.Println(a, b)

	b[0] = "XXX"      // value at zeroth index of slice b changed
	fmt.Println(a, b)
	fmt.Println(names)


	cities := make([]string, 3)
	cities[0] = "Jakarta"
	cities[1] = "Surabaya"
	cities[2] = "Semarang"
	fmt.Printf("%q", cities)

	cities := []string{"Jakarta", "Surabaya"}
	otherCities := []string{"Jakarta", "Semarang"}
	cities = append(cities, otherCities...)
	fmt.Printf("%q", cities)
}