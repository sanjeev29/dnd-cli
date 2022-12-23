package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	dice := flag.String("d", "d6", "The type of dice to roll. Format: dX where X is an integer; Default: d6")
	numRoll := flag.Int("n", 1, "The number of die to roll; Default: 1")
	sum := flag.Bool("s", false, "Get the sum of all the dice rolls")
	advantage := flag.Bool("adv", false, "Roll the dice with advantage")
	disadvantage := flag.Bool("dadv", false, "Roll the dice with disadvantage")

	flag.Parse()

	matched, _ := regexp.Match("^d\\d+", []byte(*dice))

	if matched == false {
		log.Fatal("Invalid dice format.")
		return
	}

	rolls := rollDice(dice, numRoll)
	printDice(rolls)

	if *sum {
		diceSum := sumDice(rolls)
		fmt.Printf("The sum of the dice %d\n", diceSum)
	}

	if *advantage {
		roll := RollWithAdvantage(rolls)
		fmt.Printf("The roll with advantage was %d\n", roll)
	}

	if *disadvantage {
		roll := RollWithDisdvantage(rolls)
		fmt.Printf("The roll with disadvantage was %d\n", roll)
	}

}

func rollDice(dice *string, times *int) []int {
	var rolls []int

	diceSides := (*dice)[1:]
	d, err := strconv.Atoi(diceSides)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < *times; i++ {
		rolls = append(rolls, rand.Intn(d)+1)
	}

	return rolls
}

func printDice(rolls []int) {
	for i, dice := range rolls {
		fmt.Printf("Roll %d was %d\n", i+1, dice)
	}
}

func sumDice(rolls []int) int {
	sum := 0

	for _, dice := range rolls {
		sum += dice
	}

	return sum
}

func RollWithAdvantage(rolls []int) int {
	sort.Ints(rolls)

	return rolls[len(rolls)-1]
}

func RollWithDisdvantage(rolls []int) int {
	sort.Ints(rolls)

	return rolls[0]
}
