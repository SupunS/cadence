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

type ExecutionContext struct {
	CallStack *CallStack
	NextIndex int
	currentStackFrame *StackFrame
}

func NewExecutionContext() *ExecutionContext {
	return &ExecutionContext {}
}

func (ctx *ExecutionContext) CurrentStackFrame() *StackFrame {
	return ctx.currentStackFrame
}

func (ctx *ExecutionContext) Init() {
	ctx.CallStack = NewCallStack()
	ctx.currentStackFrame = ctx.CallStack.Top()
}

func (ctx *ExecutionContext) Clear() {
	ctx.NextIndex = 0
	ctx.CallStack.Clear()
}
