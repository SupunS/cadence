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

package instructions

import vm "github.com/onflow/cadence/virtual-machine/stack-vm"

// Contains all jump instructions.

// GOTO instruction
type GOTO struct {
	Instruction int
}

var _ vm.Instruction = GOTO{}

func (i GOTO) Execute(vm *vm.VirtualMachine) {
	vm.NextIndex = i.Instruction
}

// ICOMP instruction
type ICOMP struct {
	Instruction int // instruction to jump to, if false
}

var _ vm.Instruction = ICOMP{}

func (i ICOMP) Execute(vm *vm.VirtualMachine) {
	if vm.Stack.Pop().(int) != vm.Stack.Pop().(int) {
		vm.NextIndex = i.Instruction
	}
}
