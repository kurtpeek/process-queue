package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessQueue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	executor := &CommandExecutorMock{
		ExecuteFunc: func(string) error {
			return nil
		},
	}

	done := make(chan struct{})
	go processQueue(ctx, done, executor)

	rc := pool.Get()
	defer rc.Close()

	_, err := rc.Do("RPUSH", commandsQueue, "foobar")
	require.NoError(t, err)

	<-done

	assert.Exactly(t, 1, len(executor.ExecuteCalls()))
	assert.Exactly(t, "foobar", executor.ExecuteCalls()[0].In1)
}

func TestProcessQueue2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	executor := &CommandExecutorMock{
		ExecuteFunc: func(string) error {
			return nil
		},
	}

	done := make(chan struct{})
	go processQueue(ctx, done, executor)

	rc := pool.Get()
	defer rc.Close()

	_, err := rc.Do("RPUSH", commandsQueue, "foobar")
	require.NoError(t, err)

	<-done

	assert.Exactly(t, 1, len(executor.ExecuteCalls()))
	assert.Exactly(t, "foobar", executor.ExecuteCalls()[0].In1)
}
