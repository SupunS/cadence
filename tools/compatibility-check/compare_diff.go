package compatibility_check

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kylelemons/godebug/diff"
)

func CompareFiles(oldResults, newResults *os.File) error {

	oldResultsScanner := bufio.NewScanner(oldResults)
	oldResultsScanner.Split(bufio.ScanLines)
	var oldResultsLines []string
	for oldResultsScanner.Scan() {
		oldResultsLines = append(oldResultsLines, oldResultsScanner.Text())
	}

	newResultsScanner := bufio.NewScanner(newResults)
	newResultsScanner.Split(bufio.ScanLines)
	var newResultsLines []string
	for newResultsScanner.Scan() {
		newResultsLines = append(newResultsLines, newResultsScanner.Text())
	}

	totalDiffs := 0

	const batchSize = 10_000
	const maxDiffs = 1000

	// `diff.DiffChunks` can't seem to handle when the diff is too large.
	// Hence, check the diff in batches.
	// Might not be 100% accurate if there's a very large diff at one side.
	// But it should be good enough for the reporting.
	// (i.e: worst-case: can have false positives, but no false negatives)

batchLoop:
	for batchStart := 0; batchStart < len(oldResultsLines); batchStart += batchSize {
		batchEnd := batchStart + batchSize
		oldResultsBatchEnd := len(oldResultsLines)
		if oldResultsBatchEnd > batchEnd {
			oldResultsBatchEnd = batchEnd
		}

		newResultsBatchEnd := len(newResultsLines)
		if newResultsBatchEnd > batchEnd {
			newResultsBatchEnd = batchEnd
		}

		diffChunks := diff.DiffChunks(
			oldResultsLines[batchStart:oldResultsBatchEnd],
			newResultsLines[batchStart:newResultsBatchEnd],
		)

		var prevChunk diff.Chunk

		for _, currentChunk := range diffChunks {
			if currentChunk.Added == nil && currentChunk.Deleted == nil {
				continue
			}

			printDiffChunk(prevChunk, currentChunk)

			totalDiffs++
			prevChunk = currentChunk

			// stop reporting too many errors
			if totalDiffs > maxDiffs {
				break batchLoop
			}
		}
	}

	if totalDiffs > 0 {
		return errors.New("found differences")
	}

	return nil
}

func printDiffChunk(prevChunk, currentChunk diff.Chunk) {
	sb := strings.Builder{}

	var extraLinesToPrint = 4

	// Print the last few line from the previous chunk.
	equalLines := prevChunk.Equal
	equalLinesLen := len(equalLines)
	if equalLinesLen <= extraLinesToPrint {
		extraLinesToPrint = equalLinesLen
	}

	printFrom := equalLinesLen - extraLinesToPrint
	for i := printFrom; i < equalLinesLen; i++ {
		// Thus, Print the previous lines.
		sb.WriteString(equalLines[i])
		sb.WriteRune('\n')
	}

	// Print additions
	for _, line := range currentChunk.Added {
		sb.WriteString("+")
		sb.WriteString(line)
	}

	// Print deletions
	for _, line := range currentChunk.Deleted {
		sb.WriteString("-")
		sb.WriteString(line)
	}

	// If there are 'Equal' lines, that means this is the last chunk of a diff.
	if len(currentChunk.Equal) > 0 {
		// Keep some space between this diff and the next diff.
		sb.WriteRune('\n')
	}

	fmt.Println(sb.String())
}
