package mockExample

import (
	"testing"
	"github.com/golang/mock/gomock"
)

func TestInvoke(t *testing.T) {

	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	mock := NewMockDBService(mockctrl)

	// Define expected behaviour
	call := mock.EXPECT().Init().Times(1)
	call = mock.EXPECT().Insert("Some item").Times(2).After(call)
	call = mock.EXPECT().Count().MinTimes(1).After(call)
	mock.EXPECT().Get("some item").AnyTimes().After(call)
	mock.EXPECT().DeleteAll().AnyTimes()

	// Invoke test
	mock.Init()
	mock.Insert("Some item")
	mock.Insert("Some item")
	mock.Count()
	mock.DeleteAll()

}