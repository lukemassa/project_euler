package main

import (
	"fmt"
	"log"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/constraints"
)

const gridNums = 4

const numElementsPermutation = 4

func fib(nums chan *big.Int) {
	i := big.NewInt(0)
	j := big.NewInt(1)
	for {
		nums <- j
		i, j = j, i.Add(j, i)

	}

}

func isIntPalindrome(n int) bool {
	return isPalindrome(strconv.Itoa(n))
}

func isPalindrome(s string) bool {
	strLen := len(s)
	//fmt.Printf("Starting on %s, which is %d long\n", numAsStr, strLen)
	for i := 0; i < strLen/2; i++ {
		//fmt.Printf("   Comparing %c and %c\n", numAsStr[i], numAsStr[strLen-1-i])
		if s[i] != s[strLen-1-i] {
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

func bigFactorial(n int64) *big.Int {
	if n < 0 {
		log.Fatalf("Can't get factorial of neg num %d", n)
		return nil
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
func factorial(n int) int64 {
	if n < 0 {
		panic(fmt.Sprintf("Can't get factorial of neg num %d", n))
	}
	if n < 2 {
		return 1
	}
	if n > 20 {
		panic(fmt.Sprintf("Can't get factorial of num %d, try using bigFactorial", n))
	}
	ret := int64(1)
	for i := int64(2); i <= int64(n); i++ {
		ret *= i
	}
	return ret
}

func isAbundant[T constraints.Integer](n T) bool {
	return sum(divisors(n)) > n
}

func abundants(max int) {
	abundantNumbers := mapset.NewSet[int]()
	for i := 0; i <= max; i++ {
		if isAbundant(i) {
			abundantNumbers.Add(i)
		}
	}
	sumOfAbundants := make([]int, 0)
	for i := 1; i < max; i++ {
		canBeWritten := false
		for summand := 1; summand <= i/2; summand++ {
			if !abundantNumbers.Contains(summand) {
				continue
			}
			if !abundantNumbers.Contains(i - summand) {
				continue
			}
			canBeWritten = true

		}
		if !canBeWritten {
			fmt.Printf("%d cannot be written as sum of abundant\n", i)
			sumOfAbundants = append(sumOfAbundants, i)
		}
	}
	fmt.Println(sum(sumOfAbundants))
	//return false
}

func BubbleSort[T constraints.Ordered](array []T) []T {
	for i := 0; i < len(array)-1; i++ {
		fmt.Println(i)
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func nthBiggestElement[T constraints.Ordered](arr []T, n int) T {
	// Return the n-th largest element in a not-necessarily sorted array
	sorted := make([]T, len(arr))
	copy(sorted, arr)
	start := time.Now()
	sorted = BubbleSort(sorted)
	fmt.Printf("%d,%f\n", len(arr), time.Since(start).Seconds())
	return sorted[len(sorted)-n]

}
func sortArrAfterPoint(arr []int, k int) []int {

	endPart := arr[k:]
	sort.Ints(endPart)
	return append(arr[:k], endPart...)
}
func nextPermutation(arr []int) ([]int, bool) {
	// 012
	// 021
	// 102
	// 120
	// 201
	// 210

	// 0123
	// 0132
	// 0213
	// 0231
	// 0312
	// 0321
	// 1023

	// Find the "pivot"
	n := len(arr)
	pivot := n - 2
	for ; pivot > 0; pivot-- {

		if arr[pivot] < arr[pivot+1] {
			break
		}
	}
	if pivot == -1 {
		return arr, true
	}
	arr = sortArrAfterPoint(arr, pivot+1)
	swapWithPivot := pivot + 1
	for ; swapWithPivot < n; swapWithPivot++ {
		if arr[pivot] < arr[swapWithPivot] {
			break
		}
	}
	if swapWithPivot == n {
		return arr, true
	}
	arr[pivot], arr[swapWithPivot] = arr[swapWithPivot], arr[pivot]
	arr = sortArrAfterPoint(arr, pivot+1)
	return arr, false
}
func showPermutations(n int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	count := 1
	for {
		fmt.Println(count, arr)
		if count == 1_000_000 {
			break
		}
		tmpArr, done := nextPermutation(arr)
		arr = tmpArr
		if done {
			break
		}
		count += 1
	}

}

func unitReciprocalCycles(n int) int {
	i := 0
	const runLength = 3
	seen := make(map[string]int)
	fmt.Printf("1/%d: ", n)
	numerator := 1
	current := ""
	for {
		i += 1
		quotient, remainder := (10*numerator)/n, (10*numerator)%n
		current = fmt.Sprintf("%s%d", current, quotient)
		if len(current) > runLength {
			current = current[1:]
		}
		place, found := seen[current]
		if found {
			return (i - place)
		}
		if len(current) == runLength {
			seen[current] = i
		}

		//fmt.Print(current, " ")
		//fmt.Print(quotient)
		//fmt.Print(seen)
		if remainder == 0 {
			return 0
		}
		numerator = remainder
	}
}

func distinctPowers(maxInt int64) int {
	sofar := make([]big.Int, 0)
	one := big.NewInt(1)
	//min := big.NewInt(2)
	max := big.NewInt(maxInt)

	for a := big.NewInt(2); a.Cmp(max) != 1; a.Add(a, one) {
		for b := big.NewInt(2); b.Cmp(max) != 1; b.Add(b, one) {
			res := big.NewInt(0)
			res.Exp(a, b, nil)
			//fmt.Printf("%v ^ %v = %v\n", a, b, res)
			found := false
			for i := 0; i < len(sofar); i++ {
				if sofar[i].Cmp(res) == 0 {
					found = true
					break
				}
			}
			if !found {
				sofar = append(sofar, *res)
			}

		}
	}
	return len(sofar)
}
func distinctPowers2(max uint16) int {
	decomps := make([]DecomposedInteger, max+1)
	for i := uint16(2); i <= max; i++ {
		decomps[i] = NewDecomposedInteger(i)
	}
	seen := mapset.NewSet[DecomposedInteger]()
	seenChannel := make(chan DecomposedInteger)
	for a := uint16(2); a <= max; a++ {
		go func(a uint16) {
			fmt.Printf("%d/%d\n", a, max)
			for b := uint16(2); b <= max; b++ {
				if b%100 == 0 {
					fmt.Printf("%d is on %d\n", a, b)
				}
				res := decomps[a].Pow(b)
				//fmt.Printf("%v (%v) ^ %v = %v\n", a, decomps[a], b, res)
				seenChannel <- res

			}
		}(a)
	}
	done := uint16(0)
	for res := range seenChannel {
		done += 1
		if done == (max-1)*(max-1) {
			break
		}
		//fmt.Println(res)
		seen.Add(res)
	}
	close(seenChannel)
	// Check to see if anything's still writing to the channel
	time.Sleep(100 * time.Millisecond)
	return seen.Cardinality()
}

// How many of the 1 to max numbers were *not*
// included by smaller powers
func howManyNew(power, min, max int) int {
	fmt.Printf("STARTING: power is %d\n", power)
	alreadyHit := make([]bool, max*power+1)
	for lowerPower := 1; lowerPower < power; lowerPower++ {
		//fmt.Println("   sWorking on", lowerPower)
		for i := 0; i <= max; i++ {
			alreadyHit[i*lowerPower] = true
		}
	}
	// Pretend values below min were "hit" already
	for i := 0; i < min; i++ {
		alreadyHit[i] = true
	}
	//fmt.Println(alreadyHit)
	ret := 0
	for i := 0; i <= max*power; i += power {
		//fmt.Printf("  Looking at %d, which is of course %v\n", i, alreadyHit[i])
		if !alreadyHit[i] {
			ret += 1
		}
	}
	return ret
}

func distinctPowers3(max int) int {
	ret := 0
	min := 2
	for i := min; i <= max; i++ {
		// power := largestPowerOfPrime(int64(i))
		power := gcdOfPowersOfPrime(int64(i))
		toAdd := howManyNew(power, min, max)
		fmt.Printf("%d has power %d, so adding %d\n", i, power, toAdd)
		ret += toAdd

	}
	return ret
}

func largestPowerOfPrime(n int64) int {
	if n == 12 {
		return 1
	}
	freq := make(map[int64]int)
	maxPower := 0
	// Could be sqrt?
	for i := int64(2); i <= n; i++ {
		//fmt.Println(n, i)
		if n%i == 0 {
			n /= i
			freq[i] += 1
			if freq[i] > maxPower {
				maxPower = freq[i]
			}

			i = int64(1)
			continue
		}
	}
	return maxPower
}

func smallestPowerOfPrime(n int64) int {
	//if n == 5184 {
	//	return 1
	//}
	freq := make(map[int64]int)

	// Could be sqrt?
	for i := int64(2); i <= n; i++ {
		//fmt.Println(n, i)
		if n%i == 0 {
			n /= i
			freq[i] += 1

			i = int64(1)
			continue
		}
	}
	//fmt.Println(freq)
	minPower := -1
	for _, power := range freq {
		if minPower == -1 || power < minPower {
			minPower = power
		}
	}
	return minPower
}

func gcdOfPowersOfPrime(n int64) int {
	freq := make(map[int64]int)

	// Could be sqrt?
	for i := int64(2); i <= n; i++ {
		//fmt.Println(n, i)
		if n%i == 0 {
			n /= i
			freq[i] += 1

			i = int64(1)
			continue
		}
	}
	powers := make([]int, len(freq))
	i := 0
	for _, value := range freq {
		powers[i] = value
		i += 1
	}
	return multiGCD(powers)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func multiGCD(nums []int) int {
	if len(nums) == 0 {
		panic("No numbers to take gcd of!")
	}
	if len(nums) == 1 {
		return nums[0]
	}
	runningGCD := GCD(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		runningGCD = GCD(runningGCD, nums[i])
	}
	return runningGCD
}

func sumDigitPower(n, power int) int {
	s := strconv.Itoa(n)
	sum := 0
	for _, letter := range s {
		digit := int(letter - '0')
		//fmt.Println(digit)
		res := 1
		for i := 0; i < power; i++ {
			res *= digit
		}
		//fmt.Println("bcomes", res)
		sum += res
	}
	return sum
}

func numberOfCoins(denom int) *big.Int {
	cache := make(map[int]*big.Int)
	return numCoins(denom, &cache)
}

func numCoins(denom int, cache *map[int]*big.Int) *big.Int {

	if denom == 0 {
		//fmt.Println("      ZEERO")
		return big.NewInt(0)
	}
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	sum := big.NewInt(0)
	for _, coin := range coins {
		if denom < coin {
			continue
		}

		addition, found := (*cache)[denom-coin]

		if found {
			//	fmt.Printf("    For coin %d, value %d found in cache %d\n", coin, denom-coin, addition)
		}
		if !found {
			fmt.Printf("    For coin %d, denom %d calling func(%d)\n", coin, denom, denom-coin)
			addition = numCoins(denom-coin, cache)
			fmt.Printf("    For value %d, cache is being set to %d\n", denom-coin, addition)
			var newAddition big.Int
			newAddition.Set(addition)
			(*cache)[denom-coin] = &newAddition
		}
		sum.Add(sum, big.NewInt(1))
		sum.Add(sum, addition)

		fmt.Printf("  For coin %d, addition is %d, so sum is %d\n", coin, addition, sum)
	}
	fmt.Printf("Asking about %d, answer is %d\n", denom, sum)
	return sum
}

func isPowerTwo(n uint64) bool {
	if n == 0 {
		return false
	}
	for n > 1 {
		// if n is odd
		if n&1 != 0 {
			return false
		}
		n = n >> 1
	}
	return true
}
func getDigits(i int) []int {
	digitsAsString := strconv.Itoa(i)
	numDigits := len(digitsAsString)
	digits := make([]int, numDigits)
	for i := 0; i < numDigits; i++ {
		digits[i] = int(digitsAsString[i] - '0')
	}
	return digits
}
func rotate(i int) []int {

	digits := getDigits(i)
	numDigits := len(digits)

	ret := make([]int, numDigits)
	for i := 0; i < numDigits; i++ {
		value := 0
		power := 1
		for k := 0; k < numDigits; k++ {
			position := (i + k) % numDigits
			value += digits[position] * power
			power *= 10
		}

		ret[i] = value
	}
	return ret

}

func base2Representation(n int) string {
	if n == 0 {
		return "0"
	}
	digits := make([]bool, 0)
	for n > 0 {
		digits = append(digits, n&1 == 0)
		n /= 2
	}
	ret := strings.Builder{}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] {
			ret.WriteRune('0')
		} else {
			ret.WriteRune('1')
		}
	}
	return ret.String()
}

func hasUniqueDigits(expectedLength int, num ...int) bool {
	seen := make(map[byte]struct{})
	length := 0
	for k := 0; k < len(num); k++ {
		digits := strconv.Itoa(num[k])
		length += len(digits)

		for i := 0; i < len(digits); i++ {
			if digits[i] == '0' {
				return false
			}
			_, found := seen[digits[i]]
			if found {
				return false
			}
			seen[digits[i]] = struct{}{}
		}
	}
	if expectedLength != -1 && length != expectedLength {
		return false
	}
	return true
}

func getDigitLength(n int) int {
	if n < 10 {
		return 1
	}
	if n < 100 {
		return 2
	}
	if n < 1000 {
		return 3
	}
	if n < 10_000 {
		return 4
	}
	if n < 100_000 {
		return 5
	}
	if n < 1_000_000 {
		return 6
	}
	if n < 10_000_000 {
		return 7
	}
	panic(fmt.Sprintf("can't handle %d", n))
}
func findTriangleTriples(sum int) int {
	numSolutions := 0
	for a := 1; a <= sum/3; a++ {
		for b := a; b <= (sum-a)/2; b++ {
			c := sum - (a + b)
			if !(a <= b && b <= c) {
				fmt.Printf("Not true that %d<=%d<=%d\n", a, b, c)
			}
			if a*a+b*b != c*c {
				continue
			}
			//fmt.Println(a, b, c)
			numSolutions += 1
		}
	}
	return numSolutions
}

func numNDigintNumbers(n int) int {
	power := 1
	for i := 1; i < n; i++ {
		power *= 10
	}
	return 9 * power
}

func nthDigitOfConcatenatedDecimal(n int) int {
	// First, determine digit length block where n is
	digitLength := 1
	startDigitBlock := 1
	digitLengthSize := 1
	for {
		nextStartDigitBlock := startDigitBlock + numNDigintNumbers(digitLength)*digitLength
		if nextStartDigitBlock > n {
			break
		}
		digitLength += 1
		digitLengthSize *= 10
		startDigitBlock = nextStartDigitBlock

	}
	//return startDigitBlock, digitLength
	entryInBlock := (n - startDigitBlock) / digitLength
	startingPoint := entryInBlock*digitLength + startDigitBlock
	entry := digitLengthSize + entryInBlock
	entryAsString := strconv.Itoa(entry)
	digit := int(entryAsString[n-startingPoint] - '0')
	return digit
}

func main() {
	fmt.Println(nthDigitOfConcatenatedDecimal(922337203685477580))
}
