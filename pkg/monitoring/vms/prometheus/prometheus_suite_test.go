package prometheus

import (
	"testing"

	"kubevirt.io/client-go/testutils"
)

func TestPrometheus(t *testing.T) {
	testutils.KubeVirtTestSuiteSetup(t, "Prometheus Suite")
}
