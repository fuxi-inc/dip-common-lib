package IDL

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

type SignatureData struct {
	OperatorDoi    string `json:"operator_doi" binding:"required"`    //操作者的doi，在不同的场景中，可能为DW或DU
	SignatureNonce string `json:"signature_nonce" binding:"required"` //唯一随机数，用于防止网络重放攻击。用户在不同请求间要使用不同的随机值，建议使用通用唯一识别码UUID（Universally Unique Identifier）
	Signature      string `json:"signature" binding:"required"`       //对请求进行秘钥签名
}

// CreateSignature 签名：采用sha1算法进行签名并输出为hex格式（私钥PKCS8格式）
func (s *SignatureData) CreateSignature(prvKey string) (string, error) {
	if s.OperatorDoi == "" || s.SignatureNonce == "" {
		return "", errors.New("invalid signature params")
	}
	keyBytes, err := hex.DecodeString(prvKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		return "", err
	}
	h := sha1.New()
	h.Write([]byte(s.genSignOriginData()))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	out := hex.EncodeToString(signature)
	return out, nil
}

// VerifySignature 验签：对采用sha1算法进行签名后转base64格式的数据进行验签
func (s *SignatureData) VerifySignature(pubKey string) error {
	sign, err := base64.StdEncoding.DecodeString(s.Signature)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(s.genSignOriginData()))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), sign)
}

func (s *SignatureData) genSignOriginData() string {
	return s.OperatorDoi + s.SignatureNonce
}
