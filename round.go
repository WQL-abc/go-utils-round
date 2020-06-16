package round

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Do(f float64) float64 {
	var (
		runes   []rune
		roundUp bool
	)
	ff := strconv.FormatFloat(f, 'f', -1, 64)
	sp := strings.Split(ff, ".")
	for _, r := range []rune(reverse(sp[1])) {
		if r >= 53 {
			roundUp = true
		}
		if roundUp {
			r++
		}
		if r < 53 {
			roundUp = false
		}
		if r >= 58 {
			r = 48
		}
		runes = append(runes, r)
	}
	s := fmt.Sprintf("%s.%s", sp[0], reverse(string(runes)))
	p, err := strconv.ParseFloat(s, 10)
	if err != nil {
		panic(err)
	}
	return math.Round(p)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
