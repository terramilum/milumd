package keeper_test

import (
	"fmt"

	"github.com/terramirum/mirumd/x/rental/types"
)

var contractAddress = "trm1fq5xzwrduvzqeccgjraakk9sql87uttdyr78e7"

var request = &types.MsgDeployNftRequest{
	ContractOwner: "trm1fq5xzwrduvzqeccgjraakk9sql87uttdyr78e7",
	Name:          "Rent House 1",
	Symbol:        "RHT1",
	Description:   "Rent wonderful house and get a break for a while",
	Uri:           "https://testrent.io/detail",
}

func (s *TestSuite) TestDeployContract() {
	require := s.Require()
	res, err := s.rentKeeper.DeployNft(s.ctx, request)
	require.NoError(err)
	fmt.Println(res)
}

func (s *TestSuite) TestGettingClasses_ContractAddress() {
	require := s.Require()
	res, err := s.rentKeeper.DeployNft(s.ctx, request)
	fmt.Println(res)
	require.NoError(err)
	res, err = s.rentKeeper.DeployNft(s.ctx, request)
	fmt.Println(res)
	require.NoError(err)
	req := &types.QueryClassRequest{
		ContractOwner: contractAddress,
	}
	l, err := s.rentKeeper.Classes(s.ctx, req)

	require.Equal(2, len(l.NftClass))
}
