package api

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/mars-projects/mars/common/utils/capthca"
)

type Captcha struct {
	Image string `json:"image,omitempty"`
	Id    string `json:"id,omitempty"`
}

func (handler *CaptchaApi) GenerateCaptchaHandler(ctx http.Context) error {
	id, b64s, err := capthca.DriverDigitFunc()
	if err != nil {
		handler.log.Errorf("DriverDigitFunc error, %s", err.Error())
		return err
	}
	data := Captcha{
		Image: b64s,
		Id:    id,
	}
	marshal, _ := json.Marshal(&data)
	var res = map[string]interface{}{
		"code":    200,
		"data":    string(marshal),
		"id":      id,
		"message": "captcha",
	}
	return ctx.JSON(200, res)
}
