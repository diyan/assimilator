package interfaces

type AppleCrashReport struct {
}

func init() {
	Register(&AppleCrashReport{})
}

func (*AppleCrashReport) KeyAlias() string {
	return "applecrashreport"
}

func (*AppleCrashReport) KeyCanonical() string {
	return "sentry.interfaces.AppleCrashReport"
}
