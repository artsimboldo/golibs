package golibs

/*
	Fisher-Yates shuffle
	Uses an interface to bind custom data types
*/

import (
	"math/rand"
	"time"
)

type Interface interface {
	Len() int	
	Swap(i, j int)
}

func Shuffle(data Interface) {
	rand.Seed(time.Now().UTC().UnixNano())
	l := data.Len()
	for i := 0; i < l; i++ {
		j := rand.Intn(i + 1)
		data.Swap(i, j)
	}
}
