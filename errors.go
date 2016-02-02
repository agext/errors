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
	"runtime"
)

// Error represents an error descriptor capable of storing more detailed
// information than possible with the standard errors package. It ensures that
// any implementation also satisfies the built-in `error` interface.
type Error interface {
	error
	Level() int8
	SetLevel(int8) Error
	Code() int
	SetCode(int) Error
	Text() string
	SetText(string) Error
	Info() []string
	AddInfo(...string) Error
	Log(Logger) Error
}

// Desc provides a means to convey detailed error information to New.
type Desc struct {
	Level int8
	Code  int
	Text  string
	Info  []string
}

// errorMessage stores information about one error occurrence. Pointers to it
// implement the Error interface.
type errorMessage struct {
	level int8
	code  int
	text  string
	info  []string
}

// New returns an error descriptor containing the given information. It accepts
// either a string or a Desc structure (or a pointer to either).
//
// It is a drop-in replacement for the corresponding function from the standard package.
func New(desc interface{}) Error {
	switch desc := desc.(type) {
	case string:
		return &errorMessage{level: ERROR, text: desc}
	case *string:
		return &errorMessage{level: ERROR, text: *desc}
	case Desc:
		return newFromE(&desc)
	case *Desc:
		return newFromE(desc)
	}
	return newFromE(&Desc{
		Code: ERR_NEW_ARG,
		Text: fmt.Sprintf("unsupported error descriptor type %T", desc),
		Info: []string{
			fmt.Sprintf("%T", desc),
			"debug.stack",
		},
	})
}

func newFromE(desc *Desc) Error {
	return (&errorMessage{
		level: ERROR,
		code:  desc.Code,
		text:  desc.Text,
	}).addInfo(3, desc.Info...).SetLevel(desc.Level)
}

// Log sends the error to the provided log, using the appropriate
// logging function: FATAL conditions are logged using Fatal(), PANIC using
// Panic(), and anything else using Print().
func (this *errorMessage) Log(log Logger) Error {
	switch this.level {
	case FATAL:
		log.Fatal(this)
	case PANIC:
		log.Panic(this)
	default:
		log.Print(this)
	}
	return this
}

// Level returns the error level.
func (this *errorMessage) Level() int8 {
	return this.level
}

// SetLevel sets the error level.
func (this *errorMessage) SetLevel(l int8) Error {
	if l >= minLevel && l <= maxLevel {
		this.level = l
	}
	return this
}

// Code returns the error code.
func (this *errorMessage) Code() int {
	return this.code
}

// SetCode sets the error code.
func (this *errorMessage) SetCode(c int) Error {
	this.code = c
	return this
}

// Text returns the error text.
func (this *errorMessage) Text() string {
	return this.text
}

// SetText sets the error text.
func (this *errorMessage) SetText(t string) Error {
	this.text = t
	return this
}

// Info returns the error info.
func (this *errorMessage) Info() []string {
	return this.info
}

// addInfo adds (more) error info.
func (this *errorMessage) addInfo(calldepth int, s ...string) Error {
	for i, line := range s {
		if line == "debug.stack" {
			calldepth *= 2
			buffer := make([]byte, 4096)
			buffer = buffer[:runtime.Stack(buffer, true)]
			var p1, p2, l int
			for j, c := range buffer {
				if c == 10 {
					if l == 0 {
						p1 = j + 1
					} else if l == calldepth {
						p2 = j + 1
						break
					}
					l++
				}
			}
			if p2 > 0 {
				s[i] = string(buffer[:p1]) + string(buffer[p2:])
			} else {
				s[i] = string(buffer)
			}
			break
		}
	}
	this.info = append(this.info, s...)
	return this
}

// AddInfo adds (more) error info.
func (this *errorMessage) AddInfo(s ...string) Error {
	return this.addInfo(2, s...)
}

// Error returns a text containing the error message and code;
// it is useful for satisfying the `error` interface.
func (this *errorMessage) Error() string {
	if this.code != 0 {
		return this.text + fmt.Sprintf(" (code: 0x%04x)", this.code)
	}
	return this.text
}
