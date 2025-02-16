package resource

import (
	"github.com/shopspring/decimal"
)

type Type int

const (
	ResourceEnergy Type = 1 << iota
	ResourceWater
	ResourceWood
	ResourceJunk
	ResourceFood
	ResourceMetal
	ResourceFuel
)

func (t Type) With(cnt int) Resource {
	return New(t, cnt)
}

func (t Type) Cost(cnt int) Resource {
	return New(t, -cnt)
}

func (t Type) Produce(cnt int) Resource {
	return New(t, cnt)
}

func (t Type) CanDivByHour() bool {
	if t == ResourceEnergy {
		return false
	}
	return true
}

func AllTypes() []Type {
	return []Type{
		ResourceEnergy,
		ResourceWater,
		ResourceWood,
		ResourceJunk,
		ResourceFood,
		ResourceMetal,
		ResourceFuel,
	}
}

type Resource struct {
	Type  Type
	Count decimal.Decimal
}

type Resources struct {
	AddAble    []Resource
	NotAddAble []Resource
}
