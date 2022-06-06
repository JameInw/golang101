package mysyntax

import (
	"fmt"
)

// รูปแบบการประกาศ variables
var i int
var s string
var ok bool
var i2 int = 14
var s2 string = "hello2"
var ok2 bool = true

// รูปแบบของ Type inference
var i3 = 14
var s3 = "hello3"
var ok3 = true

// i4 := 14			// การประกาศ variables นี้สามารถใช้ได้ใน func เท่านั้น
// s4 := "hello"	// การประกาศ variables นี้สามารถใช้ได้ใน func เท่านั้น
// ok4 := true		// การประกาศ variables นี้สามารถใช้ได้ใน func เท่านั้น

// Constant เป็นตัวแปรที่เปลี่ยนค่าไม่ได้แล้ว ควรจะใส่ค่าเบื้องต้นให้เลย
const defaultValue int = 1
const defaultTitle = "Go"
const defaultStatus = true

// การใช้งาน Function

// No return
func Greeting(firstName, lasstName string) {
	fmt.Println("Hello", firstName, lasstName)
}

// Return 1 value
func GreetingWithAge(name string, age uint) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
}

// Return 3 value
func GreetingWithAgeAndDrink(name string, age uint, drink string) string {
	return fmt.Sprintf("Hello, %s. You are %d years old and your favorite drink is %s.", name, age, drink)
}

// การใช้งาน IF Else

func IsOdd(n int) bool { // ถ้าใส่เลขคี่ไปให้แสดงเป็น true หรือไม่ใช่ให้แสดง false
	if n%2 != 0 {
		return true
	} else {
		return false
	}
}

// การใช้งาน Loop
func SumOffFirst(n int) int {
	sum := 0
	for i := n; i > 0; i-- {
		sum = sum + i
	}
	return sum
}
