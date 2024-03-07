package IDL

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/net/idna"
)

type SignatureData struct {
	OperatorDoi    string `json:"operator_doi" form:"operator_doi" binding:"required"`    //操作者的doi，在不同的场景中，可能为DW或DU
	SignatureNonce string `json:"signature_nonce" form:"operator_doi" binding:"required"` //唯一随机数，用于防止网络重放攻击。用户在不同请求间要使用不同的随机值，建议使用通用唯一识别码UUID（Universally Unique Identifier）
	Signature      string `json:"signature" form:"operator_doi" binding:"required"`       //对请求进行秘钥签名
}

func NewSignatureData() *SignatureData {
	return &SignatureData{}
}
func (s *SignatureData) SetOperator(operator string) *SignatureData {
	s.OperatorDoi = operator
	return s
}

func (s *SignatureData) SetNonce(nonce string) *SignatureData {
	s.SignatureNonce = nonce
	return s
}

func (s *SignatureData) SetSign(sign string) *SignatureData {
	s.Signature = sign
	return s
}

func NewSignatureDataWithSign(operator, prvKey string) *SignatureData {
	var err error
	signData := NewSignatureData()
	signData.OperatorDoi = operator
	signData.SignatureNonce = uuid.NewString()
	signData.Signature, err = signData.CreateSignature(prvKey)
	if err != nil {
		return nil
	}
	return signData
}

// CreateSignature 签名：采用sha256算法进行签名并输出为hex格式（私钥PKCS8格式）
func (s *SignatureData) CreateSignature(prvKey string) (string, error) {
	if s.OperatorDoi == "" || s.SignatureNonce == "" {
		//return "", errors.New("invalid signature params")
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
	h := sha256.New()
	originData, err := s.genSignOriginData()
	if err != nil {
		return "", err
	}
	h.Write([]byte(originData))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hash)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	out := hex.EncodeToString(signature)
	s.Signature = out
	return out, nil
}

// VerifySignature 验签：对采用sha256算法进行签名后转base64格式的数据进行验签
func (s *SignatureData) VerifySignature(pubKey string) error {
	sign, err := hex.DecodeString(s.Signature)
	if err != nil {
		return err
	}
	/*
		public, _ := base64.StdEncoding.DecodeString(pubKey)
	*/
	keyBytes, err := hex.DecodeString(pubKey)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pub, err := x509.ParsePKIXPublicKey(keyBytes)
	if err != nil {
		return err
	}
	hash := sha256.New()
	originData, err := s.genSignOriginData()
	if err != nil {
		return err
	}
	hash.Write([]byte(originData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hash.Sum(nil), sign)
}

func (s *SignatureData) genSignOriginData() (string, error) {
	operatorDoi, err := idna.ToASCII(s.OperatorDoi)
	if err != nil {
		fmt.Println("Punycode encoding error:", err)
		return "", err
	}
	return operatorDoi + s.SignatureNonce, nil
}
