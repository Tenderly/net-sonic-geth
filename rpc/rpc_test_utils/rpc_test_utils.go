package rpc_test_utils

import (
	"github.com/tenderly/net-sonic-geth/accounts"
	"github.com/tenderly/net-sonic-geth/eth/tracers"
	"github.com/tenderly/net-sonic-geth/internal/ethapi"
	"github.com/tenderly/net-sonic-geth/params"
	"github.com/tenderly/net-sonic-geth/rpc"
)

// GetRpcApis returns a list of RPC APIs for testing
func GetRpcApis() []rpc.API {
	backend := &dummyBackend{}
	apis := ethapi.GetAPIs(backend)
	apis = append(apis, tracers.APIs(nil)...)
	return apis
}

type dummyBackend struct {
	ethapi.Backend
}

func (b *dummyBackend) ChainConfig() *params.ChainConfig {
	return &params.ChainConfig{}
}

func (b *dummyBackend) AccountManager() *accounts.Manager {
	return nil
}
