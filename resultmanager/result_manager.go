package resultmanager

import (
	"sort"
)

// Capturer interface
type Capturer interface {
	// Capture which year the movie genre is found
	Capture(year int, genre string)
}

// ResultGetter interface
type ResultGetter interface {
	// GetResult returns Results, slice of yearly genre counts Result, that can be sorted by caller. Results are sorted in ascending year,
	// genre counts are not sorted to allow the caller to sort them based on their needs
	GetResult() Results
}

// ResultManager interface
type ResultManager interface {
	Capturer
	ResultGetter
}

// Implements the ResultManager interface
type resultManager struct {
	yearToGenreCounts map[int]map[string]int
}

// NewResultManager returns a new result manager for capturing data and reporting results
func NewResultManager() ResultManager {
	return &resultManager{
		yearToGenreCounts: make(map[int]map[string]int, 0),
	}
}

// Capture which year the movie genre is found
func (m *resultManager) Capture(year int, genre string) {
	if genreToCount, ok := m.yearToGenreCounts[year]; ok {
		if _, iok := genreToCount[genre]; iok {
			genreToCount[genre]++
		} else {
			genreToCount[genre] = 1
		}
	} else {
		m.yearToGenreCounts[year] = map[string]int{genre: 1}
	}
}

// GetResult returns Results, slice of yearly genre counts Result, that can be sorted by caller. Results are sorted in ascending year,
// genre counts are not sorted to allow the caller to sort them based on their needs
func (m *resultManager) GetResult() Results {
	var results Results
	for year, genreToCount := range m.yearToGenreCounts {
		var genreCounts GenreCounts
		for genre, count := range genreToCount {
			genreCounts = append(genreCounts, GenreCount{Genre: genre, Count: count})
		}

		results = append(results,
			Result{
				Year:        year,
				GenreCounts: genreCounts,
			})
	}

	// by default sort results by year ascending
	sort.Sort(results)
	return results
}
