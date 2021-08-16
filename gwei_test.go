package unit

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewGWei(t *testing.T) {
	type args struct {
		value *big.Float
	}

	ether10, _ := new(big.Int).SetString("10000000000000000000", 10)
	ether100, _ := new(big.Int).SetString("100000000000000000000", 10)
	ether1000, _ := new(big.Int).SetString("1000000000000000000000", 10)
	ether10000000, _ := new(big.Int).SetString("10000000000000000000000000", 10)

	tests := []struct {
		name string
		args args
		want *Unit
	}{
		{
			name: "10000 wei",
			args: args{value: big.NewFloat(0.00001)},
			want: &Unit{Value: big.NewInt(10000)},
		},
		{
			name: "100000 GWei",
			args: args{value: big.NewFloat(100000)},
			want: &Unit{Value: big.NewInt(100000000000000)},
		},
		{
			name: "0.1 ether",
			args: args{value: big.NewFloat(100000000)},
			want: &Unit{Value: big.NewInt(100000000000000000)},
		},
		{
			name: "1 ether",
			args: args{value: big.NewFloat(1000000000)},
			want: &Unit{Value: big.NewInt(1000000000000000000)},
		},
		{
			name: "10 ether",
			args: args{value: big.NewFloat(10000000000)},
			want: &Unit{Value: ether10},
		},
		{
			name: "100 ether",
			args: args{value: big.NewFloat(100000000000)},
			want: &Unit{Value: ether100},
		},
		{
			name: "1000 ether",
			args: args{value: big.NewFloat(1000000000000)},
			want: &Unit{Value: ether1000},
		},
		{
			name: "10000000 ether",
			args: args{value: big.NewFloat(10000000000000000)},
			want: &Unit{Value: ether10000000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGWei(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGWei() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_GWei(t *testing.T) {
	type fields struct {
		Value *big.Int
	}

	ether10, _ := new(big.Int).SetString("10000000000000000000", 10)
	ether100, _ := new(big.Int).SetString("100000000000000000000", 10)
	ether1000, _ := new(big.Int).SetString("1000000000000000000000", 10)
	ether10000000, _ := new(big.Int).SetString("10000000000000000000000000", 10)
	decEther, _ := new(big.Int).SetString("91231929312891218293912932993", 10)
	veryBigEther, _ := new(big.Int).SetString("99999999999999999899879999999999999999999999999999999999999999999999999999999000000000000000000", 10)

	tests := []struct {
		name   string
		fields fields
		want   *big.Float
	}{
		{
			name:   "1 ether",
			fields: fields{Value: big.NewInt(1000000000000000000)},
			want:   big.NewFloat(1000000000),
		},
		{
			name:   "10 ether",
			fields: fields{Value: ether10},
			want:   big.NewFloat(10000000000),
		},
		{
			name:   "100 ether",
			fields: fields{Value: ether100},
			want:   big.NewFloat(100000000000),
		},
		{
			name:   "1000 ether",
			fields: fields{Value: ether1000},
			want:   big.NewFloat(1000000000000),
		},
		{
			name:   "10000000 ether",
			fields: fields{Value: ether10000000},
			want:   big.NewFloat(10000000000000000),
		},
		{
			name:   "91231929312.891218293912932993 ether",
			fields: fields{Value: decEther},
			want:   big.NewFloat(91231929312891218293.912932993),
		},
		{
			name:   "a lot ether",
			fields: fields{Value: veryBigEther},
			want:   big.NewFloat(99999999999999999899879999999999999999999999999999999999999999999999999999999000000000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Unit{
				Value: tt.fields.Value,
			}
			if got := u.GWei(); got.String() != tt.want.String() {
				t.Errorf("GWei() = %s, want %s", got.String(), tt.want.String())
			}
		})
	}
}
