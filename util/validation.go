package util

func IsSortType(s string) bool {
	isAsc := s == "asc" || s == "ASC"
	isDesc := s == "desc" || s == "DESC"

	return isAsc || isDesc
}

func IsBoolFromStr(s string) bool {
	isTrue := s == "true"
	isFalse := s == "false"

	return isTrue || isFalse
}
