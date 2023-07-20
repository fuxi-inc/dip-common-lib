package dao

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/constants"
	"github.com/fuxi-inc/dip-common-lib/sdk/dao/idl"
	"github.com/gin-gonic/gin"
)

func (c *Client) Delete(ctx *gin.Context, request *idl.DeleteDataContentRequest) error {
	url := c.DaoHost + "/dip/data/delete"
	method := "DELETE"

	payload := strings.NewReader(request.ToString())

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
		return errors.New("delete return not success," + string(body))
	}
	return nil
}
