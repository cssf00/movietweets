package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var (
		topN        int
		currentYear int
	)
	{
		flag.IntVar(&topN, "topn", 1, "Top N most popular genre to print")
		flag.IntVar(&currentYear, "currentyear", time.Now().UTC().Year(),
			"Year to start counting backward, default to current year")
	}
	flag.Parse()

	minAcceptedDate := time.Date((currentYear - 10), time.January, 1, 0, 0, 0, 0, time.UTC)
	// Ignore dates less than this
	minAcceptedDateSec := minAcceptedDate.Unix()
	fmt.Println(minAcceptedDate)
	fmt.Println(minAcceptedDateSec)

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
