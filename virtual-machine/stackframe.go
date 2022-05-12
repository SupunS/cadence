/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package virtual_machine

type StackFrame struct {
	values []Value
}

func NewStackFrame() *StackFrame {
	return &StackFrame{
		values: make([]Value, 0, 10),
	}
}

func (s *StackFrame) Push(v Value) {
	s.values = append(s.values, v)
}

func (s *StackFrame) Pop() Value {
	// FIXME: handle empty Stack
	l := len(s.values)
	top := s.values[l-1]
	s.values = s.values[:l-1]
	return top
}

func (s *StackFrame) Set(index int, v Value) {
	if index > len(s.values) {
		panic("invalid index")
	}

	// If it's the first time var is stored, allocate a new memory location
	if index == len(s.values) {
		s.values = append(s.values, v)
		return
	}

	s.values[index] = v
}

func (s *StackFrame) Get(index int) Value {
	return s.values[index]
}
