package probe_status

import (
	"os"
	"testing"

	kubeprobev1 "github.com/erda-project/kubeprober/apis/v1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func initEvn() {
	_ = os.Setenv(kubeprobev1.ProbeNamespace, "default")
	_ = os.Setenv(kubeprobev1.ProbeName, "probe-link-test1")
	_ = os.Setenv(kubeprobev1.ProbeItemName, "probe-item-test2")
	_ = os.Setenv(kubeprobev1.ProbeStatusReportUrl, "http://localhost:8081/probe-status")
}

func TestReportStatus(t *testing.T) {
	initEvn()
	now := metav1.Now()
	checker1 := kubeprobev1.ProbeCheckerStatus{
		Name:    "probe-checker-test1",
		Status:  kubeprobev1.CheckerStatusInfo,
		Message: "probe-checker-test1 error",
		LastRun: nil,
	}
	checker2 := kubeprobev1.ProbeCheckerStatus{
		Name:    "probe-checker-test2",
		Status:  kubeprobev1.CheckerStatusInfo,
		Message: "probe-checker-test2 info",
		LastRun: &now,
	}
	checkers := []kubeprobev1.ProbeCheckerStatus{checker1, checker2}
	err := ReportProbeStatus(checkers)
	assert.NoError(t, err)
}