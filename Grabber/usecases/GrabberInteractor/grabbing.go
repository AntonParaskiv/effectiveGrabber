package GrabberInteractor

import (
	"fmt"
	"log"
	"time"
)

const (
	_ = iota
	TaskNumberOne
	TaskNumberTwo
)

func (i *Interactor) StartGrabbing() {
	go i.taskTokenChanFiller()
	i.grabbing()
}

func (i *Interactor) grabbing() {
	for {
		// wait task token
		i.waitNextTaskToken()

		// select and do task
		var err error
		switch i.selectTask() {
		case TaskNumberOne:
			err = i.doTaskOne()
		case TaskNumberTwo:
			err = i.doTaskTwo()
		}
		if err != nil {
			log.Println(err)
		}
	}
}

func (i *Interactor) waitNextTaskToken() {
	startWaitingTime := time.Now()
	<-i.taskTokenChan
	log.Println("token waited:", time.Since(startWaitingTime))
}

func (i *Interactor) selectTask() (taskNumber int) {
	// get task one if token num > 0
	if len(i.taskOneChan) > 0 && len(i.taskTokenChan) > 0 {
		<-i.taskOneChan
		taskNumber = TaskNumberOne
		return
	}

	// get task two
	if len(i.taskTwoChan) > 0 {
		<-i.taskTwoChan
		taskNumber = TaskNumberTwo
		return
	}

	// get random task
	select {
	case <-i.taskOneChan:
		taskNumber = TaskNumberOne
	case <-i.taskTwoChan:
		taskNumber = TaskNumberTwo
	}
	return
}

func (i *Interactor) doTaskOne() (err error) {
	log.Println("grab One")
	err = i.apiRepository.GetOne()
	if err != nil {
		err = fmt.Errorf("grabbing get one failed: %w", err)
		return
	}
	return
}

func (i *Interactor) doTaskTwo() (err error) {
	log.Println("grab Two")
	err = i.apiRepository.GetTwo()
	if err != nil {
		err = fmt.Errorf("grabbing get two failed: %w", err)
		return
	}
	return
}
