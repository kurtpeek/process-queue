// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"sync"
)

var (
	lockCommandExecutorMockExecute sync.RWMutex
)

// Ensure, that CommandExecutorMock does implement CommandExecutor.
// If this is not the case, regenerate this file with moq.
var _ CommandExecutor = &CommandExecutorMock{}

// CommandExecutorMock is a mock implementation of CommandExecutor.
//
//     func TestSomethingThatUsesCommandExecutor(t *testing.T) {
//
//         // make and configure a mocked CommandExecutor
//         mockedCommandExecutor := &CommandExecutorMock{
//             ExecuteFunc: func(in1 string) error {
// 	               panic("mock out the Execute method")
//             },
//         }
//
//         // use mockedCommandExecutor in code that requires CommandExecutor
//         // and then make assertions.
//
//     }
type CommandExecutorMock struct {
	// ExecuteFunc mocks the Execute method.
	ExecuteFunc func(in1 string) error

	// calls tracks calls to the methods.
	calls struct {
		// Execute holds details about calls to the Execute method.
		Execute []struct {
			// In1 is the in1 argument value.
			In1 string
		}
	}
}

// Execute calls ExecuteFunc.
func (mock *CommandExecutorMock) Execute(in1 string) error {
	if mock.ExecuteFunc == nil {
		panic("CommandExecutorMock.ExecuteFunc: method is nil but CommandExecutor.Execute was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	lockCommandExecutorMockExecute.Lock()
	mock.calls.Execute = append(mock.calls.Execute, callInfo)
	lockCommandExecutorMockExecute.Unlock()
	return mock.ExecuteFunc(in1)
}

// ExecuteCalls gets all the calls that were made to Execute.
// Check the length with:
//     len(mockedCommandExecutor.ExecuteCalls())
func (mock *CommandExecutorMock) ExecuteCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	lockCommandExecutorMockExecute.RLock()
	calls = mock.calls.Execute
	lockCommandExecutorMockExecute.RUnlock()
	return calls
}
