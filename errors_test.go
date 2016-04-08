// Copyright 2015 ALRUX Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"fmt"
	"strings"
	"testing"
)

type mockLogger struct {
	log string
}

func (this *mockLogger) Print(s ...interface{}) {
	this.log += fmt.Sprintln(s...)
}
func (this *mockLogger) Fatal(s ...interface{}) {
	this.log += "[FATAL] " + fmt.Sprintln(s...)
}
func (this *mockLogger) Panic(s ...interface{}) {
	this.log += "[PANIC] " + fmt.Sprintln(s...)
}

func TestNew(t *testing.T) {
	// Different allocations should not be equal.
	if New("abc") == New("abc") {
		t.Errorf(`New("abc") == New("abc")`)
	}
	if New("abc") == New("xyz") {
		t.Errorf(`New("abc") == New("xyz")`)
	}

	err := New(17)
	if err.Code() != ERR_NEW_ARG {
		t.Errorf(`New(17).Code() = %q, want %q`, err.Code(), ERR_NEW_ARG)
	}
	if err.Text() != "unsupported error descriptor type int" {
		t.Errorf(`New(17).Text() = %q, want %q`, err.Text(), "unsupported error descriptor type int")
	}
	if len(err.Info()) != 2 {
		t.Errorf(`len(New(17).Info()) = %q, want %q`, len(err.Info()), 2)
	} else {
		if err.Info()[0] != "int" {
			t.Errorf(`New(17).Info()[0] = %q, want %q`, err.Info()[0], "int")
		}
		if !strings.Contains(err.Info()[1], "goroutine") || !strings.Contains(err.Info()[1], "errors.TestNew") {
			t.Errorf(`New(17).Info()[1] does not contain stack trace (got %q)`, err.Info()[1])
		}
	}
}

func TestError(t *testing.T) {
	err := New("abc")
	if err.Error() != "abc" {
		t.Errorf(`New("abc").Error() = %q, want %q`, err.Error(), "abc")
	}
	s := "xyz"
	err = New(&s)
	if err.Error() != "xyz" {
		t.Errorf(`New(&"xyz").Error() = %q, want %q`, err.Error(), "xyz")
	}
	err = New(&Desc{Code: 1, Text: "abc"})
	if err.Error() != "abc (code: 0x0001)" {
		t.Errorf(`New(&Desc{Code: 1, Text: "abc"}).Error() = %q, want %q`, err.Error(), "abc (code: 0x0001)")
	}
	err = New(Desc{Code: 1, Text: "xyz"})
	if err.Error() != "xyz (code: 0x0001)" {
		t.Errorf(`New(Desc{Code: 1, Text: "xyz"}).Error() = %q, want %q`, err.Error(), "xyz (code: 0x0001)")
	}
}

func TestLevelNames(t *testing.T) {
	for level, name := range map[int8]string{
		WARNING: "WARNING",
		ERROR:   "ERROR",
		PANIC:   "PANIC",
		FATAL:   "FATAL",
		17:      "?",
	} {
		if levelName(level) != name {
			t.Errorf(`levelName(%d) = %q, want %q`, level, levelName(level), name)
		}
	}
}

func TestSetters(t *testing.T) {
	err := New("abc")
	if err.SetLevel(FATAL).Level() != FATAL {
		t.Errorf(`SetLevel(FATAL).Level() = %q, want %q`, err.Level(), levelName(FATAL))
	}
	if err.SetCode(17).Code() != 17 {
		t.Errorf(`SetCode(17).Code() = %q, want %q`, err.Code(), 17)
	}
	if err.SetText("xyz").Text() != "xyz" {
		t.Errorf(`SetText("xyz").Text() = %q, want %q`, err.Text(), "xyz")
	}
	info := err.AddInfo("line 1", "line 2", "debug.stack").Info()
	if len(info) != 3 {
		t.Errorf(`len(AddInfo("line 1", "line 2", "debug.stack").Info()) = %q, want %q`, len(info), 3)
	} else {
		if info[0] != "line 1" {
			t.Errorf(`AddInfo("line 1", "line 2", "debug.stack").Info()[0] = %q, want %q`, info[0], "line 1")
		}
		if info[1] != "line 2" {
			t.Errorf(`AddInfo("line 1", "line 2", "debug.stack").Info()[1] = %q, want %q`, info[1], "line 2")
		}
		if !strings.Contains(info[2], "goroutine") || !strings.Contains(info[2], "errors.TestSetters") {
			t.Errorf(`AddInfo("line 1", "line 2", "debug.stack").Info()[2] does not contain stack trace (got %q)`, info[2])
		}
	}
}

func TestLog(t *testing.T) {
	log := &mockLogger{}
	err := New("abc")
	err.Log(log).SetLevel(PANIC).Log(log).SetLevel(FATAL).Log(log)
	if log.log != "abc\n[PANIC] abc\n[FATAL] abc\n" {
		t.Errorf(`logging test got %q, want %q`, log.log, "abc\n[PANIC] abc\n[FATAL] abc\n")
	}
}
