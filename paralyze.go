package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

const numTrials = 1000000 // Number of trials
const paraNeeded = 500    // Number of paralysis procs needed to succeed
const attNeeded = 231     // Number of total attacks needed

// Could be calculated while running but it's quicker to just do an array lookup
var bitMaskMap = [4]byte{0, 3, 15, 63}

func attackNTimes(n int) int {
	battleArr := make([]byte, (n+3)/4)
	rand.Read(battleArr)
	count := 0
	battleArr[0] = battleArr[0] | bitMaskMap[n%4]
	for i := 0; i < len(battleArr); i++ {
		if battleArr[i]&byte(3) == 0 {
			count++
		}
		if battleArr[i]&byte(12) == 0 {
			count++
		}
		if battleArr[i]&byte(48) == 0 {
			count++
		}
		if battleArr[i]&byte(192) == 0 {
			count++
		}
	}
	return count
}

func max(curr, new int) int {
	if curr < new {
		return new
	}
	return curr
}

func main() {
	cycleStart := time.Now()
	attempts := 0
	maxParalyzeCount := 0
	for attempts < numTrials {
		maxParalyzeCount = max(maxParalyzeCount, attackNTimes(attNeeded))
		if maxParalyzeCount >= paraNeeded {
			break
		}
		attempts++
	}
	fmt.Println(time.Now().Sub(cycleStart))
	fmt.Println("Highest paralysis count: ", maxParalyzeCount)
	fmt.Println("Number of attempts: ", attempts)
}
