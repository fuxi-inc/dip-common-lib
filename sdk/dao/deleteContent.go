package dao

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/constants"
	"github.com/fuxi-inc/dip-common-lib/sdk/dao/idl"
	idl2 "github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/security"
	"github.com/fuxi-inc/dip-common-lib/utils/testpkg"
	"github.com/gin-gonic/gin"
)

func (c *Client) Delete(ctx *gin.Context, req *idl.DeleteDataContentRequest) error {
	url := c.DaoHost + "/dip/data/content"
	method := "PUT"
	request := idl.UpdateDataContentRequest{
		Doi:     req.Doi,
		DwDoi:   req.DwDoi,
		Content: "",
		Digest: &idl2.DataDigest{
			Algorithm: "SHA256",
			Result:    base64.StdEncoding.EncodeToString(security.Sha256Hash([]byte(""))),
		},
		Confirmation: func() string {
			sign, err := IDL.NewSignatureData().SetOperator("").SetNonce(base64.StdEncoding.EncodeToString(security.Sha256Hash([]byte("")))).CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
			fmt.Println("SignByPK-->:", sign, err)
			return sign
		}(),
		SecretKey:     "",
		SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
	}

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
		return errors.New("register return not success," + string(body))
	}
	return nil
}
