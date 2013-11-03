package main

import (
    "sort"
)

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *News) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(news []News) {
    ps := &NewsSorter{
        news: news,
        by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
    }
    sort.Sort(ps)
}

// NewsSorter joins a By function and a slice of Planets to be sorted.
type NewsSorter struct {
    news []News
    by      func(p1, p2 *News) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *NewsSorter) Len() int {
    return len(s.news)
}

// Swap is part of sort.Interface.
func (s *NewsSorter) Swap(i, j int) {
    s.news[i], s.news[j] = s.news[j], s.news[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *NewsSorter) Less(i, j int) bool {
    return s.by(&s.news[i], &s.news[j])
}