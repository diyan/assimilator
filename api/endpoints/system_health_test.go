package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemHealth_Get(t *testing.T) {
	client, _ := Setup(t)
	defer TearDown(t)
	res, bodyStr, errs := client.Get("http://example.com/api/0/internal/health/").End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, `{
            "healthy": {
                "WarningStatusCheck": false,
                "CeleryAppVersionCheck": true,
                "CeleryAliveCheck": true
            },
            "problems": []
        }`,
		bodyStr)
}
