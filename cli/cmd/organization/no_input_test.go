// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package organization

import (
	"strings"
	"testing"

	"github.com/daytona/clients/cli/internal"
)

func TestUseNoInputRequiresOrganizationArgument(t *testing.T) {
	oldNoInput := internal.NoInput
	internal.NoInput = true
	t.Cleanup(func() {
		internal.NoInput = oldNoInput
	})

	err := UseCmd.RunE(UseCmd, nil)
	if err == nil {
		t.Fatal("UseCmd.RunE() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "organization argument") || !strings.Contains(err.Error(), "organization use") {
		t.Fatalf("UseCmd.RunE() error = %q, want organization argument fix", err)
	}
}

func TestDeleteNoInputRequiresOrganizationArgument(t *testing.T) {
	oldNoInput := internal.NoInput
	internal.NoInput = true
	t.Cleanup(func() {
		internal.NoInput = oldNoInput
	})

	err := DeleteCmd.RunE(DeleteCmd, nil)
	if err == nil {
		t.Fatal("DeleteCmd.RunE() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "organization argument") || !strings.Contains(err.Error(), "organization delete") {
		t.Fatalf("DeleteCmd.RunE() error = %q, want organization argument fix", err)
	}
}
