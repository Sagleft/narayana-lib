package narayanalib

import (
	"fmt"
	"math/rand"
	"time"
)

func GetInt64() int64 {
	var seed int64 = 0 // TODO

	s := rand.NewSource(seed)
	return readMantra(s)
}

// returns M from [0,maxN)
func getIntFromSeed(seed int64, maxN int) int {
	rand.Seed(seed)
	return rand.Intn(maxN)
}

func sleepNanoseconds(nanoseconds int) {
	time.Sleep(time.Nanosecond * time.Duration(nanoseconds))
}

func readMantra(src rand.Source) (value int64) {
	for i := 1; i <= 1000; i++ {
		value = src.Int63()
		src.Seed(value)

		fmt.Print("om namah narayana ")
		sleepNanoseconds(getIntFromSeed(value, 100))
	}

	return value
}
