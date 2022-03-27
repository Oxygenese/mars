package transaction

import (
	"fmt"
	"github.com/mars-projects/mars/api"
	"log"
)

type H map[string]interface{}

type Sender interface {
	Send(message *api.Reply) error
}

func NewTransactionEngine(queueSize int, sender Sender) (*Engine, error) {
	if queueSize <= 0 {
		queueSize = 1 << 10
	}
	engine := Engine{}
	engine.sender = sender
	engine.executorMap = map[api.Operate]BaseExecutor{}
	engine.invokeChan = make(chan *api.Message, queueSize)
	engine.exitChan = make(chan bool, 1)
	engine.resChan = map[string]chan *api.Reply{}
	return &engine, nil
}

//type Executor interface {
//	Execute(message *api.Message, respChan chan *api.Reply, sender Sender) error
//}

type ExecutorFn func(message *api.Message, respChan chan *api.Reply, sender Sender) error

type Engine struct {
	sender      Sender
	executorMap map[api.Operate]BaseExecutor
	resChan     map[string]chan *api.Reply
	invokeChan  chan *api.Message
	exitChan    chan bool
}

func (engine *Engine) IsExecutorExists(operate api.Operate) bool {
	_, exists := engine.executorMap[operate]
	return exists
}

type BaseExecutor struct {
	ExecutorFn
	sync bool
}

func (engine *Engine) IsSync(operation api.Operate) bool {
	return engine.executorMap[operation].sync
}

func (engine *Engine) RegisterExecutor(operation api.Operate, sync bool, executor ExecutorFn) error {
	if engine.IsExecutorExists(operation) {
		return fmt.Errorf("executor already bound with message %08X", operation)
	}
	engine.executorMap[operation] = BaseExecutor{
		ExecutorFn: executor,
		sync:       sync,
	}
	return nil
}

func (engine *Engine) PushMessage(message *api.Message) error {
	// 设置返回值通道
	engine.resChan[message.GetRequestId()] = make(chan *api.Reply)
	engine.invokeChan <- message
	return nil
}

func (engine *Engine) GetResChan(requestId string) chan *api.Reply {
	return engine.resChan[requestId]
}

func (engine *Engine) Start() error {
	go engine.routine()
	return nil
}

func (engine *Engine) Stop() error {
	engine.exitChan <- true
	return nil
}

func (engine *Engine) routine() {
	exitFlag := false
	for !exitFlag {
		select {
		case <-engine.exitChan:
			exitFlag = true
		case msg := <-engine.invokeChan:
			executor, exists := engine.executorMap[msg.GetOperate()]
			if !exists {
				log.Printf("[transaction engin] no executor registered for message [%08X]", msg.GetOperate())
				break
			}
			go executeTask(executor.ExecutorFn, msg, engine.resChan[msg.GetRequestId()], engine.sender)
			break
		}
	}
	engine.exitChan <- true
}

func executeTask(executor ExecutorFn, msg *api.Message, respChan chan *api.Reply, sender Sender) {
	err := executor(msg, respChan, sender)
	if err != nil {
		log.Printf("[transaction engin]  execute task(msg: %08X) fail: %s", msg.GetOperate(), err.Error())
		//err := sender.Send(msg)
		if err != nil {
			log.Printf("[transaction engin] (msg_id: %08X) send message fail: %s", msg.GetRequestId(), err.Error())
			return
		}
	}
}
