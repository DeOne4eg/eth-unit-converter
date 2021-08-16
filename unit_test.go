package unit

import (
	"math/big"
	"reflect"
	"testing"
)

func TestUnit_String(t *testing.T) {
	type fields struct {
		Value *big.Int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "1 eth",
			fields: fields{Value: big.NewInt(1000000000000000000)},
			want:   "1000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Unit{
				Value: tt.fields.Value,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPrecision(t *testing.T) {
	type args struct {
		prec uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "set 10",
			args: args{prec: 10},
		},
		{
			name: "set 2048",
			args: args{prec: 2048},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetPrecision(tt.args.prec)
			if precision != tt.args.prec {
				t.Errorf("precision = %v, want %v", precision, tt.args.prec)
			}
		})
	}
}

func TestParseUnit(t *testing.T) {
	type args struct {
		value *big.Float
		unit  string
	}

	eth3, _ := new(big.Int).SetString("3000000000000000000", 10)

	tests := []struct {
		name    string
		args    args
		want    *Unit
		wantErr bool
	}{
		{
			name:    "Ether",
			args:    args{value: big.NewFloat(3), unit: "ether"},
			want:    &Unit{Value: eth3},
		},
		{
			name:    "Wei",
			args:    args{value: big.NewFloat(3000000000000000000), unit: "wei"},
			want:    &Unit{Value: eth3},
		},
		{
			name:    "KWei",
			args:    args{value: big.NewFloat(3000000000000000), unit: "kwei"},
			want:    &Unit{Value: eth3},
		},
		{
			name:    "MWei",
			args:    args{value: big.NewFloat(3000000000000), unit: "MWei"},
			want:    &Unit{Value: eth3},
		},
		{
			name:    "GWei",
			args:    args{value: big.NewFloat(3000000000), unit: "gWei"},
			want:    &Unit{Value: eth3},
		},
		{
			name:    "Szabo",
			args:    args{value: big.NewFloat(3000000), unit: "szabo"},
			want:    &Unit{Value: eth3},
		},
		{
			name:    "Finney",
			args:    args{value: big.NewFloat(3000), unit: "finney"},
			want:    &Unit{Value: eth3},
		},
		{
			name: "error",
			args: args{value: big.NewFloat(100), unit: "satoshi"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUnit(tt.args.value, tt.args.unit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUnit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUnit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
