package framework

import (
	"fmt"
	"log"
)

type MessageSender interface {
	SendMessage(msg Message, target string) error
	SendToSelf(msg Message) error
}

func NewTransactionEngine(queueSize int) (*TransactionEngine, error) {
	if queueSize <= 0 {
		queueSize = 1 << 10
	}
	engine := TransactionEngine{}
	engine.executorMap = map[uint32]TransactionExecutor{}
	engine.sessions = map[uint32]sessionChannel{}
	var id uint32
	for id = minSessionID; id < minSessionID+sessionCount; id++ {
		engine.sessions[id] = sessionChannel{Allocated: false}
	}

	engine.invokeChan = make(chan Message, queueSize)
	engine.pushChan = make(chan Message, queueSize)
	engine.finishChan = make(chan uint32, queueSize)

	engine.notifyChan = make(chan bool, 1)
	engine.exitChan = make(chan bool, 1)
	return &engine, nil
}

type TransactionExecutor interface {
	Execute(sessionId uint32, message Message, incoming chan Message, terminate chan bool) error
}

type sessionChannel struct {
	Allocated     bool
	pushChan      chan Message
	terminateChan chan bool
}

type TransactionEngine struct {
	executorMap map[uint32]TransactionExecutor
	sessions    map[uint32]sessionChannel
	invokeChan  chan Message
	pushChan    chan Message
	finishChan  chan uint32
	notifyChan  chan bool
	exitChan    chan bool
}

const (
	minSessionID = 1
	sessionCount = 2 << 9
)

func (engine *TransactionEngine) RegisterExecutor(initialMessage uint32, executor TransactionExecutor) error {
	_, exists := engine.executorMap[initialMessage]
	if exists {
		return fmt.Errorf("executor already bound with message %08X", initialMessage)
	}
	engine.executorMap[initialMessage] = executor
	return nil
}

func (engine *TransactionEngine) InvokeTask(message Message) error {
	_, exists := engine.executorMap[message.GetID()]
	if !exists {
		return fmt.Errorf("no executor bound with message %08X", message.GetID())
	}
	if 0 != message.GetToSession() {
		return fmt.Errorf("message [%08X] from %s require specified session [%08X]", message.GetID(), message.GetSender(), message.GetToSession())
	}
	engine.invokeChan <- message
	return nil
}

func (engine *TransactionEngine) PushMessage(message Message) error {
	id := message.GetToSession()
	if session, exists := engine.sessions[id]; exists {
		if session.Allocated {
			//pre check
			engine.pushChan <- message
			return nil
		} else {
			return fmt.Errorf("session [%08X] is not allocated", id)
		}
	}
	return fmt.Errorf("invalid session [%08X]", id)
}

func (engine *TransactionEngine) Start() error {
	go engine.routine()
	return nil
}

func (engine *TransactionEngine) Stop() error {
	engine.notifyChan <- true
	<-engine.exitChan
	return nil
}

func (engine *TransactionEngine) routine() {
	const (
		sessionQueueLength = 10
	)
	var lastID uint32 = minSessionID + sessionCount

	exitFlag := false
	for !exitFlag {
		select {
		case <-engine.notifyChan:
			exitFlag = true
		case msg := <-engine.invokeChan:
			executor, exists := engine.executorMap[msg.GetID()]
			if !exists {
				log.Printf("<trans> no executor registered for message [%08X]", msg.GetID())
				break
			}
			//allocate session
			seed := lastID
			var try uint32
			var invoked = false
			for try = 0; try < sessionCount; try++ {
				id := (seed+try)%sessionCount + minSessionID
				session, exists := engine.sessions[id]
				if !exists {
					log.Printf("<trans> warning: unexpect session [%08X]", id)
					break
				}
				if session.Allocated {
					continue
				}
				//unallocated
				lastID = id
				var pushChan = make(chan Message, sessionQueueLength)
				var tChan = make(chan bool, 1)
				invoked = true
				engine.sessions[id] = sessionChannel{true, pushChan, tChan}
				log.Printf("<trans> [%08X] session allocated", id)
				go executeTask(executor, id, msg, pushChan, tChan, engine.finishChan)
				break
			}
			if !invoked {
				log.Println("<trans> warning: no session available")
			}

		case msg := <-engine.pushChan:
			id := msg.GetToSession()
			if session, exists := engine.sessions[id]; exists {
				if session.Allocated {
					session.pushChan <- msg
					break
				} else {
					log.Printf("<trans> warning: push message %08X to deallocated session [%08X]", msg.GetID(), id)
				}
			} else {
				log.Printf("<trans> warning: message %08X push to invalid session[%08X]", msg.GetID(), id)
			}

		case id := <-engine.finishChan:
			//deallocate session
			if session, exists := engine.sessions[id]; exists {
				if session.Allocated {
					engine.sessions[id] = sessionChannel{Allocated: false}
					//log.Printf("<trans> [%08X] session deallocated", id)
				} else {
					log.Printf("<trans> warning: session [%08X] already deallocated", id)
				}
			} else {
				log.Printf("<trans> warning: try deallocate invalid session[%08X]", id)
			}
		}
	}
	engine.exitChan <- true
}

func executeTask(executor TransactionExecutor, id uint32, msg Message, pushChan chan Message,
	terminateChan chan bool, finished chan uint32) {
	if err := executor.Execute(id, msg, pushChan, terminateChan); err != nil {
		log.Printf("<trans> [%08X] execute task(msg: %08X) fail: %s", id, msg.GetID(), err.Error())
		//}else{
		//	log.Printf("<trans> [%08X] execute finished", id)
	}
	finished <- id
}
