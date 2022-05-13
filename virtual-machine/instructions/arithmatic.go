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

import (
	vm "github.com/onflow/cadence/virtual-machine"
)

// IADD instruction
type IADD struct{}

var _ vm.Instruction = IADD{}

func (i IADD) Execute(vm *vm.VirtualMachine) {
	rhsOp := vm.CurrentStackFrame().Pop().(int)
	lhsOp := vm.CurrentStackFrame().Pop().(int)
	vm.CurrentStackFrame().Push(lhsOp + rhsOp)
}

func (i IADD) String() string {
	return "IADD"
}
