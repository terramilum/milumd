package keeper_test

import (
	"fmt"

	"github.com/terramirum/mirumd/x/rental/types"
)

func (s *TestSuite) TestDeployContract() {
	require := s.Require()
	request := &types.MsgDeployNftRequest{
		ContractOwner: "trm1fq5xzwrduvzqeccgjraakk9sql87uttdyr78e7",
		Name:          "Rent House 1",
		Symbol:        "RHT1",
		Description:   "Rent wonderful house and get a break for a while",
		Uri:           "https://testrent.io/detail",
	}
	res, err := s.rentKeeper.DeployNft(s.ctx, request)
	require.NoError(err)
	fmt.Println(res)
}
