//go:build mage
// +build mage

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	return sh.Run("go", "build", "-o", "protoc-gen-whatever", "cmd/protoc-gen-whatever/main.go")
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./protoc-gen-whatever", "/usr/bin/protoc-gen-whatever")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	return sh.Run("go", "get", "github.com/stretchr/piglatin")
}

// Clean up after yourself
func Clean() error {
	fmt.Println("Cleaning...")
	os.RemoveAll("protoc-gen-whatever")
	return sh.Rm("dist")
}

func Release() (err error) {
	if os.Getenv("TAG") == "" {
		return errors.New("MSG and TAG environment variables are required")
	}
	if err := sh.RunV("git", "tag", "-a", "$TAG"); err != nil {
		return err
	}
	if err := sh.RunV("git", "push", "origin", "$TAG"); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			sh.RunV("git", "tag", "--delete", "$TAG")
			sh.RunV("git", "push", "--delete", "origin", "$TAG")
		}
	}()
	return sh.RunV("goreleaser")
}
