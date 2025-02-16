package cfgs

import (
	"time"

	resource2 "github.com/yikakia/awf_calc/game/resource"
	"github.com/yikakia/awf_calc/i18n"
	"github.com/yikakia/awf_calc/internal/tools/gslice"
)

var bs = []building{
	{
		Name:            i18n.BuildingWood_CatcherName.T(),
		UseTime:         4 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceWood.Produce(2),
			),
		},
	},
	{
		Name:            i18n.BuildingJunk_ElectricCatcherName.T(),
		UseTime:         4 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceJunk.Produce(4),
				resource2.ResourceEnergy.Cost(4),
			),
		},
	},
	{
		Name:            i18n.BuildingWood_RecyclingName.T(),
		UseTime:         4 * time.Hour,
		MaxCitizenCount: 3,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceJunk.Cost(12),
				resource2.ResourceMetal.Produce(4)),
		},
	},
	{
		Name:            i18n.BuildingWood_GreenhouseName.T(),
		UseTime:         8 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceWood.Produce(10),
				resource2.ResourceWater.Cost(14)),
		},
	},
	{
		Name:            i18n.BuildingWood_PresserName.T(),
		UseTime:         3 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceWood.Produce(18),
				resource2.ResourceJunk.Cost(9),
				resource2.ResourceEnergy.Cost(4)),
		},
	},
	{
		Name:            i18n.BuildingMetal_ElectricName.T(),
		UseTime:         6 * time.Hour,
		MaxCitizenCount: 4,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceJunk.Cost(12),
				resource2.ResourceMetal.Produce(14),
				resource2.ResourceEnergy.Cost(12),
			),
		},
	},
	{
		Name:            i18n.BuildingMetal_FuelName.T(),
		UseTime:         6 * time.Hour,
		MaxCitizenCount: 4,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceMetal.Produce(16),
				resource2.ResourceJunk.Cost(16),
				resource2.ResourceFuel.Cost(8),
			),
		},
	},
	{
		Name:            i18n.BuildingWater_WoodName.T(),
		UseTime:         6 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceWood.With(-6),
				resource2.ResourceWater.With(9),
			),
		},
	},
	{
		Name:            i18n.BuildingWater_JunkName.T(),
		UseTime:         4 * time.Hour,
		MaxCitizenCount: 3,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceJunk.With(-8),
				resource2.ResourceWater.With(14),
			),
		},
	},
	{
		Name:            i18n.BuildingWater_FuelName.T(),
		UseTime:         4 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceFuel.With(-5),
				resource2.ResourceWater.With(20),
			),
		},
	},
	{
		Name:            i18n.BuildingWater_ElectricName.T(),
		UseTime:         4 * time.Hour,
		MaxCitizenCount: 3,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceWater.With(14),
				resource2.ResourceEnergy.Cost(10),
			),
		},
	},
	{
		Name:            i18n.BuildingFood_FishingHutName.T(),
		UseTime:         5 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.New(resource2.ResourceFood, 1),
			),
		},
	},
	{
		Name:            i18n.BuildingFood_GreenhouseName.T(),
		UseTime:         6 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.New(resource2.ResourceWater, -12),
				resource2.New(resource2.ResourceFood, 12),
			),
		},
	},
	{
		Name:            i18n.BuildingFood_MushroomsName.T(),
		UseTime:         6 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.New(resource2.ResourceWater, -5),
				resource2.New(resource2.ResourceWood, -5),
				resource2.New(resource2.ResourceFood, 10),
			),
		},
	},
	{
		Name:            i18n.BuildingFood_GreenhouseElectricName.T(),
		UseTime:         6 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.New(resource2.ResourceWater, -10),
				resource2.New(resource2.ResourceFood, 14),
				resource2.ResourceEnergy.Cost(6),
			),
		},
	},
	{
		Name:            i18n.BuildingEnergy_HamsterWheelName.T(),
		UseTime:         24 * time.Hour,
		MaxCitizenCount: 1,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceEnergy.Produce(2)),
		},
	},
	{
		Name:            i18n.BuildingEnergy_WoodenWindTurbineName.T(),
		UseTime:         24 * time.Hour,
		MaxCitizenCount: 1,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceEnergy.Produce(3)),
		},
	},
	{
		Name:            i18n.BuildingEnergy_SunBoilerName.T(),
		UseTime:         8 * time.Hour,
		MaxCitizenCount: 2,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceWater.Cost(4),

				resource2.ResourceEnergy.Produce(8)),
		},
	},
	{
		Name:            i18n.BuildingFuel_WoodName.T(),
		UseTime:         6 * time.Hour,
		MaxCitizenCount: 3,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceWood.Cost(14),
				resource2.ResourceJunk.Cost(10),
				resource2.ResourceFuel.Produce(8)),
		},
	},
	{
		Name:            i18n.BuildingEnergy_FuelName.T(),
		UseTime:         8 * time.Hour,
		MaxCitizenCount: 3,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceFuel.Cost(7),
				resource2.ResourceEnergy.Produce(20)),
		},
	},
	{
		Name:            i18n.BuildingFuel_FoodName.T(),
		UseTime:         6 * time.Hour,
		MaxCitizenCount: 3,
		Resources: resource2.Resources{
			AddAble: gslice.Of(
				resource2.ResourceFood.Cost(10),
				resource2.ResourceJunk.Cost(10),
				resource2.ResourceFuel.Produce(12)),
		},
	},
}
