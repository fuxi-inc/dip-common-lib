package dao

import (
	"fmt"
	"github.com/fuxi-inc/dip-common-lib/constants"
	"github.com/fuxi-inc/dip-common-lib/sdk/dao/idl"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//GetData 调用DAO service 获取服务
func (c *Client) GetData(ctx *gin.Context, request *idl.GetDataRequest) (*idl.GetDataResponse, error) {
	daoUrl := c.DaoHost + "/dip/data/get?" + fmt.Sprintf("du_doi=%s&data_doi=%s&operator_doi=%s&signature_nonce=%s&signature=%s", request.DuDoi, request.DataDoi, request.OperatorDoi, request.SignatureNonce, request.Signature)
	method := constants.GET

	client := &http.Client{}
	req, err := http.NewRequest(method, daoUrl, nil)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	//req.Header.Add(constants.HeaderAuthorization, "<Authorization>")
	req.Header.Add(constants.HeaderContentType, constants.MIMEApplicationJSON)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}

	response := &idl.GetDataResponse{}
	err = response.Unmarshal(body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Code != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.ToString()))
		return nil, err
	}
	return response, nil
}
