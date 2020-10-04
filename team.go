package main

import "fmt"

func main() {
	var i int
	_, err := fmt.Scanf("%d\n", &i)
	if err != nil {
		panic(err)
	}

	r := 0
	for cpt := 0; cpt < i; cpt++ {
		var a, b, c int
		_, err = fmt.Scanf("%d %d %d\n", &a, &b, &c)

		if err != nil {
			panic(err)
		}

		if a+b+c >= 2 {
			r++
		}

	}
	fmt.Println(r)

}
