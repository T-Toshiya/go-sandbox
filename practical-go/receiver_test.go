package main

import "testing"

func TestReadValue1(t *testing.T) {
	type args struct {
		r SensorData
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadValue1(tt.args.r); got != tt.want {
				t.Errorf("ReadValue1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSensorData_ReadValue2(t *testing.T) {
	type fields struct {
		SensorType string
		ModelID    string
		Value      float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := SensorData{
				SensorType: tt.fields.SensorType,
				ModelID:    tt.fields.ModelID,
				Value:      tt.fields.Value,
			}
			if got := r.ReadValue2(); got != tt.want {
				t.Errorf("ReadValue2() = %v, want %v", got, tt.want)
			}
		})
	}
}
