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

func (c *Client) getAllZones(sk string) ([]string, error) {

	hubAddress := ""

	client := resty.New()

	//get private key
	privateKey, err := security.ImportPrivateKey(sk)
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	//sign
	signText := "sign-text:allzones"
	signature, err := security.SignByPK(privateKey, security.Sha256Hash([]byte(signText)))
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer: "+base64.StdEncoding.EncodeToString(signature)).
		Get(hubAddress + "/dip/dis-r/name_server/allzones/")

	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	var cr *IDL.CommonResponse
	err = json.Unmarshal(resp.Body(), cr)
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
	zones := idl.ServiceZonesInDIS{}
	err = json.Unmarshal(by, &zones)
	if err != nil {
		err = fmt.Errorf("cr unmarshal failed: %v", err)
		c.Logger.Error(err.Error())
		return nil, fmt.Errorf("cr unmarshal failed: %v", err)
	}

	return zones.Zones, nil
}
