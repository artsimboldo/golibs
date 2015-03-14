package golibs

/*
	Shuffle(Interface)
	Implments Fisher-Yates random permutation of a finite set (slice)
	http://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	Use Interface to bind custom data types
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
