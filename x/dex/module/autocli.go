package dex

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "cosmuni/api/cosmuni/dex"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "LiquidityPoolAll",
					Use:       "list-liquidity-pool",
					Short:     "List all liquidityPool",
				},
				{
					RpcMethod:      "LiquidityPool",
					Use:            "show-liquidity-pool [id]",
					Short:          "Shows a liquidityPool",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePool",
					Use:            "create-pool [token-0] [token-1] [amount-0] [amount-1]",
					Short:          "Send a createPool tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "token0"}, {ProtoField: "token1"}, {ProtoField: "amount0"}, {ProtoField: "amount1"}},
				},
				{
					RpcMethod:      "Deposit",
					Use:            "deposit [token-0] [token-1] [amount-0] [amount-1]",
					Short:          "Send a deposit tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "token0"}, {ProtoField: "token1"}, {ProtoField: "amount0"}, {ProtoField: "amount1"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
