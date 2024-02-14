package main

type car struct {
	brand string
	year  string
}

func main() {
	raptor := car{brand: "Ford", year: "2021"}
	mustang := car{brand: "Ford", year: "1964"}

	println(raptor.brand, raptor.year)
	println(mustang.brand, mustang.year)

	var unknown car
	unknown.brand = "Unknown"
	unknown.year = "Unknown"

	println(unknown.brand, unknown.year)
}
