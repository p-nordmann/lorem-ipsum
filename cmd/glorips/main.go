package main

import (
	"fmt"
	"math/rand"

	loremipsum "github.com/p-nordmann/lorem-ipsum"
	pforest "github.com/p-nordmann/prefix-forest"
	flag "github.com/spf13/pflag"
)

const maxDepth = 5

var numParagraphs int

func init() {
	flag.IntVarP(&numParagraphs, "count", "n", 3, "count of paragraphs to generate")
}

func main() {
	flag.Parse()

	f := pforest.New(maxDepth)
	f.Learn(loremipsum.LiberPrimus)

	context := "\n\n"
	lastTokenIsNewLine := false
	for count := 0; count < numParagraphs; {
		density := f.Predict(context)
		c := rune(sampleToken(density))
		if c == '\n' {
			if lastTokenIsNewLine {
				count++
				lastTokenIsNewLine = false
			} else {
				lastTokenIsNewLine = true
			}
		}
		context = context + string(c)
	}
	fmt.Printf(context[2 : len(context)-1])
}

func sampleToken(density map[pforest.Token]int) pforest.Token {
	var tokens []pforest.Token
	var values []int
	for token, value := range density {
		tokens = append(tokens, token)
		values = append(values, value)
	}
	index := sample(values)
	return tokens[index]
}

// Samples an index according to the density.
//
// TODO implement less naive approach
func sample(density []int) int {
	var total int
	var stairs []int
	for _, d := range density {
		stairs = append(stairs, total)
		total += d
	}
	dice := rand.Intn(total)
	for k, stair := range stairs {
		if dice < stair {
			return k - 1
		}
	}
	return len(stairs) - 1
}
