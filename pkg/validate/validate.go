package validate

import (
	"fmt"
	"regexp"
)

func IsValidUUID(uuid string) bool {
	// Regular expression for matching UUID format
	pattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	matched, err := regexp.MatchString(pattern, uuid)
	if err != nil {
		// Handle error if regex compilation fails
		fmt.Println("Error:", err)
		return false
	}
	return matched
}

func IsStrSortType(s string) bool {
	isAsc := s == "asc" || s == "ASC"
	isDesc := s == "desc" || s == "DESC"

	return isAsc || isDesc
}

func IsStrBool(s string) bool {
	isTrue := s == "true"
	isFalse := s == "false"

	return isTrue || isFalse
}
