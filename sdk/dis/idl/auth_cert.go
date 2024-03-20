package idl

type AuthCert struct {
	TransactionInfo   //data_doi的基础交易信息
	DataAuthorization //data_doi针对du_doi（权属所有者）的授权信息
}
