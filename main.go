package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args

	if x := len(args); x != 2 {
		x = x - 1 // args always includes the command / binary name.
		fmt.Printf("Expected one argument, got %d\n", x)
		os.Exit(1)
	}

	durStr := args[1]

	if err := sleep(durStr); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	os.Exit(0)
}

func sleep(durationStr string) error {
	// If durationStr converts to an int without error, that means
	// it is missing a duration suffix. Default to `s` for seconds.
	if _, err := strconv.Atoi(durationStr); err == nil {
		durationStr = fmt.Sprintf("%ss", durationStr)
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("invalid duration: %s: %v", durationStr, err)
	}
	time.Sleep(duration)
	return nil
}
