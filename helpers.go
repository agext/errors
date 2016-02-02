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

// Error levels match logging levels in agext/log
const (
	WARNING int8 = iota + 2
	ERROR
	PANIC
	FATAL

	minLevel = WARNING
	maxLevel = FATAL
)

// Predefined error codes
const (
	ERR_NEW_ARG int = iota
)

func levelName(l int8) string {
	switch l {
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case PANIC:
		return "PANIC"
	case FATAL:
		return "FATAL"
	}
	return "?"
}

type Logger interface {
	Fatal(...interface{})
	Panic(...interface{})
	Print(...interface{})
}
