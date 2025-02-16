package resource

import (
	"github.com/shopspring/decimal"
)

func New(t Type, count int) Resource {
	return Resource{
		Type:  t,
		Count: decimal.NewFromInt(int64(count)),
	}
}
