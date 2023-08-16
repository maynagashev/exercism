package main
import (
	"fmt"
	"strconv"
)

// 26712 -+--
func PlusMinus(num int) (resOps string) {
	strNum := strconv.Itoa(num)
	runes := []rune(strNum)
	firstDigit := int(runes[0]-'0')
	tail := strNum[1:]

	// check minus branch
	resOps, resSum := calcTail("-", tail, "", firstDigit)
	if resSum == 0 {
		return resOps
	}

	// check plus branch (if somehow first plus possible in the case)
	resOps, resSum = calcTail("+", tail, "", firstDigit)
	if (resSum == 0) {
		return resOps
	}

	return "not possible";
}

func calcTail(op string, strNumTail string, resOps string, resSum int) (string, int) {
	if len(strNumTail) == 0 {
		return resOps, resSum
	}

	runes := []rune(strNumTail)
	digit := int(runes[0]-'0')
	// fmt.Printf("calcTail%v (%v) firstDigit=%v resOps=%v resSum=%v \n", op, strNumTail, digit, resOps, resSum)

	switch op {
	case "-":
		resOps += "-"
		resSum -= digit
	case "+":
		resOps += "+"
		resSum += digit
	}

	tail := strNumTail[1:]
	if len(tail) == 0 {
		return resOps, resSum
	}

	// first try minus branch (recursion)
	minusOps, minusSum := calcTail("-", tail, resOps, resSum)
	if minusSum == 0 {
		return minusOps, minusSum
	}

	// if not yet found try plus branch (recursion)
	return calcTail("+", tail, resOps, resSum)
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(PlusMinus(readline()))

}