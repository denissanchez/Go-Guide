package main

func vocalsCounter(input string) (int, int, int, int, int) {
	vocals := []int{0, 0, 0, 0, 0}

	for _, char := range input {
		switch char {
		case 'a':
			vocals[0]++
		case 'e':
			vocals[1]++
		case 'i':
			vocals[2]++
		case 'o':
			vocals[3]++
		case 'u':
			vocals[4]++
		}
	}

	return vocals[0], vocals[1], vocals[2], vocals[3], vocals[4]
}

func main() {
	input := "Hello, World!"

	println("Input:", input)

	a, e, i, o, u := vocalsCounter(input)

	println("a:", a)
	println("e:", e)
	println("i:", i)
	println("o:", o)
	println("u:", u)
}
