package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compressString(s string) string {
	chars := strings.Split(s, "")
	dst, slow := 0, 0

	// Пробегаемся по всем символом строки
	for fast := 0; fast <= len(chars); fast++ {
		// Если дошли до конца строки или последовательность подряд идущих одинаковых символов закончилась
		if fast == len(chars) || chars[slow] != chars[fast] {
			chars[dst] = chars[slow]
			dst++
			count := fast - slow

			// Если кол-во подряд идущих символов больше 1 (есть что сокращать)
			if count > 1 {
				chars[dst] = strconv.Itoa(count)
				dst++
			}
			slow = fast
		}
	}

	return strings.Join(chars[:dst], "")
}

func main() {
	fmt.Println(compressString("aaabbbccccc"))
	fmt.Println(compressString("abc"))
	fmt.Println(compressString("c"))
	fmt.Println(compressString("a"))
	fmt.Println(compressString("abbbc"))
	fmt.Println(compressString("aac"))
	fmt.Println(compressString("abccc"))
}
