package movietweets

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

// GetMovieIDToGenreMap parses movies data file and returns a map of MovieID to Genre map
func GetMovieIDToGenreMap(dataDir string) (map[string]string, error) {
	movieIDToGenre := make(map[string]string, 0)

	// For each row sorts the genre list alphabetically because they could be out of order
	var action fileparser.RowAction = func(fields []string) error {
		var genreSlice sort.StringSlice
		// field 2 is genres
		genreSlice = strings.Split(fields[2], GenreDelimiter)
		genreSlice.Sort()
		// field 0 is movie id
		movieIDToGenre[fields[0]] = strings.Join(genreSlice, GenreDelimiter)
		return nil
	}

	if err := fileparser.ParseFile(filepath.Join(dataDir, FileNameMovies), FieldCountMovies, action); err != nil {
		return nil, err
	}

	return movieIDToGenre, nil
}

// GetYearlyGenreCountResults parses ratings data file, for each row it increments genre count by year based on the timestamp of the tweet.
// Results returned are sorted by year in ascending order. Genre counts are not sorted giving client freedom to sort in the ways they want
func GetYearlyGenreCountResults(dataDir string, currentYear int, movieIDToGenre map[string]string) (resultm.Results, error) {

	// Calculate minimum date past decade
	minAcceptedDate := time.Date((currentYear - 10), time.January, 1, 0, 0, 0, 0, time.UTC)
	// Calculate the minimum accepted seconds since epoch, ratings older than this will be ignored
	minAcceptedDateSec := minAcceptedDate.Unix()

	// Calculate maximum date allowed, include current year
	maxAcceptedDate := time.Date((currentYear + 1), time.January, 1, 0, 0, 0, 0, time.UTC)
	maxAcceptedDateSec := maxAcceptedDate.Unix()
	fmt.Printf("Capturing ratings between %s and %s\n\n",
		minAcceptedDate.Format(time.RFC3339), maxAcceptedDate.Format(time.RFC3339))

	// For each rating record increments count per genre
	rm := resultm.NewResultManager()
	var action fileparser.RowAction = func(fields []string) error {
		if secondsSinceEpoch, err := strconv.ParseInt(fields[3], 10, 64); err == nil {
			if secondsSinceEpoch >= minAcceptedDateSec &&
				secondsSinceEpoch < maxAcceptedDateSec {
				// lookup genre by movie id
				if genre, ok := movieIDToGenre[fields[1]]; ok {
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
