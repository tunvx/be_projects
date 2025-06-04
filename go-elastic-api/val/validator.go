package val

import (
	"regexp"
	"time"
)

// IsAfterValid checks if the "after" date is in the correct format and a valid date.
func IsDateValid(after string) bool {
	// Regular expression for YYYY-MM-DD format
	match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, after)
	if !match {
		return false
	}

	// Check if it is a real date
	_, err := time.Parse("2006-01-02", after)
	return err == nil
}
