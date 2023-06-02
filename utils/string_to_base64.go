package utils

import b64 "encoding/base64"

/* -------------------------------------------------------------------------- */
/*                  converts any tring to equivalent base 64                  */
/* -------------------------------------------------------------------------- */
func Str2Base64(str string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(str))
	return sEnc
}
