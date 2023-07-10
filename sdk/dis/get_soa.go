package dis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/constants"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/security"
	"github.com/go-resty/resty/v2"
)

func (c *Client) getSOA(zone string, sk string) (*idl.SOAData, error) {

	hubAddress := ""

	client := resty.New()

	//get private key
	privateKey, err := security.ImportPrivateKey(sk)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	signature, err := security.SignByPK(privateKey, security.Sha256Hash([]byte(zone)))
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer: "+base64.StdEncoding.EncodeToString(signature)).
		Get(hubAddress + "/dip/dis-r/name_server/soa/" + zone)

	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	var cr IDL.CommonResponse
	err = json.Unmarshal(resp.Body(), &cr)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if cr.Code != IDL.RespCodeType(constants.Success) || cr.Data == nil {
		err := fmt.Errorf("response code is %d", int64(cr.Code))
		c.Logger.Error(err.Error())
		return nil, err
	}

	by, err := json.Marshal(cr.Data)
	if err != nil {
		err = fmt.Errorf("cr marshal failed: %v", err)
		c.Logger.Error(err.Error())
		return nil, err
	}

	sr := idl.SOAResponse{}
	err = json.Unmarshal(by, &sr)
	if err != nil {
		err = fmt.Errorf("cr unmarshal failed: %v", err)
		c.Logger.Error(err.Error())
		return nil, err
	}

	soa, err := parseSOAData(sr.SOARR.Rdata)
	if err != nil {
		err = fmt.Errorf("not SOA data: %v", err.Error())
		c.Logger.Error(err.Error())
		return nil, err
	}
	return &soa, nil

}

// parse str to SOAData
func parseSOAData(strData string) (idl.SOAData, error) {
	var sd idl.SOAData
	bytes := []byte(strData)
	err := json.Unmarshal(bytes, &sd)
	if err != nil {
		return sd, fmt.Errorf("unmarshal SOAData failed: " + err.Error())
	}
	return sd, nil
}
