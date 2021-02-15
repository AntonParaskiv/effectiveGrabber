package GrabberInteractor

import (
	"fmt"
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/usecases/ApiRepositoryInterface"
	"github.com/AntonParaskiv/effectiveGrabber/Grabber/usecases/mock_ApiRepositoryInterface"
	"github.com/golang/mock/gomock"
	"testing"
)

var SimulatedError = fmt.Errorf("simulated error")

func TestInteractor_doTaskOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiRepositoryMockSuccess := mock_ApiRepositoryInterface.NewMockRepository(ctrl)
	apiRepositoryMockSuccess.
		EXPECT().
		GetOne().Return(nil).
		AnyTimes()

	apiRepositoryMockError := mock_ApiRepositoryInterface.NewMockRepository(ctrl)
	apiRepositoryMockError.
		EXPECT().
		GetOne().Return(SimulatedError).
		AnyTimes()

	type fields struct {
		apiRepository ApiRepositoryInterface.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				apiRepository: apiRepositoryMockSuccess,
			},
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				apiRepository: apiRepositoryMockError,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				apiRepository: tt.fields.apiRepository,
			}
			if err := i.doTaskOne(); (err != nil) != tt.wantErr {
				t.Errorf("doTaskOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInteractor_doTaskTwo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiRepositoryMockSuccess := mock_ApiRepositoryInterface.NewMockRepository(ctrl)
	apiRepositoryMockSuccess.
		EXPECT().
		GetTwo().Return(nil).
		AnyTimes()

	apiRepositoryMockError := mock_ApiRepositoryInterface.NewMockRepository(ctrl)
	apiRepositoryMockError.
		EXPECT().
		GetTwo().Return(SimulatedError).
		AnyTimes()

	type fields struct {
		apiRepository ApiRepositoryInterface.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				apiRepository: apiRepositoryMockSuccess,
			},
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				apiRepository: apiRepositoryMockError,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				apiRepository: tt.fields.apiRepository,
			}
			if err := i.doTaskTwo(); (err != nil) != tt.wantErr {
				t.Errorf("doTaskTwo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInteractor_selectTask(t *testing.T) {
	taskOneChanSelectOne := make(chan struct{}, 1)
	taskOneChanSelectOne <- struct{}{}
	taskLeftChanSelectOne := make(chan struct{}, 2)
	taskLeftChanSelectOne <- struct{}{}
	taskLeftChanSelectOne <- struct{}{}

	taskTwoChanSelectTwo := make(chan struct{}, 1)
	taskTwoChanSelectTwo <- struct{}{}
	taskLeftChanSelectTwo := make(chan struct{}, 2)
	taskLeftChanSelectTwo <- struct{}{}

	taskOneChanSelectRandomOne := make(chan struct{}, 1)
	taskOneChanSelectRandomOne <- struct{}{}
	taskLeftChanSelectRandomOne := make(chan struct{})

	taskTwoChanSelectRandomTwo := make(chan struct{}, 1)
	taskTwoChanSelectRandomTwo <- struct{}{}
	taskLeftChanSelectRandomTwo := make(chan struct{})

	type fields struct {
		taskOneChan  chan struct{}
		taskTwoChan  chan struct{}
		taskLeftChan chan struct{}
	}
	tests := []struct {
		name           string
		fields         fields
		wantTaskNumber int
	}{
		{
			name: "Select One",
			fields: fields{
				taskOneChan:  taskOneChanSelectOne,
				taskLeftChan: taskLeftChanSelectOne,
			},
			wantTaskNumber: TaskNumberOne,
		},
		{
			name: "Select Two",
			fields: fields{
				taskTwoChan:  taskTwoChanSelectTwo,
				taskLeftChan: taskLeftChanSelectTwo,
			},
			wantTaskNumber: TaskNumberTwo,
		},
		{
			name: "Select Random One",
			fields: fields{
				taskOneChan:  taskOneChanSelectRandomOne,
				taskLeftChan: taskLeftChanSelectRandomOne,
			},
			wantTaskNumber: TaskNumberOne,
		},
		{
			name: "Select Random Two",
			fields: fields{
				taskTwoChan:  taskTwoChanSelectRandomTwo,
				taskLeftChan: taskLeftChanSelectRandomTwo,
			},
			wantTaskNumber: TaskNumberTwo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				taskOneChan:   tt.fields.taskOneChan,
				taskTwoChan:   tt.fields.taskTwoChan,
				taskTokenChan: tt.fields.taskLeftChan,
			}
			if gotTaskNumber := i.selectTask(); gotTaskNumber != tt.wantTaskNumber {
				t.Errorf("selectTask() = %v, want %v", gotTaskNumber, tt.wantTaskNumber)
			}
		})
	}
}
