package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var cipher string = "FMT4BIPW7ELSZAHOV6DKRY9GNU5CJQX8"

// generate a random 24 digit number
// generate and apply a valid checksum
// jumble the 24 numbers in a preset pattern
// assemble the numbers in groups
// pack, shift and split the groups
// pass the result through a crypt table generating a limited 16 character ASCII set

func main() {
	seedData := calculateRandomSeedArray()
	seedData = applyChecksum(seedData, calculateChecksum(seedData))
	seedData = byteShiftSeed(seedData)
	calculateCharactersFromNumTrios(seedData)
}

func calculateRandomSeedArray() [24]byte {
	fmt.Println("generating seed data...")
	rand.Seed(time.Now().UnixNano())

	var seedData [24]byte
	for i := 0; i < 24; i++ {
		// 0 is not a valid number to use here to avoid it
		seedData[i] = byte(rand.Intn(8) + 1)
	}

	fmt.Println(seedData)
	return seedData
}

func calculateChecksum(seedData [24]byte) uint8 {
	fmt.Println("checksum calculation")
	// because of 0-index i numbers will look off
	// Odd digits 1/24, 3/24 etc are doubled, if > 10 then 9 is subtracted. These values are added together.

	var oddNumbersCalc byte
	for i := 0; i <= 21; i+=2 {
		var num byte = seedData[i] * seedData[i]
		if num > 10 {
			num -= 9
		}

		oddNumbersCalc += num
	}

	// Even digits 2/24, 4/24 etc are just added together
	var evenNumbersCalc byte
	for i := 1; i <= 21; i+=2 {
		evenNumbersCalc += seedData[i]
	}

	checksum := evenNumbersCalc + oddNumbersCalc

	fmt.Printf("checksum: %d\n", checksum)
	return uint8(checksum)
}

func applyChecksum(seed [24]byte, checksum uint8) [24]byte {
	if checksum < 10 || checksum > 99 {
		panic("try again")
	}

	slice := strconv.Itoa(int(checksum))
	digit1, _ := strconv.Atoi(slice[:1])
	digit2, _ := strconv.Atoi(slice[1:])

	seed[22] = uint8(digit1)
	seed[23] = uint8(digit2)

	fmt.Println(seed)
	return seed
}

// jumble the 24 numbers in a preset pattern
// 1. byte 1 swap with byte 24
// 2. byte 3 swap with byte 22
// 3. byte 9 swap with byte 16
// 4. byte 10 swap with byte 15
// 5. byte 12 swap with byte 13

func byteShiftSeed(seed [24]byte) [24]byte {
	temp := seed[23]
	seed[23] = seed[0]
	seed[0] = temp

	temp = seed[21]
	seed[21] = seed[2]
	seed[2] = temp

	temp = seed[15]
	seed[15] = seed[8]
	seed[8] = temp

	temp = seed[14]
	seed[14] = seed[9]
	seed[9] = temp

	temp = seed[12]
	seed[12] = seed[11]
	seed[11] = temp

	fmt.Println(seed)
	return seed
}

func calculateCharactersFromNumTrios(seed [24]byte) string {
	result := ""
	bytesToStr := ""
	for k := 0; k < len(seed); k++ {
		bytesToStr += strconv.Itoa(int(seed[k]))
	}

	for i := 0; i < 23; i += 3 {
		trio, _ := strconv.Atoi(bytesToStr[i:i+3])
		binary10BitNumStr := strconv.FormatInt(int64(trio), 2)


		binary5bitNum2Str := binary10BitNumStr[len(binary10BitNumStr) - 5: len(binary10BitNumStr)]
		binary5bitNum1Str := binary10BitNumStr[0:len(binary10BitNumStr)-len(binary5bitNum2Str)]

		fmt.Printf("trio: %d, binary: %s, 5-bit num 1: %05s, 5-bit num 2: %s\n", trio, binary10BitNumStr, binary5bitNum1Str, binary5bitNum2Str)

		snum1, _ := strconv.Atoi(binary5bitNum1Str)
		snum2, _ := strconv.Atoi(binary5bitNum2Str)

		num1 := binaryToDecimal(snum1)
		num2 := binaryToDecimal(snum2)

		result += getCipherCharAtPos(num1)
		result += getCipherCharAtPos(num2)
	}

	fmt.Printf("Calculated Gameshark MX ESN Registration Code: %s\n", result)
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