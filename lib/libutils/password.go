package libutils

import (
	"crypto/sha1"
	"fmt"
)

func EncryptWord(passwd string) (hashResult string) {
	h := sha1.New()
	h.Write([]byte(passwd))

	//@see https://segmentfault.com/q/1010000007510284
	hashResult = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func CombineSalt(raw string, salt string) string {
	return raw + salt
}
