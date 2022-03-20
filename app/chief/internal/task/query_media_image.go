package task

import (
	"fmt"
	"github.com/mars-projects/mars/common/framework"
	"log"
	"time"
)

type QueryMediaImageExecutor struct {
	Sender framework.MessageSender
	//ImageServer *ImageManager
}

func (executor *QueryMediaImageExecutor) Execute(id uint32, request framework.Message,
	incoming chan framework.Message, terminate chan bool) (err error) {
	fmt.Println(request)
	var originSession = request.GetFromSession()
	//var respChan = make(chan modules.ResourceResult, 1)
	//executor.ResourceModule.GetImageServer(respChan)
	//var result = <-respChan
	resp, _ := framework.NewJsonMessage(framework.QueryMediaImageResponse)
	resp.SetSuccess(false)
	resp.SetFromSession(id)
	resp.SetToSession(request.GetFromSession())
	//
	//if result.Error != nil {
	//	err := result.Error
	//	log.Printf("[%08X] get image server fail: %s", id, err.Error())
	//	resp.SetError(err.Error())
	//	return executor.Sender.SendMessage(resp, request.GetSender())
	//}

	//forward to image server
	request.SetFromSession(id)
	request.SetToSession(0)
	//var imageServer = result.Name

	//if err = executor.Sender.SendMessage(request, imageServer); err != nil {
	//	log.Printf("[%08X] forward query media to image server fail: %s", id, err.Error())
	//	resp.SetError(err.Error())
	//	return executor.Sender.SendMessage(resp, request.GetSender())
	//}
	////wait response
	timer := time.NewTimer(time.Second * 5)
	select {
	case forwardResp := <-incoming:
		if !forwardResp.IsSuccess() {
			log.Printf("[%08X] query media image fail: %s", id, forwardResp.GetError())
		}
		fmt.Println(forwardResp.GetID())
		forwardResp.SetFromSession(id)
		forwardResp.SetToSession(originSession)
		forwardResp.SetTransactionID(request.GetTransactionID())
		//forward
		return executor.Sender.SendMessage(forwardResp, request.GetSender())
	case <-timer.C:
		//timeout
		log.Printf("[%08X] query media image timeout", id)
		resp.SetError("time out")
		return executor.Sender.SendMessage(resp, request.GetSender())
	}
	return nil
}
