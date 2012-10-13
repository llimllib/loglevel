// Copyright 2012 The Oneslang Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log_test

import (
	"../log"
	"../log/layout"
	"../log/level"
	"testing"
)

func TestFatal(t *testing.T) {
	log.SetLevels(level.Trace)
	log.SetLayouts(layout.Default | layout.Level | layout.Longfile)
	log.Debug("aaa")
}
