// Copyright 2017 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package os provides template functions for interacting with the operating
// system.
package exec

import (
	"fmt"
	"bytes"
	"os/exec"

	"github.com/gohugoio/hugo/deps"
	"github.com/spf13/cast"
)

// New returns a new instance of the os-namespaced template functions.
func New(d *deps.Deps) *Namespace {

	return &Namespace{
		deps:       d,
	}
}

// Namespace provides template functions for the "os" namespace.
type Namespace struct {
	deps       *deps.Deps
}

// ReadDir lists the directory contents relative to the configured WorkingDir.
func (ns *Namespace) Exec(command_i interface{}, args_i ...interface{}) (string, error) {
	command, err := cast.ToStringE(command_i)
	if err != nil {
		return "", err
	}

	args := make([]string, len(args_i))
	for i := 0; i < len(args); i++ {
		arg, err := cast.ToStringE(args_i[i])
		if err != nil {
			return "", err
		}
		args[i] = arg
	}

	cmd := exec.Command(command, args...)
	//stdout, err := cmd.Output()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		return "", fmt.Errorf("Command execution failed: %s.\n%s", err.Error(), stderr.String())
	}

	return stdout.String(), nil
}
