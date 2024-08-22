package keeper_test

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/suite"

	"github.com/golang/mock/gomock"
	rentkeeper "github.com/terramirum/mirumd/x/rental/keeper"
	"github.com/terramirum/mirumd/x/rental/types"
	renttypes "github.com/terramirum/mirumd/x/rental/types"

	storetypes "cosmossdk.io/store/types"
	nft "cosmossdk.io/x/nft"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	nftmodule "cosmossdk.io/x/nft/module"
	nfttestutil "cosmossdk.io/x/nft/testutil"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/testutil"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
)

type TestSuite struct {
	suite.Suite

	ctx         sdk.Context
	addrs       []sdk.AccAddress
	queryClient nft.QueryClient
	nftKeeper   nftkeeper.Keeper
	rentKeeper  *rentkeeper.Keeper

	encCfg moduletestutil.TestEncodingConfig
}

func TestTestSuite(t *testing.T) {
	SetPrefixes("mirum")
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupTest() {
	// suite setup
	s.addrs = simtestutil.CreateIncrementalAccounts(3)
	s.encCfg = moduletestutil.MakeTestEncodingConfig(nftmodule.AppModuleBasic{})

	keyRental := storetypes.NewKVStoreKey(renttypes.StoreKey)
	storeServiceRental := runtime.NewKVStoreService(keyRental)
	testCtx := testutil.DefaultContextWithDB(s.T(), keyRental, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Height: 1})

	// gomock initializations
	ctrl := gomock.NewController(s.T())
	accountKeeper := nfttestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := nfttestutil.NewMockBankKeeper(ctrl)
	accountKeeper.EXPECT().GetModuleAddress("nft").Return(s.addrs[0]).AnyTimes()
	keyNft := storetypes.NewKVStoreKey(nftkeeper.StoreKey)
	storeServiceNft := runtime.NewKVStoreService(keyNft)
	memKeys := storetypes.NewMemoryStoreKeys(renttypes.MemStoreKey)

	nftKeeper := nftkeeper.NewKeeper(
		storeServiceNft,
		s.encCfg.Codec,
		accountKeeper,
		bankKeeper,
	)

	rentKeeper := rentkeeper.NewKeeper(s.encCfg.Codec, storeServiceRental, memKeys[renttypes.MemStoreKey], &nftKeeper)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, s.encCfg.InterfaceRegistry)
	nft.RegisterQueryServer(queryHelper, nftKeeper)

	s.nftKeeper = nftKeeper
	s.rentKeeper = rentKeeper
	s.queryClient = nft.NewQueryClient(queryHelper)
	s.ctx = ctx
}

func SetPrefixes(accountAddressPrefix string) {
	// Set prefixes
	accountPubKeyPrefix := accountAddressPrefix + "pub"
	validatorAddressPrefix := accountAddressPrefix + "valoper"
	validatorPubKeyPrefix := accountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := accountAddressPrefix + "valcons"
	consNodePubKeyPrefix := accountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(accountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()

	ir := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(ir)
	codec.RegisterInterfaces(ir)
	types.RegisterInterfaces(ir)
}
