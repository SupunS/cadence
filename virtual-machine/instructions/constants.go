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
	"github.com/onflow/cadence/runtime/sema"
	vm "github.com/onflow/cadence/virtual-machine"
	"math/big"
)

// IConst instruction
type IConst struct {
	Value *big.Int
}

var _ vm.Instruction = IConst{}

func NewIConst(value *big.Int) IConst {
	return IConst{
		Value: value,
	}
}

func (i IConst) Execute(vm *vm.VirtualMachine) {
	value := interpreter.NewIntValue(i.Value, sema.IntType)
	vm.CurrentStackFrame().Push(value)
}

func (i IConst) String() string {
	return fmt.Sprintf("ICONST %s", i.Value.String())
}
