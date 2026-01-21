package orders

import "context"

type Service interface {
	AddOrder (context context.Context)
}
