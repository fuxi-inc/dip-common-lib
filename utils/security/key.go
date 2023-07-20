package security

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func ImportPrivateKey(pkstr string) (*rsa.PrivateKey, error) {

	privateKeyAsBytes, err := base64.StdEncoding.DecodeString(pkstr)
	if err != nil {
		return nil, err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyAsBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func ImportPublicKey(pubKey string) (*rsa.PublicKey, error) {

	publicKeyAsBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return nil, err
	}

	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyAsBytes)
	if err != nil {
		return nil, err
	}

	return publicKey, err
}
func VerifySignature(publicKey *rsa.PublicKey, hashMsg []byte, signature []byte) error {

	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashMsg, signature)
	if err != nil {
		return fmt.Errorf("无法验证签名数据 %v", err)
	}

	return nil
}

// sign by privateKey
func SignByPK(privateKey *rsa.PrivateKey, hashMsg []byte) ([]byte, error) {

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashMsg)
	if err != nil {
		return nil, fmt.Errorf("无法签名消息 %v", err)
	}

	return signature, nil
}

// sha256 hash
func Sha256Hash(msg []byte) []byte {

	hash := sha256.New()
	_, err := hash.Write(msg)
	if err != nil {
		panic(err)
	}

	return hash.Sum(nil)
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return strings.ToUpper(hex.EncodeToString(m.Sum(nil)))
}

func HmacSHA1(data, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
}
