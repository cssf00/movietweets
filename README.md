# Movie Tweets Analyzer
This command line tool analyzes data files containing tweeter movie ratings, see [here](https://github.com/momenton/momenton-code-test-movietweetings).

It prints out the most popular movie genres, year by year, for the past ten years.

## Building executable
The entry point of this tool is in "movietweets/cmd" directory, where the main.go is.

To build the executable, change directory to "movietweets/cmd" and run the following:
```bash
~/go/src/github.com/samuel-foo/movietweets/cmd$ go build -o movietweets main.go
~/go/src/github.com/samuel-foo/movietweets/cmd$ ls
main.go  movietweets
```

## Command line options
Get help on the tool, run:
```bash
~/go/src/github.com/samuel-foo/movietweets/cmd$ go run main.go -h
```

| Option | Default       |   Description |
| ------------- | ------------- | ------------- |
| -h | n/a |  show help |
| -datadir | current directory | Directory path containing the *.dat files to analyze. Note the directory must contain files with name: movies.dat, ratings.dat  |
| -topn | 3 | Top N most popular genre to print, default to top 3 |
| -currentyear | current year | Year to start counting backward, default to current year |

Show top three most popular genre per year for the last ten years
```bash
~/go/src/github.com/samuel-foo/movietweets/cmd$ go run main.go -datadir "../test_data/big" -topn 3
Capturing ratings between 2010-01-01 00:00:00 +0000 UTC and 2021-01-01 00:00:00 +0000 UTC

---------- Result ----------
Year: 2013
   Drama : 5096
   Comedy : 4734
   Action|Crime|Thriller : 3787

----------------------------
~/go/src/github.com/samuel-foo/movietweets/cmd$
```

Build executable, show top two most popular genre per year for the last ten years
```bash
~/go/src/github.com/samuel-foo/movietweets/cmd$ go build -o movietweets main.go
~/go/src/github.com/samuel-foo/movietweets/cmd$ ls
main.go  movietweets

~/go/src/github.com/samuel-foo/movietweets/cmd$ ./movietweets -datadir "../test_data/small" -topn 2
Capturing ratings between 2010-01-01 00:00:00 +0000 UTC and 2021-01-01 00:00:00 +0000 UTC

---------- Result ----------
Year: 2013
   Comedy|Short : 8
   Comedy|Drama|Romance|Short : 1

Year: 2019
   Comedy|Short : 1
   Crime|Drama : 1

----------------------------
```

Show top two most popular genre per year for the last ten years from year 2029
```bash
~/go/src/github.com/samuel-foo/movietweets/cmd$ ./movietweets -datadir "../test_data/small" -topn 2 -currentyear 2029
Capturing ratings between 2019-01-01 00:00:00 +0000 UTC and 2030-01-01 00:00:00 +0000 UTC

---------- Result ----------
Year: 2019
   Comedy|Short : 1
   Crime|Drama : 1

----------------------------
```