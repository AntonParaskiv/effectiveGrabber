package GrabberInteractor

import (
	"time"
)

func (i *Interactor) taskTokenChanFiller() {
	// calc task token interval
	taskTokenInterval := time.Minute / time.Duration(i.rpmLimit)

	// infinite send task token
	ticker := time.NewTicker(taskTokenInterval)
	for {
		<-ticker.C

		// fill chan if not already full
		if len(i.taskTokenChan) < cap(i.taskTokenChan) {
			i.taskTokenChan <- struct{}{}
		}
	}
}
