package keeper_test

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/suite"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/golang/mock/gomock"
	rentkeeper "github.com/terramirum/mirumd/x/rental/keeper"
	"github.com/terramirum/mirumd/x/rental/types"
	renttypes "github.com/terramirum/mirumd/x/rental/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/testutil"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	nft "github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	nftmodule "github.com/cosmos/cosmos-sdk/x/nft/module"
	nfttestutil "github.com/cosmos/cosmos-sdk/x/nft/testutil"
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

	key := sdk.NewKVStoreKey(renttypes.StoreKey)
	memKeys := sdk.NewMemoryStoreKeys("test")

	testCtx := testutil.DefaultContextWithDB(s.T(), key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	// gomock initializations
	ctrl := gomock.NewController(s.T())
	accountKeeper := nfttestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := nfttestutil.NewMockBankKeeper(ctrl)
	accountKeeper.EXPECT().GetModuleAddress("nft").Return(s.addrs[0]).AnyTimes()

	nftKeeper := nftkeeper.NewKeeper(key, s.encCfg.Codec, accountKeeper, bankKeeper)
	rentKeeper := rentkeeper.NewKeeper(s.encCfg.Codec, key, memKeys[renttypes.StoreKey], &nftKeeper)
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
