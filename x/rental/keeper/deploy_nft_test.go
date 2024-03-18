package keeper_test

import (
	"fmt"

	"github.com/terramirum/mirumd/x/rental/types"
)

var contractAddress = "mirum1fq5xzwrduvzqeccgjraakk9sql87uttdyr78e7"

var deplotNftRequest = &types.MsgDeployNftRequest{
	ContractOwner: "mirum1fq5xzwrduvzqeccgjraakk9sql87uttdyr78e7",
	Name:          "Rent House 1",
	Symbol:        "RHT1",
	Description:   "Rent wonderful house and get a break for a while",
	Uri:           "https://testrent.io/detail",
	Detail:        &types.Detail{},
}

func (s *TestSuite) TestDeployContract() {
	require := s.Require()
	res, err := s.rentKeeper.DeployNft(s.ctx, deplotNftRequest)
	require.NoError(err)
	fmt.Println(res)
}

func (s *TestSuite) TestGettingClasses_ContractAddress() {
	require := s.Require()
	res, err := s.rentKeeper.DeployNft(s.ctx, deplotNftRequest)
	fmt.Println(res)
	require.NoError(err)
	res, err = s.rentKeeper.DeployNft(s.ctx, deplotNftRequest)
	fmt.Println(res)
	require.NoError(err)
	req := &types.QueryClassRequest{
		ContractOwner: contractAddress,
	}
	l, err := s.rentKeeper.Classes(s.ctx, req)

	require.Equal(2, len(l.NftClass))
}
