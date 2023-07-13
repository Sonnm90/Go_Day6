package stringFunc

import (
	"fmt"
	"strings"
)

func Demo() {
	str := "Hello, World!"

	fmt.Println(len(str))                               // Độ dài của chuỗi
	fmt.Println(strings.Contains(str, "Hello"))         // Kiểm tra xem chuỗi có chứa "Hello" không
	fmt.Println(strings.Split(str, ", "))               // Tách chuỗi thành mảng các phần tử
	fmt.Println(strings.ToUpper(str))                   // Chuyển đổi chuỗi thành chữ hoa
	fmt.Println(strings.Replace(str, "Hello", "Hi", 1)) // Thay thế "Hello" bằng "Hi" trong chuỗi
}
