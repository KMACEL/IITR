package timop

import (
	"math/rand"
	"time"
)

// Random is
func Random(begin int, end int) int {
	source := rand.NewSource(time.Now().UnixNano())
	randVariable := rand.New(source)
	return randVariable.Intn((end - begin) + begin)
}
