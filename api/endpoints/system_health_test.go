package api_test

func (t *testSuite) TestSystemHealth_Get() {
	res, bodyStr, errs := t.Client.Get("http://example.com/api/0/internal/health/").End()
	t.Nil(errs)
	t.JSONEq(`{
            "healthy": {
                "WarningStatusCheck": false,
                "CeleryAppVersionCheck": true,
                "CeleryAliveCheck": true
            },
            "problems": []
        }`,
		bodyStr)
	t.Equal(200, res.StatusCode)
}
