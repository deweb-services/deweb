package keeper

//import (
//	"bytes"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	gogotypes "github.com/gogo/protobuf/types"
//	abci "github.com/tendermint/tendermint/abci/types"
//)
//
//// ApplyAndReturnValidatorSetUpdates applies and return accumulated updates to the bonded validator set. Also,
//// * Updates the active valset as keyed by LastValidatorPowerKey.
//// * Updates the total power as keyed by LastTotalPowerKey.
//// * Updates validator status' according to updated powers.
//// * Updates the fee pool bonded vs not-bonded tokens.
//// * Updates relevant indices.
//// It gets called once after genesis, another time maybe after genesis transactions,
//// then once at every EndBlock.
////
//// CONTRACT: Only validators with non-zero power or zero-power that were bonded
//// at the previous block height or were removed from the validator set entirely
//// are returned to Tendermint.
//func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate, err error) {
//	params := k.GetParams(ctx)
//	maxValidators := params.MaxValidators
//	powerReduction := k.PowerReduction(ctx)
//	totalPower := sdk.ZeroInt()
//	amtFromBondedToNotBonded, amtFromNotBondedToBonded := sdk.ZeroInt(), sdk.ZeroInt()
//
//	// Retrieve the last validator set.
//	// The persistent set is updated later in this function.
//	// (see LastValidatorPowerKey).
//	last, err := k.getLastValidatorsByAddr(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	// Iterate over validators, highest power to lowest.
//	iterator := k.ValidatorsPowerStoreIterator(ctx)
//	defer iterator.Close()
//
//	for count := 0; iterator.Valid() && count < int(maxValidators); iterator.Next() {
//		// everything that is iterated in this loop is becoming or already a
//		// part of the bonded validator set
//		valAddr := sdk.ValAddress(iterator.Value())
//		validator := k.mustGetValidator(ctx, valAddr)
//
//		if validator.Jailed {
//			panic("should never retrieve a jailed validator from the power store")
//		}
//
//		// if we get to a zero-power validator (which we don't bond),
//		// there are no more possible bonded validators
//		if validator.PotentialConsensusPower(k.PowerReduction(ctx)) == 0 {
//			break
//		}
//
//		// apply the appropriate state change if necessary
//		switch {
//		case validator.IsUnbonded():
//			validator, err = k.unbondedToBonded(ctx, validator)
//			if err != nil {
//				return
//			}
//			amtFromNotBondedToBonded = amtFromNotBondedToBonded.Add(validator.GetTokens())
//		case validator.IsUnbonding():
//			validator, err = k.unbondingToBonded(ctx, validator)
//			if err != nil {
//				return
//			}
//			amtFromNotBondedToBonded = amtFromNotBondedToBonded.Add(validator.GetTokens())
//		case validator.IsBonded():
//			// no state change
//		default:
//			panic("unexpected validator status")
//		}
//
//		// fetch the old power bytes
//		valAddrStr, err := sdk.Bech32ifyAddressBytes(sdk.GetConfig().GetBech32ValidatorAddrPrefix(), valAddr)
//		if err != nil {
//			return nil, err
//		}
//		oldPowerBytes, found := last[valAddrStr]
//		newPower := validator.ConsensusPower(powerReduction)
//		newPowerBytes := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: newPower})
//
//		// update the validator set if power has changed
//		if !found || !bytes.Equal(oldPowerBytes, newPowerBytes) {
//			updates = append(updates, validator.ABCIValidatorUpdate(powerReduction))
//
//			k.SetLastValidatorPower(ctx, valAddr, newPower)
//		}
//
//		delete(last, valAddrStr)
//		count++
//
//		totalPower = totalPower.Add(sdk.NewInt(newPower))
//	}
//
//	noLongerBonded, err := sortNoLongerBonded(last)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, valAddrBytes := range noLongerBonded {
//		validator := k.mustGetValidator(ctx, sdk.ValAddress(valAddrBytes))
//		validator, err = k.bondedToUnbonding(ctx, validator)
//		if err != nil {
//			return
//		}
//		amtFromBondedToNotBonded = amtFromBondedToNotBonded.Add(validator.GetTokens())
//		k.DeleteLastValidatorPower(ctx, validator.GetOperator())
//		updates = append(updates, validator.ABCIValidatorUpdateZero())
//	}
//
//	// Update the pools based on the recent updates in the validator set:
//	// - The tokens from the non-bonded candidates that enter the new validator set need to be transferred
//	// to the Bonded pool.
//	// - The tokens from the bonded validators that are being kicked out from the validator set
//	// need to be transferred to the NotBonded pool.
//	switch {
//	// Compare and subtract the respective amounts to only perform one transfer.
//	// This is done in order to avoid doing multiple updates inside each iterator/loop.
//	case amtFromNotBondedToBonded.GT(amtFromBondedToNotBonded):
//		k.notBondedTokensToBonded(ctx, amtFromNotBondedToBonded.Sub(amtFromBondedToNotBonded))
//	case amtFromNotBondedToBonded.LT(amtFromBondedToNotBonded):
//		k.bondedTokensToNotBonded(ctx, amtFromBondedToNotBonded.Sub(amtFromNotBondedToBonded))
//	default: // equal amounts of tokens; no update required
//	}
//
//	// set total power on lookup index if there are any updates
//	if len(updates) > 0 {
//		k.SetLastTotalPower(ctx, totalPower)
//	}
//
//	return updates, err
//}
