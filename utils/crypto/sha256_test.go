package crypto

import (
	"go-gen-tools/utils/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256(t *testing.T) {
	str := "songthing111"
	hashed := HashSHA256(str)
	t.Log(hashed)
	isHash := helper.IsSHA256(hashed)
	assert.True(t, isHash)
}
