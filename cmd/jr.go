package cmd

import (
	"fmt"
	"time"

	"github.com/charmbracelet/huh"
)

func main() {
	currentTime := time.Now()

	var entry string
	huh.NewInput().
		Title("jr input change this haha").
		Value(&entry).
		Run() // this is blocking..

	fmt.Println(entry)
}
