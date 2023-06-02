package constants

var Errors = struct {
	VerifyOTP struct {
		InvalidOTP string
		ResendOTP  string
	}
}{
	VerifyOTP: struct {
		InvalidOTP string
		ResendOTP  string
	}{
		InvalidOTP: "invalid OTP",
		ResendOTP:  "please resend OTP",
	},
}
