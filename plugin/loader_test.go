// Copyright 2023 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0
package plugin

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_getTypeName(t *testing.T) {
	n := getTypeName[IPluggable]()
	assert.Equal(t, "IPluggable", n)
}

