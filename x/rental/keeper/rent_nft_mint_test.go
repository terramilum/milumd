package keeper_test

import (
	"fmt"
	"strconv"
	"time"

	"github.com/terramirum/mirumd/x/rental/types"
)

func (s *TestSuite) TestRentMintNft_DefineSession() {
	require := s.Require()
	nftOwner := "trm1z60lkcwptx4yuykdp9fqcmr56qgeqdh8h4zz3g"
	renter := "trm1wyyz7tw8m0z4n6elkc6fx4rwtm2jsll72eh452"

	deplotNftResponse, err := s.rentKeeper.DeployNft(s.ctx, deplotNftRequest)
	require.NoError(err)

	mintNftRequest := &types.MsgMintNftRequest{
		ContractOwner: deplotNftRequest.ContractOwner,
		Reciever:      nftOwner,
		ClassId:       deplotNftResponse.ClassId,
		NftId:         "1002",
		Detail:        &types.Detail{},
	}

	mintNftResponse, err := s.rentKeeper.MintNft(s.ctx, mintNftRequest)
	require.NoError(err)

	request := &types.MsgMintRentRequest{
		ContractOwner: nftOwner,
		ClassId:       deplotNftResponse.ClassId,
		NftId:         mintNftResponse.NftId,
		Renter:        renter,
		StartDate:     getNowUtcAddMin(10),
		EndDate:       getNowUtcAddMin(25),
	}

	res, err := s.rentKeeper.RentNftMint(s.ctx, request)
	fmt.Println(res)
}

func getNowUtcAddMin(addMin int32) int64 {
	now := time.Now().Add(time.Minute * time.Duration(addMin)).UTC()
	formatted := now.Format("200601021504")
	d, _ := strconv.ParseInt(formatted, 10, 64)
	return d
}
