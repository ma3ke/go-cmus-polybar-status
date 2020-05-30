package main

import (
	"fmt"
	"os/exec"
	"strings"
	"log"
	"math"
	"strconv"
)

func main() {
	status := getStatus()

	fmt.Println(status[2])

	stat := parseStatus(status)

	fmt.Println(parseDuration(stat.duration))
	fmt.Println(parseDuration(stat.position))

	fmt.Println(stat.artist, "\u2014", stat.title) // \u2014 represents an em dash.

	fmt.Println(progressIndicator(stat.duration, stat.position, 10))
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

type status struct {
	playing bool

	title string
	artist string
	album string

	duration int
	position int
}

func parseStatus(s []string) (status) {
	var stat status

	playing := strings.TrimPrefix(s[0], "status ")
	if playing == "playing" {
		stat.playing = true
	} else {
		stat.playing = false
	}

	stat.title = strings.TrimPrefix(s[4], "tag title ")
	stat.artist = strings.TrimPrefix(s[5], "tag artist ")
	stat.album = strings.TrimPrefix(s[6], "tag album ")

	var err1, err2 error

	stat.duration, err1 = strconv.Atoi(strings.TrimPrefix(s[2], "duration "))
	stat.position, err2 = strconv.Atoi(strings.TrimPrefix(s[3], "position "))

	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}

	return stat
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

func statusIndicator(playing bool) string {
	if playing {
		return ">"
	} else {
		return "\""
	}
}

func progressIndicator(dur, pos, len int) string {
	linechar := "-"
	pointerchar := "|"

	progress := float64(pos) / float64(dur)

	pre := strings.Repeat(linechar, int(progress * float64(len)))
	suf := strings.Repeat(linechar, int((1 - progress) * float64(len)))

	return pre + pointerchar + suf
}