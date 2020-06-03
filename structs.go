package main

// GenreCount struct
type GenreCount struct {
	Genre string
	Count int
}

// GenreCounts is a slice of GenreCount
type GenreCounts []GenreCount

// Len func
func (c GenreCounts) Len() int { return len(c) }

// Swap func
func (c GenreCounts) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

// ByMostPopular sorts by most mentioned count
type ByMostPopular struct{ GenreCounts }

// Less func
func (s ByMostPopular) Less(i, j int) bool {
	return s.GenreCounts[i].Count > s.GenreCounts[j].Count
}

// Result struct
type Result struct {
	Year int
	GenreCounts
}

// Results struct
type Results []Result

// Len func
func (r Results) Len() int { return len(r) }

// Swap func
func (r Results) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

// Less sort by year in ascending order, will need sort.Reverse() to make it in descending order
func (r Results) Less(i, j int) bool {
	return r[i].Year < r[j].Year
}
