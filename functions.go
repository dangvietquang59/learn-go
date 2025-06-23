package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

func CalculateSQRT(num int) {
	if num < 0 {
		return
	}
	fmt.Printf("%.4f", math.Sqrt(float64(num)))
}
func CheckPalindrome(num int) bool {
	origin := num
	reversed := 0

	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}
	return reversed == origin
}
func CountPerfectNum(from int, to int) int {
	count := 0
	for i := from; i <= to; i++ {
		if CheckPerfectNumber(i) {
			count++
		}
	}
	return count
}
func DescendingSlice(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return slice
}

func AscendingSlice(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

func FilterOddAndEven(nums []int) ([]int, []int) {
	odds := []int{}
	evens := []int{}

	for _, num := range nums {
		if CheckEvenOrOdd(num) {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}

	return odds, evens
}
func PrintReverseNum(num int) int {
	reversed := 0
	signed := 1
	if num < 0 {
		signed = -1
		num = -num
	}

	for num > 0 {
		reversed = reversed*10 + num%10
		num = num / 10
	}
	return reversed * signed
}
func CountEvenInArray(array []int) int {
	count := 0
	for _, num := range array {
		if CheckEvenOrOdd(num) {
			count++
		}
	}
	return count
}
func CheckAscendingArray(array []int) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i] >= array[i+1] {
			return false
		}
	}
	return true
}

func TotalPrimeFrom1ToN(n int) int {
	total := 0
	for i := 2; i <= n; i++ {
		if CheckPrime(i) {
			total += i
		}
	}
	return total
}
func SnakeToCamel(str string) string {
	parts := strings.Split(str, "_")
	if len(parts) == 0 {
		return str
	}
	camel := parts[0]
	for _, part := range parts[1:] {
		if part == "" {
			continue
		}
		camel += strings.Title(part)
	}
	return camel
}
func TriangleBy3Sides(a, b, c float64) (float64, float64) {
	perimeter := a + b + c
	s := perimeter / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return perimeter, area
}
func ConvertMinutesToHHMM(num int) string {
	hours := num / 60
	minutes := num % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func CheckValidTriangle(num1 int, num2 int, num3 int) bool {
	return (num1+num2 > num3) && (num1+num3 > num2) && (num2+num3 > num1)
}
func CheckLeapYear(num int) bool {
	return (num%4 == 0 && num%100 != 0) || num%400 == 0
}
func TotalTimesOfNumber(num int) int {
	count := 1
	if num < 0 {
		num = -num
	}
	for num >= 10 {
		count *= num % 10
		num /= 10
	}
	return count
}
func CountNumberCharactor(num int) int {
	count := 1
	if num < 0 {
		num = -num
	}
	for num >= 10 {
		num /= 10
		count++
	}
	return count
}
func FindFirstAndLastNum(num int) (int, int) {
	if num < 0 {
		num = -num
	}

	last_number := num % 10
	first_number := num

	for first_number >= 10 {
		first_number /= 10
	}

	return first_number, last_number
}

func PrintMultiplicationTable() {
	for i := 2; i <= 9; i++ {
		fmt.Printf("Table %d \n", i)
		PrintMultiplicationTableByNumber(i)
	}
}
func PrintMultiplicationTableByNumber(num int) {
	for i := 1; i <= 10; i++ {
		result := num * i
		fmt.Printf("%d x %d = %d \n", num, i, result)
	}
}
func CalculateTotalOddAndEvenFrom1ToNumber(num int) (int, int) {
	total_odd := 0
	total_even := 0

	for i := 0; i <= num; i++ {
		if CheckEvenOrOdd(i) {
			total_even += i
		} else {
			total_odd += i
		}
	}

	return total_odd, total_even
}
func CalculateTotalSquareFrom1ToNumber(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i * i
	}
	return sum
}

func CheckPerfectSquareNumber(number int) bool {
	if number < 0 {
		return false
	}

	result := math.Sqrt(float64(number))
	return math.Pow(result, 2) == float64(number)
}
func CheckPerfectNumber(number int) bool {
	if number <= 1 {
		return false
	}
	sum := 0
	for i := 1; i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}
	return sum == number
}
func SumDigits(number int) int {
	if number < 0 {
		number = -number
	}
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}
	return sum
}

func ConvertSeconds(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	remainingSeconds := seconds % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, remainingSeconds)
}
func ReturnArgument(number int) int {
	if number == 0 {
		return 0
	} else {
		return number * -1
	}
}
func CalculateAVG3Numbers(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	if len(numbers) == 0 {
		return 0
	}

	return sum / len(numbers)
}
func CalculateTotalFrom1ToN(number int) int {
	total := 0
	for i := 1; i <= number; i++ {
		total += i
	}
	return total
}
func CheckNumberType(number int) string {
	if number == 0 {
		return "0"
	} else if number > 0 {
		return "+"
	} else {
		return "-"
	}
}
func CalculatePow2Numbers(number_1 int, number_2 int) float64 {
	pow := math.Pow(float64(number_1), float64(number_2))
	return pow
}
func SwapNumber(number_1 int, number_2 int) (int, int) {
	return number_2, number_1
}
func CalculateBMI(weight float64, height int) float64 {
	bmi := weight / math.Pow((float64(height)/100), 2)
	return bmi
}
func CountPrimeLessThanInput(number int) int {
	count := 0
	for i := 2; i <= number; i++ {
		if CheckPrime(i) {
			count++
		}
	}
	return count
}
func CheckPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
func CalculateTotalOdd(number_1 int, number_2 int) int {
	total := 0
	for i := number_1; i <= number_2; i++ {
		if !CheckEvenOrOdd(i) {
			total += i
		}
	}
	return total
}
func CalculateAge(year int) {
	result := time.Now().Year() - year
	fmt.Printf("Age "+"%d", result)
}
func EnterNameAndPrintHell(name string) {
	result := PrintHelloGo(name)
	fmt.Println(result)
}
func FindMaxAndMin(numbers ...int) (int, int) {
	if len(numbers) == 0 {
		panic("no numbers provided")
	}

	max := numbers[0]
	min := numbers[0]

	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

func CheckDivide3Or5(number int) int {
	if number%3 == 0 && number%5 == 0 {
		return 1
	} else if number%3 == 0 {
		return 2
	} else if number%5 == 0 {
		return 3
	}

	return 0
}
func CheckEvenOrOdd(number int) bool {
	return number%2 == 0
}
func TranferTemperature() {
	var celsius float64
	fmt.Print("Enter temperature in Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := celsius*9/5 + 32
	fmt.Printf("Temperature in Fahrenheit: %.2f°F\n", fahrenheit)
}
func CalculateTriangle() {
	var base, height float64
	fmt.Print("Enter base: ")
	fmt.Scan(&base)
	fmt.Print("Enter height: ")
	fmt.Scan(&height)

	area := (base * height) / 2
	fmt.Printf("Area of triangle: %.2f\n", area)
}
func CalculateCircle() {
	var radius float64
	fmt.Print("Enter radius: ")
	fmt.Scan(&radius)

	circumference := 2 * math.Pi * radius
	fmt.Printf("Circumference: %.2f\n", circumference)
}
func CalculateAreaOfRectangle() {
	var a, b int

	fmt.Println("Enter the first number:")
	fmt.Scan(&a)
	fmt.Println("Enter the second number:")
	fmt.Scan(&b)

	result := Calculate2Numbers("*", a, b)
	fmt.Printf("Result: %d\n", result)
}
func Calculate2Numbers(operator string, number1 int, number2 int) int {
	var result int

	switch operator {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		result = number1 / number2
	case "%":
		result = number1 % number2
	}

	return result
}
func PrintHelloGo(name string) string {
	result := "Hello" + " " + name
	return result
}
