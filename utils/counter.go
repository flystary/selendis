package utils

import "fmt"


func Counter(metrics string, tags map[string]string) string {
	if tags == nil || len(tags) == 0 {
		return metrics
	}
	return fmt.Sprintf("%s/%s", metrics, SortedTags(tags))
}