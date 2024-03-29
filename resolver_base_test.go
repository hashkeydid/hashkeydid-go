package hashkeydid_go

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestCore_GetAddr(t *testing.T) {
	core := initTestCore()
	data, err := core.GetAddr(nil, big.NewInt(15921), 1)
	assert.Nil(t, err)
	assert.Equal(t, common.Hex2Bytes("617E266FFA5c2B168fB6B6cE1Bee9CA2E461DD58"), data)
}

func TestCore_GetText(t *testing.T) {
	core := initTestCore()
	name, err := core.GetText(nil, big.NewInt(15921), "name")
	assert.Nil(t, err)
	assert.Equal(t, "gosdktest", name)
}

func TestCore_GetDIDNameByAddr(t *testing.T) {
	//core := initTestCore()
	//// the did is claimed at 40069812
	//opts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(40069811)}
	//_, err := core.GetDIDNameByAddr(opts, common.HexToAddress("0x617E266FFA5c2B168fB6B6cE1Bee9CA2E461DD58"))
	//assert.Equal(t, ErrAddrNotClaimed, err)
	//opts.BlockNumber = new(big.Int).SetUint64(40069812)
	//_, err = core.GetDIDNameByAddr(opts, common.HexToAddress("0x617E266FFA5c2B168fB6B6cE1Bee9CA2E461DD58"))
	//assert.Equal(t, ErrAddrNotSetReverse, err)
	//// the did is set reverse at 40070952
	//opts.BlockNumber = new(big.Int).SetUint64(40070952)
	//name, err := core.GetDIDNameByAddr(opts, common.HexToAddress("0x617E266FFA5c2B168fB6B6cE1Bee9CA2E461DD58"))
	//assert.Nil(t, err)
	//assert.Equal(t, "gosdktest.key", name)
}
