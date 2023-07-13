package dao

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/constants"
	idl2 "github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) AuthInit(ctx *gin.Context, request *idl2.ApiAuthInitRequest) error {
	url := c.DaoHost + "/dip/data/initiate_auth"
	method := "POST"
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", constants.MIMEApplicationJSON)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	response := &IDL.CommonResponse{}
	err = response.Unmarshal(body)
	if err != nil {
		return err
	}
	if response.Code != 0 {
		return errors.New("AuthInit return not success," + string(body))
	}
	return nil
}
