// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package auth

import (
	"strings"
	"testing"

	"github.com/daytona/clients/cli/internal"
)

func TestLoginNoInputRequiresAPIKey(t *testing.T) {
	oldNoInput := internal.NoInput
	oldAPIKeyFlag := apiKeyFlag
	internal.NoInput = true
	apiKeyFlag = ""
	t.Cleanup(func() {
		internal.NoInput = oldNoInput
		apiKeyFlag = oldAPIKeyFlag
	})

	err := LoginCmd.RunE(LoginCmd, nil)
	if err == nil {
		t.Fatal("LoginCmd.RunE() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "--api-key") || !strings.Contains(err.Error(), "DAYTONA_API_KEY") {
		t.Fatalf("LoginCmd.RunE() error = %q, want API key fix", err)
	}
}
