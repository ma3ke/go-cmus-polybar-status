package main

import (
	"fmt"
	"os/exec"
	"strings"
	"log"
)

func main() {
	status := getStatus()

	fmt.Println(status[2])
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
