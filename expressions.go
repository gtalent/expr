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
	"github.com/gtalent/expr/lex"
	"strings"
)

var operators = func() []string {
	var ops []string
	ops = append(ops, "<=")
	ops = append(ops, ">=")
	ops = append(ops, "==")
	ops = append(ops, "!=")
	ops = append(ops, "<")
	ops = append(ops, ">")
	ops = append(ops, "+")
	ops = append(ops, "-")
	ops = append(ops, "*")
	ops = append(ops, "/")
	ops = append(ops, "(")
	ops = append(ops, ")")
	return ops
}()

func fixNegatives(exp []lex.Token) []lex.Token {
	var prev lex.Token
	prev.Set(lex.Symbol, "")
	for i, a := range exp {
		if prev.Type() == lex.Symbol && a.String() == "-" &&
			(prev.String() != ")" || prev.String() == "") {
			exp[i].Set(lex.Symbol, "i")
		}
		prev = a
	}
	return exp
}

func fixIdents(exp string, vars func(string) (Value, bool)) ([]lex.Token, bool) {
	toks := lex.Tokens(exp)
	for i := 0; i < len(toks); i++ {
		t := &toks[i]
		if t.Type() == lex.Identifier {
			v, ok := vars(t.String())
			if !ok {
				return toks, false
			}
			if v.Type == String {
				t.Set(lex.StringLiteral, v.Value.(string))
			}
			if v.Type == Int {
				t.Set(lex.IntLiteral, v.Value.(int))
			}
		}
	}
	return toks, true
}

// Evaluate returns a Value object, and an "ok" bool that indicates whether or not the expression could be parsed.
func Evaluate(expression string, variables func(string) (Value, bool)) (Value, bool) {
	expression = strings.Replace(expression, " ", "", -1)
	expression = strings.Replace(expression, "\n", "", -1)
	expression = strings.Replace(expression, "\t", "", -1)

	toks, ok := fixIdents(expression, variables)
	toks = fixNegatives(toks)

	var val Value
	if !ok {
		return val, false
	}

	tok, ok := solve(toks)
	switch tok.Type() {
	case lex.IntLiteral:
		val.Type = Int
	case lex.StringLiteral:
		val.Type = String
	case lex.BoolLiteral:
		val.Type = Bool
	}
	val.Value = tok.Value()
	return val, ok
}

func solve(exp []lex.Token) (lex.Token, bool) {
	//recursively solve the contents of parentheses place the values in expression
	for i := 0; i < len(exp); i++ {
		if exp[i].Type() == lex.Symbol && exp[i].String() == "(" {
			depth := 0
			length := 1
			var subQuery []lex.Token
			for n := i + 1; true; n++ {
				length++
				if exp[n].Type() == lex.Symbol {
					if exp[n].String() == "(" {
						depth++
					} else if exp[n].String() == ")" {
						if depth == 0 {
							break
						}
						depth--
					}
				}
				subQuery = append(subQuery, exp[n])
			}
			v, ok := solve(subQuery)
			if !ok {
				return v, false
			}

			l1 := exp[:i]
			l2 := exp[i+length:]
			l1 = append(l1, v)
			for _, a := range l2 {
				l1 = append(l1, a)
			}
			exp = l1
		}
	}
	ret := expression(exp)
	return ret.Value()
}
