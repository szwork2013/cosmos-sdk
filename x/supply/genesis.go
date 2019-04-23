package supply

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply/keeper"
	"github.com/cosmos/cosmos-sdk/x/supply/types"
)

// GenesisState is the supply state that must be provided at genesis.
type GenesisState struct {
	Supplier types.Supplier `json:"supplier"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(supplier types.Supplier) GenesisState {
	return GenesisState{Supplier: supplier}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState(types.DefaultSupplier())
}

// InitGenesis sets supply information for genesis.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data GenesisState) {
	k.SetSupplier(ctx, data.Supplier)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) GenesisState {
	return NewGenesisState(k.GetSupplier(ctx))
}

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	if err := data.Supplier.ValidateBasic(); err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}