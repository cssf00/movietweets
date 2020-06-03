package main

import (
	"sort"
	"testing"

	. "github.com/onsi/gomega"
)

func TestSortGenreCountsByMostPopular(t *testing.T) {
	g := NewGomegaWithT(t)

	gCounts := GenreCounts{
		GenreCount{Genre: "Horror", Count: 9},
		GenreCount{Genre: "Romance", Count: 4},
		GenreCount{Genre: "Thriller", Count: 15},
		GenreCount{Genre: "Kids and family", Count: 2},
	}

	sort.Sort(ByMostPopular{gCounts})

	g.Expect(gCounts[0]).Should(Equal(GenreCount{Genre: "Thriller", Count: 15}))
}

func TestSortResultsYearDescendingOrder(t *testing.T) {
	g := NewGomegaWithT(t)

	rs := Results{
		Result{Year: 2010},
		Result{Year: 2009},
		Result{Year: 2012},
	}
	sort.Sort(sort.Reverse(rs))

	g.Expect(rs[0]).Should(Equal(Result{Year: 2012}))
}
