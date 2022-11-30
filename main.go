package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

const gridNums = 4

func fib(nums chan int) {
	i := 0
	j := 1
	for {
		i, j = j, i+j
		nums <- j
	}

}

func genPrimes(nums chan int64) {
	nums <- 2
	nums <- 3
	for i := int64(6); ; i += 6 {
		if isPrime(i - 1) {
			nums <- i - 1
		}
		if isPrime(i + 1) {
			nums <- i + 1
		}
	}
}

func factor(n int64) []int64 {
	factors := make([]int64, 0)
	for {
		factor := nextFactor(n)
		if factor == -1 {
			factors = append(factors, n)
			return factors
		}
		factors = append(factors, factor)
		n = n / factor
	}
}

func divisors[T constraints.Integer](n T) []T {
	divisors := make([]T, 0)
	for i := T(1); i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}

	}
	return divisors
}

func isPrime(n int64) bool {
	factor := nextFactor(n)
	return factor == -1
}

func nextFactor(n int64) int64 {
	//fmt.Printf("Working on %d\n", n)
	max := int64(math.Sqrt(float64(n)))
	for i := int64(2); i <= max; i++ {
		if n%i == 0 {
			//fmt.Printf("   Found %d\n", i)
			return i
		}
	}
	return -1
}

func isPalindrome(n int) bool {
	numAsStr := strconv.Itoa(n)
	strLen := len(numAsStr)
	//fmt.Printf("Starting on %s, which is %d long\n", numAsStr, strLen)
	for i := 0; i < strLen/2; i++ {
		//fmt.Printf("   Comparing %c and %c\n", numAsStr[i], numAsStr[strLen-1-i])
		if numAsStr[i] != numAsStr[strLen-1-i] {
			return false
		}
	}
	return true
}

func parseGrid(digits string) [][]int {
	ret := make([][]int, 0)
	for _, line := range strings.Split(digits, "\n") {
		line = strings.Trim(line, "\t")
		row := make([]int, 0)
		for _, entry := range strings.Split(line, " ") {
			digit, err := strconv.Atoi(entry)
			if err != nil {
				panic(err)
			}
			row = append(row, digit)
		}
		ret = append(ret, row)
	}
	return ret
}

func getFreq[T comparable](sortedList []T) map[T]int {
	freq := make(map[T]int)
	for i := 0; i < len(sortedList); i++ {
		// Skip ahead until they aren't equal anymore
		start := i
		for ; i < len(sortedList)-1 && sortedList[i] == sortedList[i+1]; i++ {
		}
		freq[sortedList[start]] = i - start + 1
	}

	return freq
}

func numDivisors(n int64) int {
	// Special case; 1 doesn't have any prime factors so the below calculation won't work
	if n == 1 {
		return 1
	}
	ret := 1
	factors := factor(n)
	for _, freq := range getFreq(factors) {
		// Each prime contributes 0, 1, ..., p to a given divisor
		ret *= freq + 1
	}
	return ret
}

func sumDivisors[T constraints.Integer](n T) T {
	return sum(divisors(n))
}
func sum[T constraints.Integer](arr []T) T {
	// Special case; 1 doesn't have any prime factors so the below calculation won't work
	var sum T
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func adjDigits(digits string, chanDigits chan [gridNums]int) {

	grid := parseGrid(digits)
	for _, line := range grid {
		fmt.Println(line)
	}
	y := len(grid)
	x := len(grid[0])

	// Horizontal
	for i := 0; i < y; i++ {

		for j := 0; j <= x-gridNums; j++ {
			var candidate [gridNums]int
			for k := 0; k < gridNums; k++ {
				candidate[k] = grid[i][j+k]
			}
			chanDigits <- candidate

		}
	}
	// Vertical
	for i := 0; i < y; i++ {

		for j := 0; j <= x-gridNums; j++ {
			var candidate [gridNums]int
			for k := 0; k < gridNums; k++ {
				candidate[k] = grid[j+k][i]
			}
			chanDigits <- candidate

		}
	}
	// Diagonal up from left to right
	for i := gridNums - 1; i < y; i++ {

		for j := 0; j <= x-gridNums; j++ {
			var candidate [gridNums]int
			for k := 0; k < gridNums; k++ {
				candidate[k] = grid[j+k][i-k]
			}
			chanDigits <- candidate

		}
	}
	// Diagonal down from left to right
	for i := 0; i <= y-gridNums; i++ {

		for j := 0; j <= x-gridNums; j++ {
			var candidate [gridNums]int
			for k := 0; k < gridNums; k++ {
				candidate[k] = grid[j+k][i+k]
			}
			chanDigits <- candidate

		}
	}
	close(chanDigits)
}

func collatzLength(n int64) int64 {
	i := int64(1)
	bool := 1
	fmt.Println(bool)
	for ; n != int64(1); i++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return i
}

func next(path []int) ([]int, bool) {
	max := 1
	for i := 0; i < len(path); i++ {
		path[i] += 1
		if path[i] <= max {
			return path, true
		}
		path[i] = 0
	}
	return path, false
}

func numberInWords(n int) string {
	earlyNumbers := [20]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}
	tensStrings := [10]string{
		"",
		"",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}
	if n < 20 {
		return earlyNumbers[n]
	}
	if n < 100 {
		ones, tens := n%10, n/10
		suffix := ""
		if ones != 0 {
			suffix = "-" + earlyNumbers[ones]
		}
		return tensStrings[tens] + suffix
	}
	if n < 1000 {
		underAHundred, hundreds := n%100, n/100
		suffix := ""
		if underAHundred != 0 {
			suffix = " and " + numberInWords(underAHundred)
		}
		return earlyNumbers[hundreds] + " hundred" + suffix
	}
	if n < 1_000_000 {
		underAThousand, thousands := n%1000, n/1000
		suffix := ""
		if underAThousand != 0 {
			suffix = " " + numberInWords(underAThousand)
		}
		return numberInWords(thousands) + " thousand" + suffix
	}
	panic(fmt.Sprintf("Can't handle %d", n))
}

func writtenOutLength(n int) int {
	word := numberInWords(n)
	word = strings.Replace(word, " ", "", -1)
	word = strings.Replace(word, "-", "", -1)

	return len(word)
}

func parseTriangle(tri string) [][]int {
	ret := make([][]int, 0)
	for _, line := range strings.Split(tri, "\n") {
		line = strings.Trim(line, " \t")
		lineOfStr := strings.Split(line, " ")
		lineOfNums := make([]int, len(lineOfStr))
		for i := 0; i < len(lineOfStr); i++ {
			num, err := strconv.Atoi(lineOfStr[i])
			if err != nil {
				panic(err)
			}
			lineOfNums[i] = num
		}
		ret = append(ret, lineOfNums)
	}
	return ret
}

func valueOfPath(path []int, tri [][]int) int {
	sum := tri[0][0]
	index := 0
	for i := 0; i < len(path); i++ {
		index += path[i]
		entry := tri[i+1][index]
		sum += entry
	}
	return sum
}
func bruteForceTriangle(tri [][]int) int {
	longest := 0
	path := make([]int, len(tri)-1)
	success := true
	for success {
		value := valueOfPath(path, tri)
		if value > longest {
			longest = value
		}
		path, success = next(path)

	}
	return longest
}

func solveTriangle(tri [][]int) int {
	height := len(tri)
	choices := make([][]int, height)

	lastRowChoices := make([]int, height)
	for i := 0; i < height; i++ {
		lastRowChoices[i] = tri[height-1][i]

	}
	choices[height-1] = lastRowChoices

	for i := height - 2; i >= 0; i-- {
		row := make([]int, i+1)
		for j := 0; j < len(row); j++ {
			var below int
			if choices[i+1][j] > choices[i+1][j+1] {
				below = choices[i+1][j]
			} else {
				below = choices[i+1][j+1]
			}
			row[j] = tri[i][j] + below

		}
		choices[i] = row
	}
	for i := 0; i < len(choices); i++ {
		fmt.Println(choices[i])
	}
	return choices[0][0]
}

func dedupe[T comparable](x []T) []T {
	exists := make(map[T]bool)
	for i := 0; i < len(x); i++ {
		exists[x[i]] = true
	}
	ret := make([]T, 0)
	for k, _ := range exists {
		ret = append(ret, k)
	}
	return ret
}

func factorial(n int64) *big.Int {
	if n < 0 {
		log.Fatalf("Can't get factorial of neg num %d", n)
	}
	if n < 2 {
		return big.NewInt(1)
	}
	ret := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		ret.Mul(ret, big.NewInt(i))
	}
	return ret
}

//func nextStartingAndOffset(starting, offset int) (int, int) {

//}
func main() {
	a := make([]int, 5)
	a[0] = 1
	a[1] = 4
	a[2] = 4
	fmt.Println(slices.Contains(a, 6))
}
