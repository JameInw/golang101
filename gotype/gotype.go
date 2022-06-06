package gotype

import "fmt"

// Basic Type

// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte //alias for uint8

// rune // alias for int32
// represents a Unicode code point

// float32 float64

// complex64 complex128

//    ┌────────┬────────────────────────────────┐
//    │        │                                │
//    │   s    │                                │
//    │        │ 0x12345                        │
//    ├────────┤    ▲  ┌─────────────────────┐  │
//    │        │    │  │  "hello world       │  │
//    │ string │    │  └─────────────────────┘  │
//    │        │    │                           │
//    └────────┴────┼───────────────────────────┘
//                  │
//                  │
//                  └────────────┐
// 								 │
//								 │
//    ┌────────┬─────────────────┼──────────────┐
//    │        │                 │              │
//    │   P    │                 │              │
//    │        │ 0x33221         │      		│
//    │		   │				 │				│
//    ├────────┤       ┌─────────┴───────────┐  │
//    │        │       │       0x12345       │  │
//    │ *string│       └─────────────────────┘  │
//    │        │                                │
//    └────────┴────────────────────────────────┘
//
// Pointer
//

// var s string
// var p *string

// P = &s
// s = "hello"
// *p += " world"

func Double(n *int) {
	*n = *n * 2
	fmt.Println("Double : ", *n)
}
func AppendGreeting(s *string) {
	*s = "Hi, " + *s
	fmt.Println(*s)
}

// Zero Value
// n_numbers = 0
// s_string = ""
// b_boolean = false
// p_pointer = nil

// Array
// var fourNum [4]int //ให้ขนาดเท่าไรก็ใช้ได้เท่านั้นไม่สามารถเพิ่มภายหลังได้

// fourNum[0] = 1
// fourNum[1] = 3

// Slice
// var nums []int
// nums = make([]int, 4)

// nums[0] = 1
// nums[2] = 2
// nums = append(nums, 20)

func Abc() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	s = s[0:3]
	fmt.Println(s)
	// [apple banana coconut]
}
func Efg() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	s = s[4:7]
	fmt.Println(s)
	// [elderberries figs grapes]
}
func Cde() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	s = s[2:5]
	fmt.Println(s)
	// [coconut durian elderberries]
}
