/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
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

package wasm

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuf_writeULEB128(t *testing.T) {

	t.Parallel()

	t.Run("DWARF spec", func(t *testing.T) {

		t.Parallel()

		// DWARF Debugging Information Format, Version 3, page 140

		type testCase struct {
			value   uint32
			encoded []byte
		}

		for _, tc := range []testCase{
			{
				value:   2,
				encoded: []byte{2},
			},
			{
				value:   127,
				encoded: []byte{127},
			},
			{
				value:   128,
				encoded: []byte{0 + 0x80, 1},
			},
			{
				value:   129,
				encoded: []byte{1 + 0x80, 1},
			},
			{
				value:   130,
				encoded: []byte{2 + 0x80, 1},
			},
			{
				value:   12857,
				encoded: []byte{57 + 0x80, 100},
			},
		} {
			var b buf
			err := b.writeULEB128(tc.value)
			require.NoError(t, err)
			require.Equal(t, tc.encoded, b.data)

			b.offset = 0

			actual, err := b.readULEB128()
			require.NoError(t, err)
			require.Equal(t, tc.value, actual)
		}
	})

	t.Run("write: max byte count", func(t *testing.T) {

		t.Parallel()

		// This test ensures that only up to the maximum number of bytes are written
		// when writing a LEB128-encoded 32-bit number (see max32bitLEB128ByteCount),
		// i.e. test that only up to 5 bytes are written.

		var b buf
		err := b.writeULEB128(math.MaxUint32)
		require.NoError(t, err)
		require.GreaterOrEqual(t, max32bitLEB128ByteCount, len(b.data))
	})

	t.Run("read: max byte count", func(t *testing.T) {

		t.Parallel()

		// This test ensures that only up to the maximum number of bytes are read
		// when reading a LEB128-encoded 32-bit number (see max32bitLEB128ByteCount),
		// i.e. test that only 5 of the 8 given bytes are read,
		// to ensure the LEB128 parser doesn't keep reading infinitely.

		b := buf{data: []byte{0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88}}
		_, err := b.readULEB128()
		require.NoError(t, err)
		require.Equal(t, offset(max32bitLEB128ByteCount), b.offset)
	})
}

func TestBuf_writeSLEB128(t *testing.T) {

	t.Parallel()

	t.Run("DWARF spec", func(t *testing.T) {

		t.Parallel()

		// DWARF Debugging Information Format, Version 3, page 141

		type testCase struct {
			value   int32
			encoded []byte
		}

		for _, tc := range []testCase{
			{
				value:   2,
				encoded: []byte{2},
			},
			{
				value:   -2,
				encoded: []byte{0x7e},
			},
			{
				value:   127,
				encoded: []byte{127 + 0x80, 0},
			},
			{
				value:   -127,
				encoded: []byte{1 + 0x80, 0x7f},
			},
			{
				value:   128,
				encoded: []byte{0 + 0x80, 1},
			},
			{
				value:   -128,
				encoded: []byte{0 + 0x80, 0x7f},
			},
			{
				value:   129,
				encoded: []byte{1 + 0x80, 1},
			},
			{
				value:   -129,
				encoded: []byte{0x7f + 0x80, 0x7e},
			},
		} {
			var b buf
			err := b.writeSLEB128(tc.value)
			require.NoError(t, err)
			require.Equal(t, tc.encoded, b.data)

			b.offset = 0

			actual, err := b.readSLEB128()
			require.NoError(t, err)
			require.Equal(t, tc.value, actual)
		}
	})

	t.Run("write: max byte count", func(t *testing.T) {

		t.Parallel()

		// This test ensures that only up to the maximum number of bytes are written
		// when writing a LEB128-encoded 32-bit number (see max32bitLEB128ByteCount),
		// i.e. test that only up to 5 bytes are written.

		var b buf
		err := b.writeSLEB128(math.MaxInt32)
		require.NoError(t, err)
		require.GreaterOrEqual(t, max32bitLEB128ByteCount, len(b.data))

		var b2 buf
		err = b2.writeSLEB128(math.MinInt32)
		require.NoError(t, err)
		require.GreaterOrEqual(t, max32bitLEB128ByteCount, len(b.data))
	})

	t.Run("read: max byte count", func(t *testing.T) {

		t.Parallel()

		// This test ensures that only up to the maximum number of bytes are read
		// when reading a LEB128-encoded 32-bit number (see max32bitLEB128ByteCount),
		// i.e. test that only 5 of the 8 given bytes are read,
		// to ensure the LEB128 parser doesn't keep reading infinitely.

		b := buf{data: []byte{0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88}}
		_, err := b.readSLEB128()
		require.NoError(t, err)
		require.Equal(t, offset(max32bitLEB128ByteCount), b.offset)
	})
}

func TestBuf_WriteSpaceAndSize(t *testing.T) {

	t.Parallel()

	var b buf

	err := b.WriteByte(101)
	require.NoError(t, err)
	err = b.WriteByte(102)
	require.NoError(t, err)

	off, err := b.writeFixedUint32LEB128Space()
	require.NoError(t, err)
	require.Equal(t, offset(2), off)
	require.Equal(t,
		[]byte{
			101, 102,
			0, 0, 0, 0, 0,
		},
		b.data,
	)

	err = b.WriteByte(104)
	require.NoError(t, err)
	err = b.WriteByte(105)
	require.NoError(t, err)
	err = b.WriteByte(106)
	require.NoError(t, err)

	err = b.writeUint32LEB128SizeAt(off)
	require.NoError(t, err)
	require.Equal(t,
		[]byte{
			101, 102,
			0x83, 0x80, 0x80, 0x80, 0,
			104, 105, 106,
		},
		b.data,
	)
}
