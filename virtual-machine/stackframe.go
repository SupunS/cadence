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

import "sync"

type StackFrame struct {
	values   []Value
	size     int
	capacity int
}

const initialCapacity = 20
const growFactor = 1.5

var stackFramePool = sync.Pool{
	New: func() interface{} {
		return &StackFrame{
			values:   make([]Value, initialCapacity),
			capacity: initialCapacity,
		}
	},
}

func NewStackFrame() *StackFrame {
	return stackFramePool.Get().(*StackFrame)
}

func (s *StackFrame) Push(v Value) {
	s.grow()
	s.values[s.size] = v
	s.size++
}

func (s *StackFrame) Pop() Value {
	// FIXME: handle empty Stack
	last := s.size - 1
	top := s.values[last]

	//s.values = s.values[:last]

	s.size--

	return top
}

func (s *StackFrame) Set(index int, v Value) {
	if index > s.size {
		panic("invalid index")
	}

	if index < s.size {
		s.values[index] = v
		return
	}

	// If it's the first time var is stored, allocate a new memory location
	s.Push(v)
}

func (s *StackFrame) Get(index int) Value {
	return s.values[index]
}

func (s *StackFrame) Release() {
	s.values = s.values[:0]
	stackFramePool.Put(s)
}

func (s *StackFrame) grow() {
	if s.size != s.capacity {
		return
	}

	values := s.values

	s.capacity = int(float64(s.capacity) * growFactor)

	s.values = make([]Value, s.capacity)
	copy(s.values, values)
}
