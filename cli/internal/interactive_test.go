// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package internal

import (
	"strings"
	"testing"
)

func TestRequireInteractiveFailsWhenNoInput(t *testing.T) {
	oldNoInput := NoInput
	NoInput = true
	t.Cleanup(func() {
		NoInput = oldNoInput
	})

	err := RequireInteractive("test prompt", "provide --value")
	if err == nil {
		t.Fatal("RequireInteractive() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "--no-input") || !strings.Contains(err.Error(), "provide --value") {
		t.Fatalf("RequireInteractive() error = %q, want no-input fix", err)
	}
}

func TestRequireInteractiveFailsWithoutTerminal(t *testing.T) {
	oldNoInput := NoInput
	NoInput = false
	t.Cleanup(func() {
		NoInput = oldNoInput
	})

	err := RequireInteractive("test prompt", "provide --value")
	if err == nil {
		t.Skip("test process is attached to an interactive terminal")
	}
	if !strings.Contains(err.Error(), "interactive terminal") || !strings.Contains(err.Error(), "provide --value") {
		t.Fatalf("RequireInteractive() error = %q, want terminal fix", err)
	}
}
