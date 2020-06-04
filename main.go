package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	resultm "github.com/samuel-foo/movietweets/resultmanager"
)

func main() {
	var (
		dataDir     string
		currentYear int
		topN        int
	)
	{
		flag.StringVar(&dataDir, "datadir", "",
			"Directory path containing the *.dat files to analyse")
		flag.IntVar(&topN, "topn", 1, "Top N most popular genre to print")
		flag.IntVar(&currentYear, "currentyear", time.Now().UTC().Year(),
			"Year to start counting backward, default to current year")
	}

	// Parse command line arguments
	flag.Parse()

	// Validate files must exist
	for _, f := range []string{FileNameMovies, FileNameRatings} {
		if _, err := os.Stat(filepath.Join(dataDir, f)); os.IsNotExist(err) {
			log.Fatalf("File %s does not exists\n", f)
		}
	}

	movieID2Genre, err := getMovieIDToGenreMap(dataDir)
	if err != nil {
		os.Exit(2)
	}

	results, err := processRatings(dataDir, currentYear, movieID2Genre)
	if err != nil {
		os.Exit(3)
	}

	// Print results
	fmt.Println("---------- Result ----------")
	for _, r := range results {
		fmt.Printf("Year: %d\n", r.Year)

		genreCounts := r.GenreCounts
		sort.Sort(resultm.ByMostPopular{GenreCounts: genreCounts})

		// Prevent printing out-of-range elements.
		// If number of genre-counts is less than topN, set topN to the len
		tempTopN := topN
		if len(genreCounts) < topN {
			tempTopN = len(genreCounts)
		}
		for _, gc := range genreCounts[:tempTopN] {
			fmt.Printf("-- %s : %d\n", gc.Genre, gc.Count)
		}

		fmt.Println() // separate year section if there are multiple
	}
	fmt.Println("----------------------------")
}
