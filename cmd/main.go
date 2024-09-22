package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

	// slices
	// t := []byte("string")

	// // get the lenght of the slice and the items in the slice
	// fmt.Println(len(t), t)

	// // gets the 3rd element in the slice
	// fmt.Println(t[2])

	// //gets from the start of the slice till the elemnt eofre th index 2
	// fmt.Println(t[:2])

	// //get from the index 3 and till the index 4 (5 means it is not included)
	// fmt.Println(t[3:5])

	// //gets from the index 2 till the end of the slice
	// fmt.Println(t[2:])

	// //append to the slice, this creates a new slice with the new element, as we need to have a new underlying array of bytes to store the new element,
	// // there was enough memory
	// t = append(t, '!')
	// fmt.Println(t)

	scan := bufio.NewScanner(os.Stdin)
	//create a map to store the words and their count (like a dictionary in C#)
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)

	// scan the input and count the words, if the word was there before incremenet it's count
	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words found")

	// creates a new struct (like a class in C#)
	type kv struct {
		key string
		val int
	}

	// create a slice of this struct to store the key value pairs
	var ss []kv

	for k, v := range words {
		// add a new struct in the slice for each key value pair in the words range
		ss = append(ss, kv{k, v})
	}

	sort.SliceStable(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, kv := range ss {
		fmt.Println(kv.key, "appears", kv.val, "times")
	}
}
