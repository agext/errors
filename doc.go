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

/*
Package errors provides extended error handling.

If you want to retain more information about an error message than a single string allows, just substitute this package for the one in the standard library.

The `New` function still accepts a single string as argument, so no code will be broken. Where you need to include additional information, you can provide it to `New` in a `Desc` structure instead of the string, or you can add it to the error message using one of its setter methods.

The additional information can be used for smarter error handling and logging:
- `Level` differentiates between warnings, regular errors, panics converted to errors, and fatal errors;
- `Code` allows custom classification and prioritizing, by using ranges or bit-level masks;
- `Info` offers a store for arbitrary data and messages, besides the main error `Text`; the special string "debug.stack", if present as an element in the Info slice, is automatically replaced by a stack trace at the point the error message has been created.

*/
package errors
