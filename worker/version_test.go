package worker_test

import (
	"context"
	"testing"

	"github.com/itsranveer/sj/worker"
	"github.com/stretchr/testify/assert"
)

func TestCompareVersions(t *testing.T) {
	app := worker.NewVersion()
	ctx := context.Background()

	assert.Equal(t, 1, app.CompareVersion(ctx, "1.0.1", "1"))
	assert.Equal(t, 1, app.CompareVersion(ctx, "2.1", "2"))
	assert.Equal(t, 0, app.CompareVersion(ctx, "1.01", "1.0001"))
	assert.Equal(t, 0, app.CompareVersion(ctx, "1.0", "1.0.0.0"))
	assert.Equal(t, -1, app.CompareVersion(ctx, "0.1", "1.1"))
	assert.Equal(t, -1, app.CompareVersion(ctx, "6.7.1.5", "6.7.2"))
	assert.Equal(t, -1, app.CompareVersion(ctx, "6.07.01.5", "6.007.002"))
}
