package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// DataDelimiter is the character(s) used to separate each field in a data line
	DataDelimiter = "::"
	// GenreDelimiter is the character that separate each genre in a genre list
	GenreDelimiter = "|"
	// FileNameMovies is the name of movies data file
	FileNameMovies = "movies.dat"
	// FileNameRatings is the name of ratings data file
	FileNameRatings = "ratings.dat"
	// FieldCountMovies is the number of fields in the movies data file
	FieldCountMovies = 3
	// FieldCountRatings is the number of fields in the ratings data file
	FieldCountRatings = 4
)

func main() {
	var (
		topN        int
		currentYear int
		dataDir     string
	)
	{
		flag.IntVar(&topN, "topn", 1, "Top N most popular genre to print")
		flag.IntVar(&currentYear, "currentyear", time.Now().UTC().Year(),
			"Year to start counting backward, default to current year")
		flag.StringVar(&dataDir, "datadir", "",
			"Directory path containing the *.dat files to analyse")
	}
	flag.Parse()

	for _, f := range []string{FileNameMovies, FileNameRatings} {
		if _, err := os.Stat(filepath.Join(dataDir, f)); os.IsNotExist(err) {
			log.Fatalf("File %s does not exists\n", f)
		}
	}

	moviesFile, err := os.Open(filepath.Join(dataDir, FileNameMovies))
	if err != nil {
		log.Fatalf("Fails to open file %s: %s\n", FileNameMovies, err)
	}
	defer moviesFile.Close()

	// store a map of movie id to sorted list of genres
	movieID2Genre := make(map[string]string, 0)
	moviesScanner := bufio.NewScanner(moviesFile)
	moviesScanner.Split(bufio.ScanLines)
	for {
		if ok := moviesScanner.Scan(); !ok {
			if err := moviesScanner.Err(); err != nil {
				log.Fatalf("Fails to scan movies: %s\n", err)
			}
			// exit the loop when either error or end of file
			break
		}

		line := moviesScanner.Text()
		fields := strings.Split(line, DataDelimiter)
		if len(fields) != FieldCountMovies {
			log.Fatalf("Missing fields in movies file on line: %s\n", line)
		}

		// Sort the genre list alphabetically because they could be out of order
		var genreSlice sort.StringSlice
		genreSlice = strings.Split(fields[2], GenreDelimiter)
		genreSlice.Sort()
		movieID2Genre[fields[0]] = strings.Join(genreSlice, GenreDelimiter)
	}
	//fmt.Printf("%+v", movieID2Genre)

	rm := NewResultManager()

	minAcceptedDate := time.Date((currentYear - 10), time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("minAcceptedDate: %v\n", minAcceptedDate)
	// Ignore dates less than this
	minAcceptedDateSec := minAcceptedDate.Unix()

	ratingsFile, err := os.Open(filepath.Join(dataDir, FileNameRatings))
	if err != nil {
		log.Fatalf("Fails to open file %s: %s\n", FileNameRatings, err)
	}
	defer ratingsFile.Close()

	ratingsScanner := bufio.NewScanner(ratingsFile)
	ratingsScanner.Split(bufio.ScanLines)
	for {
		if ok := ratingsScanner.Scan(); !ok {
			if err := ratingsScanner.Err(); err != nil {
				log.Fatalf("Fails to scan ratings: %s\n", err)
			}
			// exit the loop when either error or end of file
			break
		}

		line := ratingsScanner.Text()
		fields := strings.Split(line, DataDelimiter)
		if len(fields) != FieldCountRatings {
			log.Fatalf("Missing fields in rating file on line: %s\n", line)
		}

		secondsSinceEpoch, err := strconv.ParseInt(fields[3], 10, 64)
		if err != nil {
			log.Fatalf("Invalid seconds since epoch %s: %s\n", fields[3], err)
		}

		if secondsSinceEpoch >= minAcceptedDateSec {
			// lookup genre by movie id
			if genre, ok := movieID2Genre[fields[1]]; ok {
				rm.Capture(time.Unix(secondsSinceEpoch, 0).Year(), genre)
			} else {
				log.Printf("Warning movie id cannot be found in movies.dat: %s\n", fields[1])
			}
		}
	} // end for

	results := rm.GetResult()
	for _, r := range results {
		fmt.Printf("Year: %d\n", r.Year)

		genreCounts := r.GenreCounts
		sort.Sort(ByMostPopular{genreCounts})

		// Prevent printing out-of-range elements.
		// If number of genre-counts is less than topN, set topN to the len
		tempTopN := topN
		if len(genreCounts) < topN {
			tempTopN = len(genreCounts)
		}
		for _, gc := range genreCounts[:tempTopN] {
			fmt.Printf("-- %s : %d\n", gc.Genre, gc.Count)
		}
	}

	// i, err := strconv.ParseInt(os.Args[1], 10, 64)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// ratingTime := time.Unix(i, 0).UTC()
	// ratingTimeStr := ratingTime.Format(time.RFC3339Nano)
	// fmt.Printf("ratingTime local: %s\n", ratingTimeStr)

	// yr := ratingTime.Year()
	// fmt.Println(yr)
}
