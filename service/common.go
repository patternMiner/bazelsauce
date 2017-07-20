package service

import (
	"sort"
)

// Common types
type RankMap map[string]int
type StringSlice []string
type StringSliceMap map[string]StringSlice
type StringSet map[string]bool
type StringSetMap map[string]StringSet

// A data structure to hold a key/value pair.
type Pair struct {
	Key string
	Value int
}
// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) SortByValue() {sort.Sort(sort.Reverse(p))}


func makePairList(m RankMap) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	return p
}

func (slice StringSlice) makeStringSet() StringSet {
	ss := make(StringSet)
	for _, s := range slice {
		ss[s] = true
	}
	return ss
}

func (stringSet StringSet) append (s string) {
	stringSet[s] = true
}

func (stringSet StringSet) merge (otherSet StringSet) {
	for other := range otherSet {
		stringSet[other] = true
	}
}

func (stm StringSetMap) stringSet (k string) StringSet {
	s, ok := stm[k]
	if !ok {
		s = make(StringSet)
		stm[k] = s
	}
	return s
}