package util

import "fmt"

func Days(days int) string {
	if days == 1 || days == -1 {
		return fmt.Sprintf("%v day", days)
	} else {
		return fmt.Sprintf("%v days", days)
	}
}
