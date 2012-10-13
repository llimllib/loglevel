// Copyright 2012 The Oneslang Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log_test

import (
	"../log"
	"testing"
)

func TestFatal(t *testing.T) {
	log.SetPriority(log.Pinfo)
	log.SetLayouts(log.Ldefault | log.Lpriority | log.Llongfile)
	log.Fatal("=====fatal=====")
	log.Error("=====error=====")
	log.Warn("=====warn=====")
	log.Info("=====info=====")
	log.Debug("=====debug=====")
	log.Trace("=====trace=====")
}
