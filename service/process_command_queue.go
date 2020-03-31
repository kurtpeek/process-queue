package service

import (
	"context"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const commandsQueue = "queuedCommands:"

var pool = redis.Pool{
	MaxIdle:   50,
	MaxActive: 1000,
	Dial: func() (redis.Conn, error) {
		conn, err := redis.Dial("tcp", ":6379")
		if err != nil {
			logrus.WithError(err).Fatal("initialize Redis pool")
		}
		return conn, err
	},
}

// CommandExecutor executes a command
type CommandExecutor interface {
	Execute(string) error
}

func processQueue(ctx context.Context, done chan<- struct{}, executor CommandExecutor) error {
	rc := pool.Get()
	defer rc.Close()

	for {
		select {
		case <-ctx.Done():
			done <- struct{}{}
			return nil
		default:
			// If the commands queue does not exist, BLPOP blocks until another client
			// performs an LPUSH or RPUSH against it. The timeout argument of zero is
			// used to block indefinitely.
			reply, err := redis.Strings(rc.Do("BLPOP", commandsQueue, 0))
			if err != nil {
				logrus.WithError(err).Errorf("BLPOP %s %d", commandsQueue, 0)
				return errors.Wrapf(err, "BLPOP %s %d", commandsQueue, 0)
			}

			if len(reply) < 2 {
				logrus.Errorf("Expected a reply of length 2, got one of length %d", len(reply))
				return errors.Errorf("Expected a reply of length 2, got one of length %d", len(reply))
			}

			// BLPOP returns a two-element multi-bulk with the first element being the
			// name of the key where an element was popped and the second element
			// being the value of the popped element (cf. https://redis.io/commands/blpop#return-value)
			if err := executor.Execute(reply[1]); err != nil {
				return errors.Wrapf(err, "execute scheduled command: %s", reply[0])
			}
			done <- struct{}{}
		}
	}
}
