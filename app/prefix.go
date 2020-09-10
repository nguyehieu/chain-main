package app

import (
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	CoinType       = 394
	FundraiserPath = "44'/394'/0'/0/1"
)

var (
	AccountAddressPrefix   = "cro"
	AccountPubKeyPrefix    = "cropub"
	ValidatorAddressPrefix = "crocncl"
	ValidatorPubKeyPrefix  = "crocnclpub"
	ConsNodeAddressPrefix  = "crocnclcons"
	ConsNodePubKeyPrefix   = "crocnclconspub"
	HumanCoinUnit          = "cro"
	BaseCoinUnit           = "basecro" // 10^-8 AKA "carson"
)

func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)

	config.SetCoinType(CoinType)
	config.SetFullFundraiserPath(FundraiserPath)

	croUnit := sdk.OneDec()
	err := sdk.RegisterDenom(HumanCoinUnit, croUnit)
	if err != nil {
		log.Fatal(err)
	}

	carsonUnit := sdk.NewDecWithPrec(1, 8) // 10^-8 (carson)
	err = sdk.RegisterDenom(BaseCoinUnit, carsonUnit)

	if err != nil {
		log.Fatal(err)
	}

	config.Seal()
}
