package interfaces

type Device struct {
}

func (device *Device) DecodeRecord(record interface{}) error {
	return nil
}

func (device *Device) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("device", "sentry.interfaces.Device", request, device)
	return err
}
