# Movie Tweets Analyzer
This command line tool analyzes data files containing tweeter movie ratings, see [here](https://github.com/momenton/momenton-code-test-movietweetings).

It prints out the most popular movie genres, year by year, for the past decade.

## Command line options
| Option | Default       |   Description |
| ------------- | ------------- | ------------- |
| -datadir | current directory | Directory path containing the *.dat files to analyse. Note the directory must contain files with name: movies.dat, ratings.dat  |
| -topn | 1 | Top N most popular genre to print |
| -currentYear | current year | Year to start counting backward, default to current year |

```bash
sfoo@SamLenoX270:~/go/src/github.com/samuel-foo/movietweets$ go run . -datadir "test_data/small" -topn 3
Minimum accepted date: 2010-01-01 00:00:00 +0000 UTC
---------- Result ----------
Year: 2013
-- Comedy|Short : 8
-- Comedy|Drama|Romance|Short : 1
-- Drama : 1

Year: 2019
-- Comedy|Short : 1
-- Crime|Drama : 1
-- Drama|History|Romance|War : 1

----------------------------
sfoo@SamLenoX270:~/go/src/github.com/samuel-foo/movietweets$
```