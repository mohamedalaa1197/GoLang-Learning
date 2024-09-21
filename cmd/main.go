package main

func main() {
	// a := 2
	// b := 2.01
	// fmt.Printf("a: %T %v\n", a, a)
	// fmt.Printf("b: %T %v\n", b, b)

	// // implicit type conversion
	// a = int(b)
	// fmt.Printf("a: %T %v\n", a, a)

	// calculation example
	// var sum float64
	// var n int

	// for {
	// 	var val float64

	// 	_, err := fmt.Fscanln(os.Stdin, &val)
	// 	if err != nil {
	// 		break
	// 	}
	// 	sum += val
	// 	n++
	// }

	// if n == 0 {
	// 	fmt.Fprintln(os.Stderr, "no values")
	// }

	// fmt.Println("the avrage is:", sum/float64(n))

	//  string example
	// strings are immutable, it can't change once it is created
	// strings has descriptors that point to the underlying array of bytes, and also the length of the string
	// s := "hello" => s is a string descriptor that points to the underlying array of bytes
	// a := s[:1] => a is a new string descriptor that points to the same underlying array of bytes, but with a different length
	// d := s[:2] + "world!" :s[:1] => d is a new string descriptor that points to a new underlying array of bytes (doesn't point to the same bytes of s), with a different length
	// s+= "world!" => s is a new string descriptor that points to a new underlying array of bytes (doesn't point to the same bytes of s), with a different length, and the old s value is still there as there are variables points to it.

}
