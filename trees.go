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
	"strconv"
	"wombat/engine/expr/lex"
)

func isStringOp(node1, node2 lex.Token) bool {
	return node1.Type() == lex.StringLiteral || node2.Type() == lex.StringLiteral
}

func isBoolOp(node1, node2 lex.Token) bool {
	return node1.Type() == lex.BoolLiteral || node2.Type() == lex.BoolLiteral
}


type expression []lex.Token

func (me expression) Value() (lex.Token, bool) {
	expression := []lex.Token(me)
	size := len(expression)
	if size == 0 {
		var t lex.Token
		t.SetInt(0)
		return t, true
	}
	//and/or operators
	for i := 0; i < size; i++ {
		if expression[i].Type() == lex.Symbol {
			o := expression[i]
			switch o.String() {
			case "&&", "||":
				var tree parseTree
				tree.build(expression, i)
				return tree.Value()
			}
		}
	}
	//not operator
	for i := 0; i < size; i++ {
		if expression[i].Type() == lex.Symbol {
			if expression[i].String() == "!" {
				var tree parseTree
				tree.build(expression, i)
				return tree.Value()
			}
		}
	}
	//comparison operators
	for i := 0; i < size; i++ {
		if expression[i].Type() == lex.Symbol {
			o := expression[i]
			switch o.String() {
			case "<", ">", "<=", ">=", "!=", "==":
				var tree parseTree
				tree.build(expression, i)
				return tree.Value()
			}
		}
	}
	//add and subtract
	for i := 0; i < size; i++ {
		if expression[i].Type() == lex.Symbol {
			if expression[i].String() == "+" || expression[i].String() == "-" {
				var tree parseTree
				tree.build(expression, i)
				return tree.Value()
			}
		}
	}
	//multiplication and division
	for i := 0; i < size; i++ {
		if expression[i].Type() == lex.Symbol {
			if expression[i].String() == "*" || expression[i].String() == "/" {
				var tree parseTree
				tree.build(expression, i)
				return tree.Value()
			}
		}
	}

	var t lex.Token
	if expression[0].Type() != lex.Symbol {
		t = expression[0]
	} else if expression[0].Type() == lex.Symbol { //is negative int literal
		var t lex.Token
		if expression[0].String() == "i" {
			retval, _ := strconv.Atoi(expression[1].String())
			t.SetInt(-retval)
			return t, true
		} else {
			return t, false
		}
	}
	return t, true
}


type parseTree struct {
	node1    expression
	operator string
	node2    expression
}

func (me *parseTree) build(expression expression, splitPoint int) {
	me.operator = expression[splitPoint].String()

	me.node1 = expression[0:splitPoint]
	me.node2 = expression[splitPoint+1 : len(expression)]
}

func (me *parseTree) Value() (t lex.Token, ok bool) {
	a, ok := me.node1.Value()
	if !ok {
		return t, false
	}
	b, ok := me.node2.Value()
	if !ok {
		return t, false
	}
	ok = true
	switch me.operator {
	case "+":
		if isStringOp(a, b) {
			t.Set(lex.StringLiteral, a.String()+b.String())
		} else {
			t.Set(lex.IntLiteral, a.Int()+b.Int())
		}
	case "-":
		if isStringOp(a, b) {
			ok = false
		} else {
			t.Set(lex.IntLiteral, a.Int()-b.Int())
		}
	case "*":
		if isStringOp(a, b) {
			ok = false
		} else {
			t.Set(lex.IntLiteral, a.Int()*b.Int())
		}
	case "/":
		if isStringOp(a, b) {
			ok = false
		} else {
			t.Set(lex.IntLiteral, a.Int()/b.Int())
		}
	case "==":
		if isStringOp(a, b) {
			t.Set(lex.StringLiteral, a.String() == b.String())
		} else {
			t.Set(lex.BoolLiteral, a.Int() == b.Int())
		}
	case "!=":
		if isStringOp(a, b) {
			t.Set(lex.StringLiteral, a.String() != b.String())
		} else {
			t.Set(lex.BoolLiteral, a.Int() != b.Int())
		}
	case "<=":
		if isStringOp(a, b) {
			ok = false
		} else {
			t.Set(lex.BoolLiteral, a.Int() <= b.Int())
		}
	case ">=":
		if isStringOp(a, b) {
			ok = false
		} else {
			t.Set(lex.BoolLiteral, a.Int() >= b.Int())
		}
	case "<":
		if isStringOp(a, b) {
			ok = false
		} else {
			t.Set(lex.BoolLiteral, a.Int() < b.Int())
		}
	case ">":
		if isStringOp(a, b) {
			ok = false
		} else {
			t.Set(lex.BoolLiteral, a.Int() > b.Int())
		}
	case "!":
		if b.Type() == lex.BoolLiteral {
			t.SetBool(!b.Bool())
		} else {
			ok = false
		}
	case "&&":
		if isBoolOp(a, b) {
			t.SetBool(a.Bool() && b.Bool())
		} else {
			ok = false
		}
	case "||":
		if isBoolOp(a, b) {
			t.SetBool(a.Bool() || b.Bool())
		} else {
			ok = false
		}
	}
	return t, ok
}
