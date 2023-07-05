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
		Get(hubAddress + "/name_server/allzones/")

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

	zones, ok := cr.Data.(idl.ServiceZonesInDIS)
	if !ok {
		err := fmt.Errorf("error: interface claim to ServiceZonesInDIS failed")
		c.Logger.Error(err.Error())
		return nil, err
	}

	return zones.Zones, nil
}
