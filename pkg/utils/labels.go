package utils

import (
	"fmt"
	"strings"
)

// SplitLabelString splits a string in the format `key=value` to map[string]string
func SplitLabelString(labels string) map[string]string {
	mapLabels := make(map[string]string)

	if labels == "" {
		return mapLabels
	}

	// Split a list of labels to the single label pairs
	slicedLabels := strings.Split(labels, ",")
	for _, slicedLabel := range slicedLabels {

		// Split a label into it's key and value
		splitLabel := strings.Split(slicedLabel, "=")
		if len(splitLabel) != 2 {
			// incomplete label, don't know how to handle this...
			continue
		}

		mapLabels[splitLabel[0]] = splitLabel[1]
	}

	return mapLabels
}

// MergeLabels merges a map[string]string to a comma separated string in the format `key=value`
func MergeLabels(labelsMap map[string]string) string {
	labels := make([]string, 0)
	for key, value := range labelsMap {
		labels = append(labels, fmt.Sprintf("%s=%s", key, value))
	}

	if len(labels) == 0 {
		labels = append(labels, "[None]")
	}

	return strings.Join(labels, ",")
}
