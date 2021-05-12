package main

import "fmt"
import "os"

const MAX_DEG = 360
const HOUR_PER_ROUND = 12
const MINUTES_PER_HOUR = 60
const MINUTES_PER_ROUND = MINUTES_PER_HOUR * HOUR_PER_ROUND
const DEGREES_PER_HOUR = MAX_DEG / HOUR_PER_ROUND

// у нас целые числа и соответственно нам нужно обратное отношение сколько минут в одном градусе
const MINUTES_IN_DEGREE = MINUTES_PER_ROUND / MAX_DEG

// поиск часов и минут по смещению часовой стрелки
func main() {

	var d uint
	fmt.Scan(&d)
	if d <= 0 || d >= MAX_DEG {
		fmt.Fprintf(os.Stderr, "Invalid degree %d should be (0..360)", d)
		os.Exit(1)
	} else {
		hours := d / DEGREES_PER_HOUR
		minutes := d % DEGREES_PER_HOUR * MINUTES_IN_DEGREE
		fmt.Printf("It is %d hours %d minutes.", hours, minutes)
	}
}
