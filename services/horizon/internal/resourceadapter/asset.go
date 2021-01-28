package resourceadapter

import (
	"context"

	protocol "github.com/AnneNamuli/go-stellar/protocols/horizon"
	"github.com/AnneNamuli/go-stellar/xdr"
)

func PopulateAsset(ctx context.Context, dest *protocol.Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
