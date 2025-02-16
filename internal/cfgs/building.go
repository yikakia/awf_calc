package cfgs

import (
	_ "embed"
	"encoding/json"
	"sync"
	"time"

	"github.com/samber/lo"
	"github.com/yikakia/awf_calc/game/resource"
	"github.com/yikakia/awf_calc/viewmodel"
)

//go:embed buildings.json
var buildingsFile []byte

var buildings = sync.OnceValue(func() []*viewmodel.BuildingConfig {
	var rs []building
	err := json.Unmarshal(buildingsFile, &rs)
	if err != nil {
		panic(err)
	}
	return lo.Map(rs, func(item building, _ int) *viewmodel.BuildingConfig {
		return &viewmodel.BuildingConfig{
			Name:            item.Name,
			UseTime:         item.UseTime,
			MaxCitizenCount: item.MaxCitizenCount,
			Resources:       item.Resources,
		}
	})
})

type building struct {
	// 建筑物名字
	Name string
	// 建筑物产出时间
	UseTime time.Duration
	// 每个建筑最多工人数量
	MaxCitizenCount int
	// 产出资源列表
	resource.Resources
}

func GetBuildingConfigs() []*viewmodel.BuildingConfig {
	return lo.Map(bs, func(item building, _ int) *viewmodel.BuildingConfig {
		return &viewmodel.BuildingConfig{
			Name:            item.Name,
			UseTime:         item.UseTime,
			MaxCitizenCount: item.MaxCitizenCount,
			Resources:       item.Resources,
		}
	})
}
