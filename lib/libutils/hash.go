package libutils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func Sha1(passwd string) (hashResult string) {
	h := sha1.New()

	h.Write([]byte(passwd))

	// 16进制输出的结果才和php是一样的。  php默认按16进制进行输出。
	//@see https://segmentfault.com/q/1010000007510284
	hashResult = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func Md5(content string) (hashResult string) {
	h := md5.New()
	h.Write([]byte(content))

	// 16进制输出的结果才和php是一样的。  php默认按16进制进行输出。
	//@see https://segmentfault.com/q/1010000007510284
	hashResult = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func Md5Byte(content []byte) (hashResult string) {
	h := md5.New()
	h.Write(content)

	// 16进制输出的结果才和php是一样的。  php默认按16进制进行输出。
	//@see https://segmentfault.com/q/1010000007510284
	return hex.EncodeToString(h.Sum(nil))
}
