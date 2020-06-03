package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

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

func TestStringSliceSort(t *testing.T) {
	g := NewGomegaWithT(t)

	var genreSlice sort.StringSlice
	genreSlice = strings.Split("Short|Drama|Fantasy", "|")
	genreSlice.Sort()

	g.Expect(len(genreSlice)).Should(Equal(3))
	g.Expect(genreSlice[0]).Should(Equal("Drama"))
	g.Expect(genreSlice[1]).Should(Equal("Fantasy"))
	g.Expect(genreSlice[2]).Should(Equal("Short"))
}

func TestExtractYears(t *testing.T) {
	g := NewGomegaWithT(t)

	var i int64 = 1365029107
	ratingTime := time.Unix(i, 0)
	ratingTimeStr := ratingTime.Format(time.RFC3339Nano)
	fmt.Println(ratingTimeStr)
	yr := ratingTime.Year()
	fmt.Println(yr)
	g.Expect(1).Should(Equal(1))
}
