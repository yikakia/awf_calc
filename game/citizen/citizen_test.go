package citizen

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/yikakia/awf_calc/game/resource"
)

func TestPerson_calcSingleResource(t *testing.T) {
	type args struct {
		r resource.Resource
	}
	tests := []struct {
		name    string
		citizen Citizen
		args    args
		want    resource.Resource
	}{
		{
			name:    "worker do nothing",
			citizen: NewCitizenWithInt(1, Worker),
			args: args{
				r: resource.Resource{
					Type:  resource.ResourceFood,
					Count: decimal.NewFromInt(10),
				},
			},
			want: resource.Resource{
				Type:  resource.ResourceFood,
				Count: decimal.NewFromInt(10),
			},
		},
		{
			name:    "sailor sub them",
			citizen: NewCitizenWithInt(2, Sailor),
			args: args{
				r: resource.Resource{
					Type:  resource.ResourceFood,
					Count: decimal.NewFromInt(10),
				},
			},
			want: resource.Resource{
				Type:  resource.ResourceFood,
				Count: decimal.NewFromInt(8),
			},
		},
		{
			name:    "engineer add them",
			citizen: NewCitizenWithInt(3, Engineer),
			args: args{
				r: resource.Resource{
					Type:  resource.ResourceFood,
					Count: decimal.NewFromInt(10),
				},
			},
			want: resource.Resource{
				Type:  resource.ResourceFood,
				Count: decimal.NewFromInt(13),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.citizen
			if got := p.calcSingleResource(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcSingleResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCitizenWithInt(t *testing.T) {
	c := NewCitizenWithInt(1, Worker)
	c.Count.Set(3)
	assert.Equal(t, 3, c.GetCount())
}
