package utils

import (
	"fmt"
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
