package command

import "context"

type Command interface {
	Handle(ctx context.Context) error
}
