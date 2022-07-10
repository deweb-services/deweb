package testutil

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	nftcli "github.com/deweb-services/deweb/x/nftmodule/client/cli"
)

func BurnNFTExec(clientCtx client.Context, from string, denomID string, tokenID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		tokenID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, nftcli.GetCmdRemoveDomain(), args)
}

func MintNFTExec(clientCtx client.Context, from string, denomID string, tokenID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		tokenID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, nftcli.GetCmdMintNFT(), args)
}

func EditNFTExec(clientCtx client.Context, from string, denomID string, tokenID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		tokenID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, nftcli.GetCmdEditNFT(), args)
}

func TransferNFTExec(clientCtx client.Context, from string, recipient string, denomID string, tokenID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		recipient,
		denomID,
		tokenID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, nftcli.GetCmdTransferNFT(), args)
}

func QueryDenomExec(clientCtx client.Context, denomID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, nftcli.GetCmdQueryDenom(), args)
}

func QueryCollectionExec(clientCtx client.Context, denomID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, nftcli.GetCmdQueryCollection(), args)
}

func QueryDenomsExec(clientCtx client.Context, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, nftcli.GetCmdQueryDenoms(), args)
}

func QuerySupplyExec(clientCtx client.Context, denom string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denom,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, nftcli.GetCmdQuerySupply(), args)
}

func QueryOwnerExec(clientCtx client.Context, address string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		address,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, nftcli.GetCmdQueryOwner(), args)
}

func QueryNFTExec(clientCtx client.Context, denomID string, tokenID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		tokenID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, nftcli.GetCmdQueryDomain(), args)
}
