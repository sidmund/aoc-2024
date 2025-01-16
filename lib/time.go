package lib

import (
	"log"
	"time"
)

func Measure(start time.Time, task string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", task, elapsed)
}
