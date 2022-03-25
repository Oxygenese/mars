package task

import (
	"fmt"
	"github.com/mars-projects/mars/common/transaction"
)

func NewCreateImageExecutor() *CreateImageExecutor {
	return &CreateImageExecutor{}
}

type CreateImageExecutor struct {
}

type Image struct {
	Path string
	Cap  string
}

func (executor *CreateImageExecutor) Execute(request transaction.Message, respChan chan transaction.Message) (err error) {
	var image Image

	err = request.GetData(&image)
	fmt.Println(image)
	if err != nil {
		return err
	}
	return nil
}
