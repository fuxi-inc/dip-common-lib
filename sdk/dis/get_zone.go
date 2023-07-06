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

func (c *Client) getZone(zone string, sk string) (*idl.ZoneResponse, error) {

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
		Get(hubAddress + "/dip/dis-r/name_server/zone/" + zone)

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
	zr := idl.ZoneResponse{}
	err = json.Unmarshal(by, &zr)
	if err != nil {
		err = fmt.Errorf("cr unmarshal failed: %v", err)
		c.Logger.Error(err.Error())
		return nil, err
	}

	if len(zr.ZoneDatas) == 0 {
		err = fmt.Errorf("zone response of %s is empty", zone)
		c.Logger.Error(err.Error())
		return nil, err
	}
	return &zr, nil
}
