package main

import "fmt"

func isPalindrome(s string) bool {
	is := true

	for i := range s {
		if s[i] == s[len(s)-(i+1)] {
			continue
		} else {
			is = false
			break
		}
	}

	return is
}

func verifyPalindrome(s string, verifier func(string) bool, onTrue func(), onFalse func()) {
	if verifier(s) {
		onTrue()
	} else {
		onFalse()
	}
}

func printEsPalindromo() {
	fmt.Println("Es palindromo")
}

func printNoEsPalindromo() {
	fmt.Println("No es palindromo")
}

func main() {
	var arr [4]int

	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	fmt.Println(arr)

	slice := []int{1, 2, 3}

	fmt.Println(slice[:2])
	fmt.Println(slice[1:2])
	fmt.Println(slice[1:])

	fmt.Println("Length", len(slice))
	fmt.Println("Capacity", cap(slice))

	otherNumbers := []int{8, 9, 10}

	slice = append(slice, otherNumbers...)
	fmt.Println(slice)

	for i, v := range slice {
		fmt.Println(i, v)
	}

	verifyPalindrome("radar", isPalindrome, printEsPalindromo, printNoEsPalindromo)
	verifyPalindrome("amor a roma", isPalindrome, printEsPalindromo, printNoEsPalindromo)
	verifyPalindrome("casaca", isPalindrome, printEsPalindromo, printNoEsPalindromo)
}
