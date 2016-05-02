// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package zbark_test

import (
	"github.com/uber-common/zap"
	"github.com/uber-common/zap/zbark"
)

func Example() {
	zapLogger := zap.NewJSON()
	// Stub the current time in tests.
	zapLogger.StubTime()

	// Wrap our structured logger to mimic bark.Logger.
	logger := zbark.Barkify(zapLogger)

	// The wrapped logger has bark's usual fluent API.
	logger.WithField("errors", 0).Infof("%v accepts arbitrary types.", "Bark")

	// Output:
	// {"msg":"Bark accepts arbitrary types.","level":"info","ts":0,"fields":{"errors":0}}
}