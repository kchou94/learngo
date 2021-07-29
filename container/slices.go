package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// fmt.Println("arr[2:6] =", arr[2:6])
	// fmt.Println("arr[:6] =", arr[:6])
	// fmt.Println("arr[2:] =", arr[2:])
	// fmt.Println("arr[:] =", arr[:])
	// slice 是对 arr 的一个 view

	s1 := arr[2:]
	fmt.Println(s1)
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	s2 := arr[:]
	fmt.Println(s2)
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s3 := arr2[2:6]
	s4 := s3[3:5]
	// s5 := s4[2:3]
	fmt.Println(s3, s4)
	// fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	// fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))

	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	fmt.Println(s5, s6, s7, arr2)
	// s6 s7 view different array
}
