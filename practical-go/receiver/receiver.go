package receiver

type SensorData struct {
	SensorType string
	ModelID    string
	Value      float32
}

func ReadValue1(r SensorData) float32 {
	if r.SensorType == "Fahrenheit" {
		return (r.Value * 9 / 5) + 32
	}
	return r.Value
}

func (r SensorData) ReadValue2() float32 {
	if r.SensorType == "Fahrenheit" {
		return (r.Value * 9 / 5) + 32
	}
	return r.Value
}
