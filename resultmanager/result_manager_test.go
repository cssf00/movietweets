package resultmanager

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCaptureAndResults(t *testing.T) {
	g := NewGomegaWithT(t)

	// Arrange
	rm := NewResultManager()
	rm.Capture(2012, "Horror")
	rm.Capture(2012, "Horror")
	rm.Capture(2013, "Thriller")

	// Act
	results := rm.GetResult()

	// Assert
	g.Expect(len(results)).Should(Equal(2))

	g.Expect(results[0].Year).Should(Equal(2012))
	g.Expect(results[0].GenreCounts).Should(Equal(
		GenreCounts{
			GenreCount{Genre: "Horror", Count: 2},
		},
	))

	g.Expect(results[1].Year).Should(Equal(2013))
	g.Expect(results[1].GenreCounts).Should(Equal(
		GenreCounts{
			GenreCount{Genre: "Thriller", Count: 1},
		},
	))
}
