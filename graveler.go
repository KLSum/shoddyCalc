package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

const numTrials = 1000000 //Number of trials
const paraNeeded = 177    //Number of paralysis procs needed to succeed
const attNeeded = 231     //Number of total attacks needed

// Could be calculated while running but it's quicker to just do an array lookup
var bitMaskMap = [4]byte{0, 3, 15, 63}

func attackNTimes(n int) int {
	//(n+3)/4 makes battleArr always round up to the nearest byte
	battleArr := make([]byte, (n+3)/4)
	rand.Read(battleArr)
	count := 0

	//This masks n*2 bits and makes them all 1's and basically ignores them since we're looking for 0's
	battleArr[0] = battleArr[0] | bitMaskMap[n%4]

	//This loop battleArr which has roughly 1/4 of the bytes compared to the amount of attacks needed
	//Each byte has 8 bits and thus 4 pairs of bits, each pair has 4 combinations and we only care about 00
	for i := 0; i < len(battleArr); i++ {
		if battleArr[i]&byte(0b00000011) == 0 {
			count++
		}
		if battleArr[i]&byte(0b00001100) == 0 {
			count++
		}
		if battleArr[i]&byte(0b00110000) == 0 {
			count++
		}
		if battleArr[i]&byte(0b11000000) == 0 {
			count++
		}
	}
	return count
}

// Gives the higher of 2 values
func max(curr, new int) int {
	if curr < new {
		return new
	}
	return curr
}

func main() {
	cycleStart := time.Now()
	attempts := 0         //Amount of simulations attempted
	maxParalyzeCount := 0 //Highest amount of paralysis procs
	for attempts < numTrials {
		//Run a simulation and take the highest result
		maxParalyzeCount = max(maxParalyzeCount, attackNTimes(attNeeded))

		//Break if we ever meet the goal (We won't)
		if maxParalyzeCount >= paraNeeded {
			break
		}
		attempts++
	}

	//Print results
	fmt.Println("Amount of time taken to simulate: ", time.Now().Sub(cycleStart))
	fmt.Println("Highest paralysis count: ", maxParalyzeCount)
	fmt.Println("Number of attempts: ", attempts)
}
