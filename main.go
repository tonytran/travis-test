package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/vsui/travis-test/mathutils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Not an integer")
			continue
		}
		pow, err := mathutils.Pow(2, n)
		if err != nil {
			fmt.Println("No negative numbers!")
			continue
		}
		if n == 12 {
			fmt.Println("YOU HIT THE JACKPOT!~")
			os.Exit(0)
		}
		fmt.Printf("2^%d = %d\n", n, pow)
	}
}
