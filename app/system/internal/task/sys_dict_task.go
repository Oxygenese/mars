package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
	"github.com/mars-projects/mars/common/utils"
)

type QueryDictDataSelectExecutor struct {
	*biz.SysDictData
}

//Execute 根据字典类型查询字典数据信息
func (e QueryDictDataSelectExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	req := dto.SysDictDataGetPageReq{}
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]models.SysDictData, 0)
	err = e.GetAll(&req, &list)
	l := make([]dto.SysDictDataGetAllResp, 0)
	for _, i := range list {
		d := dto.SysDictDataGetAllResp{}
		utils.Translate(i, &d)
		l = append(l, d)
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, &l)
	return nil
}
