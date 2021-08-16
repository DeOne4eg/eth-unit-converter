package unit

import (
	"math/big"
	"reflect"
	"testing"
)

func TestUnit_Ether(t *testing.T) {
	type fields struct {
		Value *big.Int
	}

	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "1 ether",
			fields: fields{Value: big.NewInt(1000000000000000000)},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Unit{
				Value: tt.fields.Value,
			}
			if got := u.Ether(); got != tt.want {
				t.Errorf("Ether() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_NewEther(t *testing.T) {
	type args struct {
		value *big.Float
	}

	eth10, _ := new(big.Int).SetString("10000000000000000000", 10)

	tests := []struct {
		name   string
		args   args
		want   *Unit
	}{
		{
			name: "1 ether",
			args: args{value: big.NewFloat(1)},
			want: &Unit{Value: big.NewInt(1e18)},
		},
		{
			name: "10 ether",
			args: args{value: big.NewFloat(10)},
			want: &Unit{Value: eth10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEther(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEther() = %v, want %v", got, tt.want)
			}
		})
	}
}
