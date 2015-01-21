package bit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetFlagByPos(t *testing.T) {
	var bitInt BitInt
	var pos uint64 = 0

	pos = 1
	assert.True(t, bitInt.GetFlag(pos) == 0)
	bitInt.SetFlag(pos)
	assert.True(t, bitInt == 2) // 2#10
	assert.True(t, bitInt.GetFlag(pos) == 1)

	var pos5 uint64 = 5
	assert.True(t, bitInt.GetFlag(pos5) == 0)
	bitInt.SetFlag(pos5)
	assert.True(t, bitInt == 34) // 2#100010
	assert.True(t, bitInt.GetFlag(pos5) == 1)

	bitInt.UnSetFlag(pos5)
	assert.True(t, bitInt == 2) // 2#10
	assert.True(t, bitInt.GetFlag(pos5) == 0)
	assert.False(t, bitInt.IsAllZero())

}

func TestBitIntGetTrueLen(t *testing.T) {
	var bitInt BitInt = 0
	assert.True(t, 0 == bitInt.GetTrueLen())
	bitInt = 1 // 2#1
	assert.True(t, 1 == bitInt.GetTrueLen())
	bitInt = 3 // 2#11
	assert.True(t, 2 == bitInt.GetTrueLen())
	assert.True(t, 1 == bitInt.GetTrueLen(1))

	bitInt = 5 // 2#101
	assert.True(t, 2 == bitInt.GetTrueLen())
	assert.True(t, 1 == bitInt.GetTrueLen(2))

}
