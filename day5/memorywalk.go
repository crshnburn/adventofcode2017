package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)

func MemoryWalk(numbers []int) int {
	steps := 0
	
	for pos := 0; pos >= 0 && pos < len(numbers); {
		newPos := pos + numbers[pos]
		steps++
		numbers[pos] ++
		pos = newPos
	}

	return steps
}

func MemoryWalk2(numbers []int) int {
	steps := 0
	
	for pos := 0; pos >= 0 && pos < len(numbers); {
		newPos := pos + numbers[pos]
		steps++
		if numbers[pos] >= 3 {
			numbers[pos]--
		} else {
			numbers[pos] ++
		}
		pos = newPos
	}

	return steps
}

func main(){
	file, err := os.Open("./memory.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	var memory []int
    for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		memory = append(memory, i)
	}

	// fmt.Println("Escape in ", MemoryWalk(memory), " steps")
	fmt.Println("Escape 2 in ", MemoryWalk2(memory), " steps")
}