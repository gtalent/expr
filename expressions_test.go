/*
   Copyright 2011 gtalent2@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package expr

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	val, ok := Evaluate("-1 + (3 * 7) * 2 + -2", nil)
	if !ok || val.Int() != 39 {
		t.Error("Non-variable multiplication with paretheses is broken.")
	}
	val, ok = Evaluate("(3 * 7) - 2", nil)
	if !ok || val.Int() != 19 {
		t.Error("Non-variable multiplication with paretheses is broken.")
	}
	val, ok = Evaluate("(3 * 7) * 2", nil)
	if !ok || val.Int() != 42 {
		t.Error("Non-variable multiplication with paretheses is broken.")
	}
	val, _ = Evaluate("2 * 21", nil)
	if val.Int() != 42 {
		t.Error("Non-variable multiplication broken.")
	}
	val, _ = Evaluate("2 == 21", nil)
	if val.Bool() {
		t.Error("Non-variable == comparison broken.")
	}
	val, _ = Evaluate("\"Narf!\" == \"Narf!\"", nil)
	if !val.Bool() {
		t.Error("Non-variable == comparison broken.")
	}
	val, _ = Evaluate("\"Narf!\" != \"Narf!\"", nil)
	if val.Bool() {
		t.Error("Non-variable == comparison broken.")
	}
	val, _ = Evaluate("\"Narf!\" != \"Murr!\"", nil)
	if !val.Bool() {
		t.Error("Non-variable == comparison broken.")
	}
	val, _ = Evaluate("21 == 21", nil)
	if !val.Bool() {
		t.Error("Non-variable == comparison broken.")
	}
	val, _ = Evaluate("2 != 21", nil)
	if !val.Bool() {
		t.Error("Non-variable != comparison broken.")
	}
	val, _ = Evaluate("2 != 21 && 4 > 2", nil)
	if !val.Bool() {
		t.Error("Non-variable && comparison broken.")
	}
	val, _ = Evaluate("2 == 21 || 4 > 2", nil)
	if !val.Bool() {
		t.Error("Non-variable || comparison broken.")
	}
	val, _ = Evaluate("!(2 == 21)", nil)
	if !val.Bool() {
		t.Error("Non-variable noted parentheses boolean expressions broken.")
	}
}
