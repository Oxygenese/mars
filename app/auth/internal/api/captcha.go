package api

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/mars-projects/mars/lib/utils/capthca"
)

func (handler *CaptchaApi) GenerateCaptchaHandler(ctx http.Context) error {
	id, b64s, err := capthca.DriverDigitFunc()
	if err != nil {
		handler.log.Errorf("DriverDigitFunc error, %s", err.Error())
		return err
	}
	var res = map[string]interface{}{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	}
	return ctx.JSON(200, res)
}
