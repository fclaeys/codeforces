package main

import "fmt"

func main() {
	var i int
	_, err := fmt.Scanf("%d\n", &i)
	if err != nil {
		panic(err)
	}

	for cpt := 0; cpt < i; cpt++ {
		var word string
		_, err = fmt.Scanf("%s\n", &word)

		if err != nil {
			panic(err)
		}

		if len(word) > 10 {
			d := len(word) - 2
			fmt.Printf("%v%d%v\n", string(word[0]), d, string(word[len(word)-1]))

		} else {
			fmt.Println(word)
		}
	}

}
