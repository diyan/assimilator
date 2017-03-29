package interfaces

type Device struct {
}

func init() {
	Register(&Device{})
}

func (*Device) KeyAlias() string {
	return "device"
}

func (*Device) KeyCanonical() string {
	return "sentry.interfaces.Device"
}
