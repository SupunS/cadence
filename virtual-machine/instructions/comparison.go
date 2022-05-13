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
	"fmt"
	"github.com/onflow/cadence/runtime/interpreter"
	vm "github.com/onflow/cadence/virtual-machine"
)

// Equal instruction
type Equal struct{}

var _ vm.Instruction = Equal{}

var EqualIns = NotEqual{}

func (i Equal) Execute(vm *vm.VirtualMachine) {
	rhsOp := vm.CurrentStackFrame().Pop().(interpreter.EquatableValue)
	lhsOp := vm.CurrentStackFrame().Pop().(interpreter.EquatableValue)

	isEqual := lhsOp.Equal(nil, nil, rhsOp)
	vm.CurrentStackFrame().Push(interpreter.BoolValue(isEqual))
}

func (i Equal) String() string {
	return fmt.Sprintf("EQ")
}

// NotEqual instruction
type NotEqual struct{}

var _ vm.Instruction = NotEqual{}

var NotEqualIns = NotEqual{}

func (i NotEqual) Execute(vm *vm.VirtualMachine) {
	rhsOp := vm.CurrentStackFrame().Pop().(interpreter.EquatableValue)
	lhsOp := vm.CurrentStackFrame().Pop().(interpreter.EquatableValue)

	isEqual := lhsOp.Equal(nil, nil, rhsOp)
	vm.CurrentStackFrame().Push(interpreter.BoolValue(!isEqual))
}

func (i NotEqual) String() string {
	return fmt.Sprintf("NEQ")
}
