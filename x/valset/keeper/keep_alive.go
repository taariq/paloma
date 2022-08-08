package keeper

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keeperutil "github.com/palomachain/paloma/util/keeper"
	"github.com/vizualni/whoops"
)

const (
	// TODO: make this a param
	defaultKeepAliveDuration = 5 * time.Minute
)

type keepAliveData struct {
	ValAddr     sdk.ValAddress
	ContactedAt time.Time
	AliveUntil  time.Time
}

func (k Keeper) KeepValidatorAlive(ctx sdk.Context, valAddr sdk.ValAddress) error {
	if err := k.CanAcceptValidator(ctx, valAddr); err != nil {
		if !errors.Is(err, ErrValidatorNotInKeepAlive) {
			return err
		}
	}

	store := k.keepAliveStore(ctx)
	data := keepAliveData{
		ValAddr:     valAddr,
		ContactedAt: ctx.BlockTime(),
		AliveUntil:  ctx.BlockTime().Add(defaultKeepAliveDuration),
	}
	bz, err := json.Marshal(data)
	if err != nil {
		return err
	}
	store.Set(valAddr, bz)
	return nil
}

func (k Keeper) IsValidatorAlive(ctx sdk.Context, valAddr sdk.ValAddress) (bool, error) {
	aliveUntil, err := k.ValidatorAliveUntil(ctx, valAddr)
	if err != nil {
		return false, err
	}
	return ctx.BlockTime().Before(aliveUntil), nil
}

func (k Keeper) ValidatorAliveUntil(ctx sdk.Context, valAddr sdk.ValAddress) (time.Time, error) {
	store := k.keepAliveStore(ctx)
	if !store.Has(valAddr) {
		return time.Time{}, ErrValidatorNotInKeepAlive.Format(valAddr)
	}
	dataBz := store.Get(valAddr)
	var data keepAliveData
	err := json.Unmarshal(dataBz, &data)
	if err != nil {
		return time.Time{}, err
	}
	return data.AliveUntil, nil
}

func (k Keeper) CanAcceptValidator(ctx sdk.Context, valAddr sdk.ValAddress) error {
	stakingVal := k.staking.Validator(ctx, valAddr)

	if stakingVal == nil {
		return ErrValidatorWithAddrNotFound.Format(valAddr.String())
	}

	if stakingVal.IsJailed() {
		return ErrValidatorCannotBePigeon.Format(valAddr.String()).WrapS("validator is jailed")
	}

	if !stakingVal.IsBonded() {
		return ErrValidatorCannotBePigeon.Format(valAddr.String()).WrapS("validator is not bonded")
	}

	return nil
}

func (k Keeper) JailInactiveValidators(ctx sdk.Context) error {
	store := k.keepAliveStore(ctx)
	keys, _, _ := keeperutil.IterAllRaw(store, k.cdc)
	var g whoops.Group
	for _, bz := range keys {
		valAddr := sdk.ValAddress(bz)
		alive, err := k.IsValidatorAlive(ctx, valAddr)
		if err != nil {
			g.Add(err)
			continue
		}
		if alive {
			continue
		}
		store.Delete(valAddr)
		consAddr, err := k.staking.Validator(ctx, valAddr).GetConsAddr()
		if err != nil {
			g.Add(err)
			continue
		}
		k.staking.Jail(ctx, consAddr)

	}
	return g.Return()
}

func (k Keeper) keepAliveStore(ctx sdk.Context) sdk.KVStore {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte("keep-alive/"))
}