package main

func main() {

	count := 10

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	increment(&count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

func increment(inc *int) {
	*inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "\tValue Points To[", *inc, "]")
}
