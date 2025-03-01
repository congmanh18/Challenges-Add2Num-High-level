package bigsum

import (
	"log"
	"strconv"
	"strings"
)

type BigNumber struct{}

func (b *BigNumber) Sum(x, y string) string {
	log.Printf("Thực hiện phép tính cộng 2 số %s và %s", x, y)
	var result strings.Builder
	carry := 0

	// lấy độ dài của x và y
	nx, ny := len(x), len(y)

	// so sánh độ dài của x và y để lấy độ dài lớn nhất
	maxLen := nx
	if ny > nx {
		maxLen = ny
	}

	// khởi tạo vòng lặp
	for i := 0; i < maxLen; i++ {
		// khởi tạo x và y tại vị trí i
		digitX := 0
		digitY := 0

		if nx-i-1 >= 0 {
			digitX = int(x[nx-i-1] - '0')
		}

		if ny-i-1 >= 0 {
			digitY = int(y[ny-i-1] - '0')
		}

		sum := digitX + digitY + carry

		beforeCarry := carry
		carry = sum / 10
		result.WriteString(strconv.Itoa(sum % 10))
		log.Printf("Bước %d: Lấy %d cộng với %d được %d. Cộng tiếp với nhớ %d được %d. Lưu %d vào kết quả. Nhớ %d",
			i+1,
			digitX,
			digitY,
			digitX+digitY,
			beforeCarry,
			sum,
			sum%10,
			carry)
	}

	if carry > 0 {
		result.WriteString(strconv.Itoa(carry))
		log.Printf("Bước %d: Nhớ %d được lưu vào kết quả", maxLen+1, carry)
	}

	log.Println("--------------------------------")
	return reverse(result.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
