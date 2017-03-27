package interfaces

type AppleCrashReport struct {
}

func (*AppleCrashReport) KeyAlias() string {
	return "applecrashreport"
}

func (*AppleCrashReport) KeyCanonical() string {
	return "sentry.interfaces.AppleCrashReport"
}

func (report *AppleCrashReport) DecodeRecord(record interface{}) error {
	return DecodeRecord(record, report)
}

func (report *AppleCrashReport) DecodeRequest(request map[string]interface{}) error {
	return DecodeRequest(request, report)
}
