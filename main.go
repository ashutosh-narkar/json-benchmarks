package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
)

func main() {
	analyzeAddString()
	analyzeAddSlice()
	analyzeAddObject()
}

func analyzeAddString() {
	for i := 0; i < 50; i++ {
		input := randomString(i + 1)
		bytes, err := json.Marshal(input)
		if err != nil {
			panic(err)
		}

		runTest(len(input), bytes)
	}
}

func analyzeAddSlice() {
	var input []string
	for i := 0; i < 50; i++ {
		input = append(input, strconv.Itoa(i+1))
		bytes, err := json.Marshal(input)
		if err != nil {
			panic(err)
		}

		runTest(len(input), bytes)
	}
}

func analyzeAddObject() {
	input := make(map[string]int)
	for i := 0; i < 50; i++ {
		input[strconv.Itoa(i+1)] = i + 1
		bytes, err := json.Marshal(input)
		if err != nil {
			panic(err)
		}

		runTest(len(input), bytes)
	}
}

func runTest(num int, input []byte) {
	var result interface{}

	before := getMemUsage()
	if err := json.Unmarshal(input, &result); err != nil {
		panic(err)
	}
	after := getMemUsage()

	fmt.Printf("%v\t%v\t%v\n", num, len(input), after.Alloc-before.Alloc)
}

func getMemUsage() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
