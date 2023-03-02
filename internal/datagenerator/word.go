package datagenerator

import (
	"math/rand"
	"time"
)

func GetRandomAdjective() string {
	rand.Seed(time.Now().UnixNano())

	adjectives := []string{
		"authentic",
		"initial",
		"aboriginal",
		"beginning",
		"first",
		"infant",
		"opening",
		"pioneer",
		"primary",
		"starting",
		"archetypal",
		"autochthonous",
		"commencing",
		"early",
		"elementary",
		"embryonic",
		"first-hand",
		"genuine",
		"inceptive",
		"introductory",
		"prime",
		"primeval",
		"primitive",
		"primordial",
		"pristine",
		"prototypal",
		"rudimental",
		"rudimentary",
		"underivative",
		"underived",
		"random",
		"crazy",
	}

	max := len(adjectives) - 1
	return adjectives[rand.Intn(max+1)]
}
