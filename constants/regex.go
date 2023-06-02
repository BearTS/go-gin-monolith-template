package constants

var Regex = struct {
	REGEXP_MOBILE string
	REGEX_EMAIL   string
}{
	REGEXP_MOBILE: "^[5-9]{1}[0-9]{9}$",
	REGEX_EMAIL:   "^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$",
}
