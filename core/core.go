package core

import "context"

type Inserter interface {
	Insert(ctx context.Context, contents []string) error
}
