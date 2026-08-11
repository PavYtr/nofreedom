package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/PavYtr/nofreedom/internal/cli"
)

func executeCommand(t *testing.T, args ...string) (stdout string, stderr string, err error) {
	t.Helper()
	cmd := cli.NewRootCmd()

	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&errOut)
	cmd.SetArgs(args)

	err = cmd.Execute()
	stdout = out.String()
	stderr = errOut.String()

	return stdout, stderr, err
}

func TestConvertCommandSuccess(t *testing.T) {
	stdout, stderr, err := executeCommand(t, "convert", "1000", "kg", "g")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if stderr != "" {
		t.Errorf("Expected no error, got: %v", stderr)
	}
	if stdout != "1000000\n" {
		t.Errorf("Expected output '1000000', got: %v", stdout)
	}
}

func TestConvertInvalidValue(t *testing.T) {
	stdout, stderr, err := executeCommand(t, "convert", "invalid", "kg", "g")
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	if stdout != "" {
		t.Errorf("Expected no output, got: %v", stdout)
	}
	if stderr != "" {
		t.Errorf("expected no stderr, got: %v", stderr)
	}
	if !strings.Contains(err.Error(), "invalid value") {
		t.Errorf("Expected error message to contain 'invalid value', got: %v", err)
	}
}

func TestConvertCommandNotEnoughArgs(t *testing.T) {
	_, stderr, err := executeCommand(t, "convert", "1000", "kg")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if stderr != "" {
		t.Errorf("expected no stderr, got: %v", stderr)
	}
	if !strings.Contains(err.Error(), "accepts 3 arg(s), received 2") {
		t.Errorf("Expected error message to contain 'accepts 3 arg(s), received 2', got: %v", err)
	}
}

func TestConvertCommandTooManyArgs(t *testing.T) {
	_, _, err := executeCommand(t, "convert", "1000", "kg", "g", "extra")
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestConvertCommandInvalidUnit(t *testing.T) {
	_, _, err := executeCommand(t, "convert", "1000", "bebra", "obeme")
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if !strings.Contains(err.Error(), "unknown unit") {
		t.Errorf("Expected error message, got nil")
	}
}

func TestListCommand(t *testing.T) {
	stdout, stderr, err := executeCommand(t, "list")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if stderr != "" {
		t.Errorf("Expected no error, got: %v", stderr)
	}
	if !strings.Contains(stdout, "[length]:") || !strings.Contains(stdout, "[mass]:") {
		t.Errorf("Expected output to contain '[length]:' and '[mass]:', got: %v", stdout)
	}
}
