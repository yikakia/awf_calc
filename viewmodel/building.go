package viewmodel

import (
	"time"

	"fyne.io/fyne/v2/data/binding"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	citizen2 "github.com/yikakia/awf_calc/game/citizen"
	"github.com/yikakia/awf_calc/game/resource"
	"github.com/yikakia/awf_calc/internal/tools/gresult"
)

type Building struct {
	cfg     *BuildingConfig
	citizen citizen2.Citizen
	Data    struct {
		// 建筑物名字
		Name string
		// 该建筑物总数
		Count binding.Int
		// 居民类型
		CitizenType citizen2.Type
		// 居民数量
		CitizenCount binding.Int
		// 产出物产数量
		*AllResources
	}
}

type BuildingConfig struct {
	// 建筑物名字
	Name string
	// 建筑物产出时间
	UseTime time.Duration
	// 每个建筑最多工人数量
	MaxCitizenCount int
	// 产出资源列表
	resource.Resources
	// 当 Data.Count 变动时，需要回调的
	OnChanges []func() `json:"-"`
}

func NewBuilding(dep *BuildingConfig, opts ...func(building *Building)) *Building {
	b := &Building{cfg: dep}
	b.init()
	for _, opt := range opts {
		opt(b)
	}

	return b
}

func (b *Building) init() {
	// 判断应该用哪个工人类别
	cType := citizen2.Which(b.cfg.Resources)
	b.citizen = citizen2.NewCitizen(cType)

	// 初始化 Data
	b.Data.Name = b.cfg.Name
	b.Data.Count = binding.NewInt()
	b.Data.CitizenType = b.citizen.Type
	b.Data.CitizenCount = b.citizen.Count
	b.Data.AllResources = NewAllResources()

	// 再注册回调
	b.Data.Count.AddListener(binding.NewDataListener(func() {
		b.updateCitizenCnt()
		go b.calc()
		for _, onChange := range b.cfg.OnChanges {
			onChange()
		}
	}))
}

func (b *Building) calc() {
	// 乘count
	mul := func(item resource.Resource, _ int) resource.Resource {
		cnt := gresult.Of(b.Data.Count.Get()).ValueOrZero()
		item.Count = item.Count.Mul(decimal.NewFromInt(int64(cnt)))
		return item
	}
	newR := resource.Resources{
		AddAble: lo.Map(b.cfg.Resources.AddAble, mul),
	}

	// 加减数量
	afterCitizen := b.citizen.CalcResource(newR)
	b.Data.AllResources.Clear()
	for _, r := range afterCitizen.AddAble {
		b.Data.AllResources.Add(r.Type, r.Count.InexactFloat64())
	}

	// 除以时间
	for _, t := range resource.AllTypes() {
		cnt := b.Data.AllResources.GetBy(t)
		tmp := gresult.Of(cnt.Get()).MustGet()
		if t.CanDivByHour() {
			b.Data.AllResources.Set(t, tmp/b.cfg.UseTime.Hours())
		}

	}

}

func (b *Building) updateCitizenCnt() {
	b.citizen.Count.Set(b.cfg.MaxCitizenCount * gresult.Of(b.Data.Count.Get()).ValueOrZero())
}
