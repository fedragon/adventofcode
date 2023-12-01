package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	PacketLength  = 4
	MessageLength = 14
)

func main() {
	f, err := os.Open("../data/day06")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()

	start := 0
	startOfPacketMarker, startOfMessageMarker := PacketLength, MessageLength
	packetMarkerFound, messageMarkerFound := false, false

	for {
		if packetMarkerFound && messageMarkerFound {
			break
		}

		if isMarker(line[start:start+PacketLength], PacketLength) {
			packetMarkerFound = true
		}

		if isMarker(line[start:start+MessageLength], MessageLength) {
			messageMarkerFound = true
		}

		if !packetMarkerFound {
			startOfPacketMarker++
		}

		if !messageMarkerFound {
			startOfMessageMarker++
		}

		start++
	}

	fmt.Println(startOfPacketMarker)
	fmt.Println(startOfMessageMarker)
}

func isMarker(input string, length int) bool {
	m := make(map[rune]struct{})

	for _, r := range input {
		m[r] = struct{}{}
	}

	return len(m) == length
}
