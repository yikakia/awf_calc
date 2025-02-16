package citizen

import (
	"github.com/samber/lo"
	"github.com/yikakia/awf_calc/game/resource"
)

func Which(resources resource.Resources) Type {
	p := lo.SumBy(resources.AddAble, func(item resource.Resource) float64 {
		return item.Count.InexactFloat64()
	})
	if p == 0 {
		return Worker
	} else if p > 0 {
		return Sailor
	} else {
		return Engineer
	}
}
