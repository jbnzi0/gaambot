package main

import "math/rand"

func getRandomAdjective() string {
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

	min := 0
	max := len(adjectives)
	return adjectives[rand.Intn(max-min+1)+min]
}