package private

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"time"
	"strings"
	"strconv"
)

func SignRequest(apiSecret string, timestamp time.Time, verb string, path string, body string) string {
	// Create a new Keyed-Hash Message Authentication Code (HMAC) using SHA512 and API Secret
	mac := hmac.New(sha512.New, []byte(apiSecret))
	// Convert timestamp to nanoseconds then divide by 1000000 to get the milliseconds
	timestampString := strconv.FormatInt(timestamp.UnixNano()/1000000, 10)
	mac.Write([]byte(timestampString))
	mac.Write([]byte(strings.ToUpper(verb)))
	mac.Write([]byte(path))
	mac.Write([]byte(body))
	// Gets the byte hash from HMAC and converts it into a hex string
	return hex.EncodeToString(mac.Sum(nil))
}