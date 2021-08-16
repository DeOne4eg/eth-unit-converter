package unit

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewWei(t *testing.T) {
	type args struct {
		value *big.Int
	}

	i, _ := new(big.Int).SetString("100000000000000000000000", 10)

	tests := []struct {
		name string
		args args
		want *Unit
	}{
		{
			name: "100000 eth",
			args: args{value: i},
			want: &Unit{Value: i},
		},
		{
			name: "1521952167015402080 wei",
			args: args{value: big.NewInt(1521952167015402080)},
			want: &Unit{Value: big.NewInt(1521952167015402080)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWei(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWei() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_Wei(t *testing.T) {
	type fields struct {
		Value *big.Int
	}

	tests := []struct {
		name   string
		fields fields
		want   *big.Int
	}{
		{
			name:   "1521952167015402080 wei",
			fields: fields{Value: big.NewInt(1521952167015402080)},
			want:   big.NewInt(1521952167015402080),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Unit{
				Value: tt.fields.Value,
			}
			if got := u.Wei(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wei() = %v, want %v", got, tt.want)
			}
		})
	}
}
