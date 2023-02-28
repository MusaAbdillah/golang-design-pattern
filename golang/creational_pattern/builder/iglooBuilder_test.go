package main

import (
	"fmt"
	"testing"
)

func TestIglooBuilder_setWindowType(t *testing.T) {
	type fields struct {
		windowType string
		doorType   string
		floor      int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			fields: fields{
				windowType: "Wood",
				doorType:   "Wood",
				floor:      0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IglooBuilder{
				windowType: tt.fields.windowType,
				doorType:   tt.fields.doorType,
				floor:      tt.fields.floor,
			}
			i.setWindowType()
			if i.windowType != tt.fields.windowType {
				fmt.Errorf("setWindowType is not working")
			}
		})
	}
}
