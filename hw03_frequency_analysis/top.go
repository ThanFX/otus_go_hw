package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const MAX_RES_COUNT int = 10

func minDigit(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Top10(text string) []string {
	type sortedMap struct {
		k string
		v int
	}
	var (
		words []string       = make([]string, 0)
		wm    map[string]int = make(map[string]int)
		sm    []sortedMap    = make([]sortedMap, 0)
		out   []string       = make([]string, 0)
	)

	words = strings.Fields(text)

	// Подсчитали количество повторений слов в тексте
	for _, w := range words {
		if value, ok := wm[w]; ok {
			wm[w] = value + 1
		} else {
			wm[w] = 1
		}
	}

	for k, v := range wm {
		sm = append(sm, sortedMap{k, v})
	}
	// Сортируем мапу с количеством слов по количеству DESC и лексиграфически ASC в случае равенства количества
	sort.Slice(sm, func(i, j int) bool {
		if sm[i].v == sm[j].v {
			return sm[i].k < sm[j].k
		}
		return sm[i].v > sm[j].v
	})
	// Берём для вывода отсортированный список слов, но не больше допустимого
	for _, v := range sm[:minDigit(len(sm), MAX_RES_COUNT)] {
		out = append(out, v.k)
	}
	return out
}
