// Copyright Daytona Platforms Inc.
// SPDX-License-Identifier: AGPL-3.0

package internal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func RequireInteractive(action, fix string) error {
	if NoInput {
		return fmt.Errorf("%s requires interactive input, but --no-input was provided; %s", action, fix)
	}

	if !isTerminal(os.Stdin) || !isTerminal(os.Stdout) {
		return fmt.Errorf("%s requires an interactive terminal; %s", action, fix)
	}

	return nil
}

func isTerminal(file *os.File) bool {
	return file != nil && term.IsTerminal(int(file.Fd()))
}
