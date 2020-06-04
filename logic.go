package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/samuel-foo/movietweets/fileparser"
	resultm "github.com/samuel-foo/movietweets/resultmanager"
)

const (
	// GenreDelimiter is the character that separate each genre in a genre list
	GenreDelimiter = "|"

	// FileNameMovies is the name of movies data file
	FileNameMovies = "movies.dat"

	// FieldCountMovies is the number of fields in the movies data file
	FieldCountMovies = 3

	// FileNameRatings is the name of ratings data file
	FileNameRatings = "ratings.dat"

	// FieldCountRatings is the number of fields in the ratings data file
	FieldCountRatings = 4
)

// Parse movies data file and returns movie id to genre map
func getMovieIDToGenreMap(dataDir string) (map[string]string, error) {
	movieID2Genre := make(map[string]string, 0)

	// For each row sorts the genre list alphabetically because they could be out of order
	var action fileparser.RowAction = func(fields []string) error {
		var genreSlice sort.StringSlice
		// field 2 is genres
		genreSlice = strings.Split(fields[2], GenreDelimiter)
		genreSlice.Sort()
		// field 0 is movie id
		movieID2Genre[fields[0]] = strings.Join(genreSlice, GenreDelimiter)
		return nil
	}

	if err := fileparser.ParseFile(filepath.Join(dataDir, FileNameMovies), FieldCountMovies, action); err != nil {
		return nil, err
	}

	return movieID2Genre, nil
}

// Parse ratings data file and returns result
func processRatings(dataDir string, currentYear int, movieID2Genre map[string]string) (resultm.Results, error) {

	// Calculate minimum date past decade
	minAcceptedDate := time.Date((currentYear - 10), time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Minimum accepted date: %v\n", minAcceptedDate)
	// Calculate the minimum accepted seconds since epoch, ratings older than this will be ignored
	minAcceptedDateSec := minAcceptedDate.Unix()

	// For each rating record increments count per genre
	rm := resultm.NewResultManager()
	var action fileparser.RowAction = func(fields []string) error {
		if secondsSinceEpoch, err := strconv.ParseInt(fields[3], 10, 64); err == nil {
			if secondsSinceEpoch >= minAcceptedDateSec {
				// lookup genre by movie id
				if genre, ok := movieID2Genre[fields[1]]; ok {
					rm.Capture(time.Unix(secondsSinceEpoch, 0).Year(), genre)
				} else {
					log.Printf("Warning, movie id cannot be found in movies.dat: %s, skipping...\n", fields[1])
				}
			}
		} else {
			log.Printf("Warning, invalid seconds since epoch %s: %s...skipping\n", fields[3], err)
		}
		return nil
	}

	if err := fileparser.ParseFile(filepath.Join(dataDir, FileNameRatings), FieldCountRatings, action); err != nil {
		return nil, err
	}

	return rm.GetResult(), nil
}
