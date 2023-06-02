package constants

var Charset = struct {
	ALPHABETS string
	ALPHANUMS string
	NUMERIC   string
}{
	ALPHABETS: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	ALPHANUMS: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	NUMERIC:   "0123456789",
}
