/*
Package ethrpc implements RPC methods to interact with the ethreum node geth.
*/
package ethrpc

import (
	"errors"

	jsonrpc "github.com/KeisukeYamashita/go-jsonrpc"
)

/*
RPCClient ...
*/
type RPCClient struct {
	*jsonrpc.RPCClient
}

/*
RPCer ...
Interface for other jsonrpc.
*/
type RPCer interface {
	GetBlockNumber() (string, error)
}

/*
NewRPCClient ...
NewRPCClient creates JSONRPC clients for your bitcoin node.
*/
func NewRPCClient(endpoint string) *RPCClient {
	c := new(RPCClient)
	c.RPCClient = jsonrpc.NewRPCClient(endpoint)
	return c
}

/*
GetBlockNumber ...
GetBlockNumber gets the most resent block height
*/
func (c *RPCClient) GetBlockNumber() (string, error) {
	resp, err := c.RPCClient.Call("eth_blockNumber")
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	var heightHex string
	resp.GetObject(&heightHex)
	return heightHex, nil
}

/*
GetBalance ...
GetBalance gets the balance of eth with the given address
*/
func (c *RPCClient) GetBalance(addr string) (string, error) {
	resp, err := c.RPCClient.Call("eth_getBalance", addr)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	var balance string
	resp.GetObject(&balance)
	return balance, nil
}
