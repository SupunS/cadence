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

type CallStack struct {
	stackFrames []*StackFrame

	// cache for fast access
	top *StackFrame
}

var callStackPool = sync.Pool{
	New: func() interface{} {
		return &CallStack{
			stackFrames: make([]*StackFrame, 0, 10),
		}
	},
}

func NewCallStack() *CallStack {
	callstack := callStackPool.Get().(*CallStack)
	callstack.PushNew()

	return callstack
}

func (s *CallStack) PushNew() {
	s.Push(NewStackFrame())
}

func (s *CallStack) Push(frame *StackFrame) {
	s.top = frame
	s.stackFrames = append(s.stackFrames, frame)
}

func (s *CallStack) Pop() *StackFrame {
	// FIXME: handle empty Stack
	l := len(s.stackFrames)
	top := s.stackFrames[l-1]
	s.stackFrames = s.stackFrames[:l-1]

	// update the new top of stack
	s.top = s.stackFrames[l-2]

	return top
}

func (s *CallStack) Top() *StackFrame {
	return s.top
}

func (s *CallStack) Clear() {
	s.top = nil
	s.stackFrames = s.stackFrames[:0]
}
