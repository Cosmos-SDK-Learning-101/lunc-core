package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/core/x/treasury/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	m.keeper.SetBurnSplitRate(ctx, types.DefaultBurnTaxSplit)
	return nil
}
