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

package benchmarks

import (
	"io/ioutil"

	"github.com/mattermost/logr/v2"
	"github.com/mattermost/logr/v2/formatters"
	"github.com/mattermost/logr/v2/targets"
)

func newDisabledMattermostLogrLogger() *logr.Logger {
	filter := &logr.StdFilter{Lvl: logr.Error, Stacktrace: logr.Panic}
	return newMattermostLogrLoggerWithFilter(filter)
}

func newMattermostLogrLogger() *logr.Logger {
	filter := &logr.StdFilter{Lvl: logr.Debug, Stacktrace: logr.Panic}
	return newMattermostLogrLoggerWithFilter(filter)
}

func newMattermostLogrLoggerWithFilter(filter *logr.StdFilter) *logr.Logger {
	lgr, _ := logr.New()
	formatter := &formatters.JSON{}
	target := targets.NewWriterTarget(ioutil.Discard)
	_ = lgr.AddTarget(target, "writer", filter, formatter, 1000)
	logger := lgr.NewLogger()
	return &logger
}

func fakeMattermostLogrFields() []logr.Field {
	return []logr.Field{
		logr.Int("int", _tenInts[0]),
		logr.Array("ints", _tenInts),
		logr.String("string", _tenStrings[0]),
		logr.Array("strings", _tenStrings),
		logr.Time("time", _tenTimes[0]),
		logr.Array("times", _tenTimes),
		logr.Any("user1", _oneUser),
		logr.Any("user2", _oneUser),
		logr.Array("users", _tenUsers),
		logr.NamedErr("error", errExample),
	}
}
