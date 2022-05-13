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

type Value interface{}

type VirtualMachine struct {
	CallStack *CallStack
	NextIndex int
}

func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{}
}

func (vm *VirtualMachine) Execute(instructions []Instruction) {
	vm.CallStack = NewCallStack()

	for vm.NextIndex != NO_OP {
		instruction := instructions[vm.NextIndex]
		vm.NextIndex++

		instruction.Execute(vm)
	}
}

func (vm *VirtualMachine) CurrentStackFrame() *StackFrame {
	return vm.CallStack.Top()
}
