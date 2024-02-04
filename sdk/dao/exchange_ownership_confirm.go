package dao

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/constants"
	idl2 "github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/gin-gonic/gin"
)

func (c *Client) ExchangeOwnershipConfirm(ctx *gin.Context, request *idl2.ApiExchangeOwnershipRequest) error {
	url := c.DaoHost + "/dip/data/confirm_exchange_ownership"
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
		return errors.New("ExchangeOwnership confirm return not success," + string(body))
	}
	return nil
}
