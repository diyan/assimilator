package interfaces

type Device struct {
}

func (*Device) KeyAlias() string {
	return "device"
}

func (*Device) KeyCanonical() string {
	return "sentry.interfaces.Device"
}

func (device *Device) DecodeRecord(record interface{}) error {
	return DecodeRecord(record, device)
}

func (device *Device) DecodeRequest(request map[string]interface{}) error {
	return DecodeRequest(request, device)
}
