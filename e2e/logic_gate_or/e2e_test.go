package test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

// Check that logical OR works
func TestOR(t *testing.T) {
	cmd := exec.Command("neva", "run", "main")

	out, err := cmd.CombinedOutput()
	require.NoError(t, err)
	require.Equal(
		t,
		"false\n",
		string(out),
	)

	require.Equal(t, 0, cmd.ProcessState.ExitCode())
}
