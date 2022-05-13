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

import (
	"sync"

	"github.com/onflow/cadence/runtime/interpreter"
)

type Value interpreter.Value

type VirtualMachine struct {
	threadPool sync.Pool
}

func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		threadPool: sync.Pool{
			New: func() interface{} {
				return NewExecutionContext()
			},
		},
	}
}

func (vm *VirtualMachine) Execute(ins []Instruction) {
	ctx := vm.NewContext()

	for ctx.NextIndex != 1000 {
		instruction := ins[ctx.NextIndex]
		ctx.NextIndex++

		switch instruction.(type) {
		default:
			instruction.Execute(nil)

		}
	}

	vm.ReleaseContext(ctx)
}

func (vm *VirtualMachine) NewContext() *ExecutionContext {
	ctx := vm.threadPool.Get().(*ExecutionContext)
	ctx.Init()
	return ctx
}

func (vm *VirtualMachine) ReleaseContext(ctx *ExecutionContext) {
	ctx.Clear()
	vm.threadPool.Put(ctx)
}

func staticFunction() {

}
