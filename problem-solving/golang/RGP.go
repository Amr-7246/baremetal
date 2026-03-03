package main
import "fmt"

func rgb(r, g, b int) string {
	hexaNums := "0123456789ABCDEF"
	result := ""
	colors := []int{r, g, b}

	for _, num := range colors {
		if num > 255 {
			num = 255
		} else if num < 0 {
			num = 0
		}
		quotient := num / 16
		remainder := num % 16
	
		result += string(hexaNums[quotient]) + string(hexaNums[remainder])
	}

	return result
}

func main() {
	fmt.Println(rgb(0, 224, 231))
}