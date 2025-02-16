package citizen

import (
	"log/slog"

	"fyne.io/fyne/v2/data/binding"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/yikakia/awf_calc/game/resource"
	"github.com/yikakia/awf_calc/i18n"
	"github.com/yikakia/awf_calc/internal/tools/gresult"
	"github.com/yikakia/awf_calc/internal/tools/optional"
)

type Type int

const (
	Worker = iota + 1
	Sailor
	Engineer
)

func AllTypes() []Type {
	return []Type{Worker, Sailor, Engineer}
}

func (t Type) String() string {
	switch t {
	case Worker:
		return i18n.CitizenWorkerName.T()
	case Sailor:
		return i18n.CitizenSailorName.T()
	case Engineer:
		return i18n.CitizenEngineerName.T()
	default:
		return "unknown"
	}
}

type Citizen struct {
	Count binding.Int
	// TODO 技能
	skillEnabled binding.Bool
	Type
}

func NewCitizen(citizenType Type, fns ...func(c *Citizen)) Citizen {
	c := &Citizen{
		Count:        binding.NewInt(),
		skillEnabled: binding.NewBool(),
		Type:         citizenType,
	}
	for _, fn := range fns {
		fn(c)
	}

	return *c
}

func NewCitizenWithInt(count int, citizenType Type) Citizen {
	c := NewCitizen(citizenType)
	c.Count.Set(count)
	return c
}

// 如果为空 则认为全部ok
// 不为空 则挨个判断，长度为0表示一个都不允许
func (c Citizen) Satisfy(allows optional.O[[]Type]) bool {
	if allows.IsNil() {
		return true
	}
	for _, allow := range allows.MustGet() {
		if allow == c.Type {
			return true
		}
	}
	return false
}

func (c Citizen) GetCount() int {
	v := gresult.Of(c.Count.Get()).MustGet()
	if v <= 0 {
		return 0
	}
	return v
}

func (c Citizen) getCountI64() int64 {
	return int64(c.GetCount())
}

func (c Citizen) CalcResource(r resource.Resources) resource.Resources {
	add := lo.Map(r.AddAble, func(item resource.Resource, _ int) resource.Resource {
		return c.calcSingleResource(item)
	})

	return resource.Resources{
		AddAble: add,
	}
}

func (c Citizen) calcSingleResource(r resource.Resource) resource.Resource {
	// 都没人的话 置为0
	if c.getCountI64() == 0 {
		r.Count = decimal.Zero
		return r
	}

	// 没有技能 啥也不做
	if gresult.Of(c.Count.Get()).MustGet() == 0 {
		return r
	}

	// 不能减到小于0
	switch c.Type {
	case Worker:
		return r
	case Sailor:
		r.Count = r.Count.Sub(decimal.NewFromInt(c.getCountI64()))
		return r
	case Engineer:
		r.Count = r.Count.Add(decimal.NewFromInt(c.getCountI64()))
		return r
	default:
		slog.Error("unknown resource type", slog.Any("type", c.Type))
		return r
	}
}
