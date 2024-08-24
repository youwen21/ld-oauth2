package libutils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"runtime"
)

var (
	PublicKeyStr = `-----BEGIN PUBLIC KEY-----
-----END PUBLIC KEY-----`

	PrivateKeyStr = `-----BEGIN RSA PRIVATE KEY-----
-----END RSA PRIVATE KEY-----`
)

// js
//https://github.com/travist/jsencrypt

// rsa 加密经验+1
// 加密对象长度有限制， 突然就更明白https为什么是这样设计了。
// https://cloud.tencent.com/developer/section/1140761

type rsaUtil struct {
}

var RsaUtil = &rsaUtil{}

func (r *rsaUtil) Decrypt(cryptText, privateKey []byte) (plainText []byte, err error) {
	block, _ := pem.Decode(privateKey)

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the privateKey is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return []byte{}, err
	}
	plainText, err = rsa.DecryptPKCS1v15(rand.Reader, key, cryptText)
	if err != nil {
		return []byte{}, err
	}
	return plainText, nil
}

func (r *rsaUtil) Encrypt(plainText, publicKey []byte) (cryptText []byte, err error) {
	block, _ := pem.Decode(publicKey)

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the publicKey is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return []byte{}, err
	}
	cryptText, err = rsa.EncryptPKCS1v15(rand.Reader, key.(*rsa.PublicKey), plainText)
	if err != nil {
		return []byte{}, err
	}
	return cryptText, nil
}
