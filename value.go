/*
   Copyright 2011-2017 gtalent2@gmail.com

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
	"strconv"
)

// Value holds a value in an expression. The value can be a string, int, or bool.
type Value struct {
	Type  int
	Value interface{}
}

// SetInt sets the Value as an in of value val.
func (me *Value) SetInt(val int) {
	me.Type = Int
	me.Value = val
}

// Int returns the value of the Value as an int.
func (me Value) Int() int {
	return me.Value.(int)
}

// SetString sets the Value as an in of value string.
func (me *Value) SetString(val string) {
	me.Type = String
	me.Value = val
}

// String returns the value of the Value as a string.
func (me Value) String() string {
	switch me.Value.(type) {
	case int:
		return strconv.Itoa(me.Value.(int))
	case bool:
		return strconv.FormatBool(me.Value.(bool))
	case string:
		return me.Value.(string)
	}
	return ""
}

// SetBool sets the Value as an in of value bool.
func (me *Value) SetBool(val bool) {
	me.Type = Bool
	me.Value = val
}

// Bool returns the value of the Value as a bool.
func (me Value) Bool() bool {
	return me.Value.(bool)
}
