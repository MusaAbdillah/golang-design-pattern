package main

import (
	"fmt"
	"testing"
)

func TestGun_setName(t *testing.T) {
	type fields struct {
		name  string
		power int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Success",
			fields: fields{
				name:  "Ak47",
				power: 4,
			},
			args: args{
				name: "ak47",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gun{
				name:  tt.fields.name,
				power: tt.fields.power,
			}
			g.setName(tt.args.name)

			fmt.Println("are you executed?")

		})
	}
}
