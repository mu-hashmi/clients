// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package config

import (
	"testing"

	"github.com/daytona/clients/cli/internal"
)

func TestGetActiveProfileUsesDefaultAPIURLWithAPIKeyEnv(t *testing.T) {
	oldDaytonaApiUrl := internal.DaytonaApiUrl
	internal.DaytonaApiUrl = ""
	t.Cleanup(func() {
		internal.DaytonaApiUrl = oldDaytonaApiUrl
	})

	t.Setenv(DAYTONA_API_KEY_ENV_VAR, "test-key")
	t.Setenv(DAYTONA_API_URL_ENV_VAR, "")

	profile, err := (&Config{}).GetActiveProfile()
	if err != nil {
		t.Fatalf("GetActiveProfile() returned error: %v", err)
	}

	if profile.Id != "env" {
		t.Fatalf("profile id = %q, want env", profile.Id)
	}
	if profile.Api.Url != defaultDaytonaApiUrl {
		t.Fatalf("profile API URL = %q, want %q", profile.Api.Url, defaultDaytonaApiUrl)
	}
	if profile.Api.Key == nil || *profile.Api.Key != "test-key" {
		t.Fatalf("profile API key = %v, want test-key", profile.Api.Key)
	}
}

func TestGetActiveProfileUsesExplicitAPIURLWithAPIKeyEnv(t *testing.T) {
	t.Setenv(DAYTONA_API_KEY_ENV_VAR, "test-key")
	t.Setenv(DAYTONA_API_URL_ENV_VAR, "https://example.test/api")

	profile, err := (&Config{}).GetActiveProfile()
	if err != nil {
		t.Fatalf("GetActiveProfile() returned error: %v", err)
	}

	if profile.Api.Url != "https://example.test/api" {
		t.Fatalf("profile API URL = %q, want explicit env URL", profile.Api.Url)
	}
}

func TestGetDaytonaApiUrlUsesBuildInfoBeforeDefault(t *testing.T) {
	oldDaytonaApiUrl := internal.DaytonaApiUrl
	internal.DaytonaApiUrl = "https://build.example.test/api"
	t.Cleanup(func() {
		internal.DaytonaApiUrl = oldDaytonaApiUrl
	})

	t.Setenv(DAYTONA_API_URL_ENV_VAR, "")

	if got := GetDaytonaApiUrl(); got != "https://build.example.test/api" {
		t.Fatalf("GetDaytonaApiUrl() = %q, want build info URL", got)
	}
}
