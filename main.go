package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// keygen steps
// generate a random 24 digit number
// generate and apply a valid checksum
// jumble the 24 numbers in a preset pattern
// assemble the numbers in groups
// pack, shift and split the groups
// pass the result through a crypt table generating a limited 16 character ASCII set

var cipher string = "FMT4BIPW7ELSZAHOV6DKRY9GNU5CJQX8"

func main() {
	fmt.Println("generating seed data...")
	seed := calculateRandomSeed()
	fmt.Println(seed)

	checksum := calculateChecksumFromSeed(seed)
	fmt.Println(checksum)

	validKey := applyChecksumToSeed(seed, checksum)
	fmt.Println(validKey)

	scrambledKey := byteShiftKey(validKey)
	fmt.Printf("Scrambled / Byte shifted key: %s\n", scrambledKey)

	validRegistrationCode := calculateKeyCharsFromNumTrioGroups(scrambledKey)
	fmt.Printf("Generated Shark MX ESN: %s\n", validRegistrationCode)
}

func calculateRandomSeed() string {
	rand.Seed(time.Now().UnixNano())

	str := ""
	for i := 0; i < 24; i++ {
		// 0 is not a valid number to use here to avoid it
		str += strconv.Itoa(rand.Intn(8) + 1)
	}

	return str
}

func calculateChecksumFromSeed(seed string) string {
	sum := 0
	zero := "0"
	for i := 0; i < 11; i++ {
		// odd numbers
		a := (charCodeAt(seed, 2 * i) - charCodeAt(zero, 0)) * 2
		if a > 9 {
			a -= 9
		}

		sum += a

		// even numbers
		sum += charCodeAt(seed, 2 * i + 1) - charCodeAt(zero, 0)
	}

	sum %= 256

	tens := 0
	units := 0
	if sum != 0 {
		units = int(math.Ceil(float64(sum) / 10)) % 10
		tens = (10 - (sum % 10)) % 10
	}

	return fmt.Sprintf("%d%d", tens, units)
}

func applyChecksumToSeed(seed string, checksum string) string {
	return fmt.Sprintf("%s%s", seed[0:len(seed) - 2], checksum)
}

func byteShiftKey(key string) string {
	str := key

	temp := getCharInStringFromPos(str, 24)
	str = setCharInStringAtPos(str, getCharInStringFromPos(str, 1), 24)
	str = setCharInStringAtPos(str, temp, 1)

	temp = getCharInStringFromPos(str, 22)
	str = setCharInStringAtPos(str, getCharInStringFromPos(str, 3), 22)
	str = setCharInStringAtPos(str, temp, 3)

	temp = getCharInStringFromPos(str, 16)
	str = setCharInStringAtPos(str, getCharInStringFromPos(str, 9), 16)
	str = setCharInStringAtPos(str, temp, 9)

	temp = getCharInStringFromPos(str, 15)
	str = setCharInStringAtPos(str, getCharInStringFromPos(str, 10), 15)
	str = setCharInStringAtPos(str, temp, 10)

	temp = getCharInStringFromPos(str, 13)
	str = setCharInStringAtPos(str, getCharInStringFromPos(str, 12), 13)
	str = setCharInStringAtPos(str, temp, 12)

	return str
}

func calculateKeyCharsFromNumTrioGroups(scrambledKey string) string {
	result := ""

	for i := 0; i < 23; i += 3 {
		trio, _ := strconv.Atoi(scrambledKey[i:i+3])
		binary10BitNumStr := strconv.FormatInt(int64(trio), 2)

		binary5bitNum2Str := binary10BitNumStr[len(binary10BitNumStr) - 5: len(binary10BitNumStr)]
		binary5bitNum1Str := binary10BitNumStr[0:len(binary10BitNumStr)-len(binary5bitNum2Str)]

		snum1, _ := strconv.Atoi(binary5bitNum1Str)
		snum2, _ := strconv.Atoi(binary5bitNum2Str)

		num1 := binaryToDecimal(snum1)
		num2 := binaryToDecimal(snum2)

		cipherChar1 := getCipherCharAtPos(num1)
		cipherChar2 := getCipherCharAtPos(num2)

		fmt.Printf("trio: %d converted to binary is: %s. Split up 10-bit number to 5-bit numbers..\n5-bit num 1: %05s = %d -> cipher at index: %s, 5-bit num 2: %s = %d -> cipher at index: %s\n\n", trio, binary10BitNumStr, binary5bitNum1Str, num1, cipherChar1, binary5bitNum2Str, num2, cipherChar2)

		result += cipherChar1 + cipherChar2
	}

	return result
}

func binaryToDecimal(num int) int {
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0{
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder * int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}

func getCipherCharAtPos(pos int) string {
	return cipher[pos:pos+1]
}

func charCodeAt(s string, n int) int {
	i := 0
	for _, r := range s {
		if i == n {
			return int(r)
		}
		i++
	}
	return 0
}

func getCharInStringFromPos(str string, pos int) string {
	return str[pos-1:pos]
}

func setCharInStringAtPos(str string, replacement string, pos int) string {
	return str[:pos-1] + string(replacement) + str[pos:]
}
