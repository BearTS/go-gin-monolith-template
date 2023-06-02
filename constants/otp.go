package constants

var OtpTypes = struct {
	EMAIL string
}{
	EMAIL: "email",
}

var OtpStatuses = struct {
	PENDING  string
	VERIFIED string
	EXPIRED  string
	FAILED   string
}{
	PENDING:  "pending",
	VERIFIED: "verified",
	EXPIRED:  "expired",
	FAILED:   "failed",
}
