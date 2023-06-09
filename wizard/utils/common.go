package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func ChunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// GetNextItem to get next item from slice
func GetNextItem(s []string, current string) (string, error) {
	var nextItem string
	var found bool

	for _, v := range s {
		if found {
			nextItem = v
			break
		}
		if v == current {
			found = true
		}
	}

	if !found {
		return "", fmt.Errorf("element %s not found in slice", current)
	}

	if nextItem == "" {
		// return the first item if current value is the last element in the slice
		return s[0], nil
	}

	return nextItem, nil
}

// get previous item from slice
func GetPrevItem(s []string, current string) (string, error) {
	var prevItem string
	var found bool

	for _, v := range s {
		if v == current {
			found = true
			break
		}
		prevItem = v
	}

	if !found {
		return "", fmt.Errorf("element %s not found in slice", current)
	}

	if prevItem == "" {
		// return the last item if current value is the first element in the slice
		return s[len(s)-1], nil
	}

	return prevItem, nil
}

func GetRandomItem(options []string, valueToIgnore string) string {
	var item string
	rand.Seed(time.Now().UnixNano())
	for {
		item = options[rand.Intn(len(options))]
		if item != valueToIgnore {
			break
		}
	}
	return item
}

func AddNewLines(text string, charLimit int) string {
	var result strings.Builder
	lines := strings.Split(text, "\n")

	for lineIdx, line := range lines {
		words := strings.Fields(line)
		var currentLine strings.Builder

		for _, word := range words {
			if currentLine.Len()+len(word) > charLimit {
				if result.Len() > 0 {
					result.WriteString("\n")
				}
				result.WriteString(currentLine.String())
				currentLine.Reset()
			} else if currentLine.Len() > 0 {
				currentLine.WriteString(" ")
			}

			currentLine.WriteString(word)
		}

		if currentLine.Len() > 0 {
			if result.Len() > 0 && lineIdx > 0 {
				result.WriteString("\n")
			}
			result.WriteString(currentLine.String())
		}

		if lineIdx < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}
