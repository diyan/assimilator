package interfaces

type AppleCrashReport struct {
}

func (report *AppleCrashReport) DecodeRecord(record interface{}) error {
	return nil
}

func (report *AppleCrashReport) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("applecrashreport", "sentry.interfaces.AppleCrashReport", request, report)
	return err
}
