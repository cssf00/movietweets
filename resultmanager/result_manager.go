package resultmanager

import (
	"sort"
)

// Capturer interface
type Capturer interface {
	Capture(year int, genre string)
}

// ResultGetter interface
type ResultGetter interface {
	GetResult() Results
}

// ResultManager interface
type ResultManager interface {
	Capturer
	ResultGetter
}

// Implements the ResultManager interface
type resultManager struct {
	year2GenreCounts map[int]map[string]int
}

// NewResultManager returns a new result manager
func NewResultManager() ResultManager {
	return &resultManager{
		year2GenreCounts: make(map[int]map[string]int, 0),
	}
}

// Capture result
func (m *resultManager) Capture(year int, genre string) {
	if genre2Count, ok := m.year2GenreCounts[year]; ok {
		genre2Count[genre]++
	} else {
		m.year2GenreCounts[year] = map[string]int{genre: 1}
	}
}

// GetResult returns results in ascending year, genre counts are not sorted to allow the caller
// to sort them based on their needs
func (m *resultManager) GetResult() Results {
	var results Results
	for year, genre2Count := range m.year2GenreCounts {
		var genreCounts GenreCounts
		for genre, count := range genre2Count {
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
