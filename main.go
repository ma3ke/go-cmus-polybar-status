package main

import (
	"fmt"
	"os/exec"
	"strings"
	"log"
	"math"
)

func main() {
	status := getStatus()

	fmt.Println(status[2])

	fmt.Println(parseDuration(371))
}

func getStatus() []string {
	cmd := exec.Command("cmus-remote", "-Q")

	status, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(status))

	output := strings.Split(string(status), "\n")

	return output
}

func parseDuration(seconds int) (int, int){
	if seconds < 0 {
		return 0, 0
	} else {
		minutes := float64(seconds) / 60.0
		min := math.Floor(minutes)
		sec := math.Floor((minutes - min) * 60)

		return int(min), int(sec)
	}
}