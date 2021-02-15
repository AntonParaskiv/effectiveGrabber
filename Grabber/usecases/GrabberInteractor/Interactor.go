package GrabberInteractor

import (
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/usecases/ApiRepositoryInterface"
)

type Interactor struct {
	apiRepository ApiRepositoryInterface.Repository
	taskOneChan   chan struct{}
	taskTwoChan   chan struct{}
	rpmLimit      int64
	taskTokenChan chan struct{}
}

func New() (i *Interactor) {
	i = new(Interactor)
	i.taskTokenChan = make(chan struct{}, 10)
	return
}

func (i *Interactor) SetApiRepository(apiRepository ApiRepositoryInterface.Repository) *Interactor {
	i.apiRepository = apiRepository
	return i
}

func (i *Interactor) SetTaskOneChan(taskChan chan struct{}) *Interactor {
	i.taskOneChan = taskChan
	return i
}

func (i *Interactor) SetTaskTwoChan(taskChan chan struct{}) *Interactor {
	i.taskTwoChan = taskChan
	return i
}

func (i *Interactor) SetRpmLimit(rpmLimit int64) *Interactor {
	i.rpmLimit = rpmLimit
	return i
}
