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

func (c *Client) getZone(zone string) (*idl.ZoneResponse, error) {

	hubAddress := "39.107.180.231:8991"
	sk := "MIIEpQIBAAKCAQEA6RJNkDeYkK9D2xbwlxCH9e1yGf6eKQAKZ/fpWkS0RJdW0G1lSDdBLowrU5nlD4EB8WhXGI9A7lYyD1gUkKDtJjdKzK/oGrpKlsJRcEkEPRhMvODKTgDP7Y3jv37JGekHFzmXVPUBBrj7UVqjIUWlHm5TZW6t58E4HQOO6Pm3IRRgkWfMD0mwEY7fDdpGHaZOa9coMJNGOb1JwhCkD0tap+oK+e7ktP24oY6I89XVCVf8dsBwkVM+J+2CA1hJ7BHIn1Ew8AF6Z1xNc2k612QGZZ48irSLRAm3BuJ/0OlSZf5kFBu6r8z5XizMpCpQBwta7Y+1Jcd6ZA8wSeXNiGcjawIDAQABAoIBAQCa+IiKM+1FhsPfkUk3qM4+Pzt2/U28hUJmDvkiBDy4vsgCnjqR4et9P8YvOEyOGoFoOrBurBHIZ5exjCdgYyBMIMejgfMaknO1+k6cEnC3Dsk3bWAWrYOTuJyOX4bnq4IDd0+gNagRCD5SwalmHiDEMydQcl8/kcH1yL/lOOUDIHyVc5TYaCwbNYTRGHh/5UKves3FfnefgaGHwO2ZFJ1TpYir1AutEcZe7hk53+sfPYoeljkWOh7pRH/NKx3h3rLlLeSjA/yUWTBr/Ra3FfQ88QsYpxQaCh48JTyPsaWJdgGIuuXpw7/iZ8AT6Fso/7dyVae7z2cxcBsatKdk1BoBAoGBAPB7Q8ZpP6pzFsdRM+HD6IRKDLmJ0F1emda3gdCipqZm0AI1fjlb+0OPQ9x23IMwWfIrTn/ULO298Cx7rtETVIc1sgVIddyzCD5TRCPGdhrYuv8oHzDZMfR6kDcw8eu09zoFoNYZem+tBN+HqAaFbyoMm8Yi7Wk/hG06zWLvsHrLAoGBAPgcoCGW1Ea9JkS48/oVy394hHfEcH3OQOlH7F4+5qbzxcmwA97+CShpureIb02fal0I2K5BzqHLlcpbp4b2voshVWtg/0gGQnQkP5TPJri6ZeUcuyivwLqpoMRLqXJ4R7jwYASQ0Fe+XVlHsqdmqcVVnqqetBMcmnrOVURUx8XhAoGBAJvX1hEzzc2BLgbXKd8qHXn1nw9yLzUDBIEwhREXBxaZVD3KOtWjNU8P/fNWmArwf6m/xOx0LPWY/JdnQ+A8PbAt7QMddTwkTbhOL79bcthtEaofby8x4fzEnMcXkf+lU/4m/ZciBcVvg9P0TsCDIopGh4C8y8xEHGbJD6cJG/5HAoGBAMwPVxBpKcvIU8ofx+G+vTDo4SJT+SUvVqnG0rusxi03v0ujLXvguY14b+31E33nCCbeWL/xNc5ST68v9LgHBi4Ny2BWVX9aTpuRGI3+Vd2GyLlICRJsgnnDe4sWo5WXIX4UjRQUOpZ/5ezBEEmoISdAxQ+VGYvCcbdhLjXV+LxhAoGAQn5G+TbMt+ADrJoFnQ6lLBaZEh0DKfNtyzcuTlU8IFNU1VgK1gyU5o0TeVybvlzMvGLCIeHuK6yG9NhCE7+skH5zCO/X11MjuMDgqN+VnldgLBlEtxVcZEGiPVqSxrwXDocxDtfJOvU2uFwhR1C0mU+quwf7w5Hqe5DeHftO5Yw="

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
