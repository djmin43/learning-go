package main

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	type firstPerson struct {
		name string
		age  int
	}

	type secondPerson struct {
		name string
		age  int
	}

	type thirdPerson struct {
		age  int
		name string
	}
}
