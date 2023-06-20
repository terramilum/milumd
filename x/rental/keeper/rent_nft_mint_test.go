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

func (s *TestSuite) TestRentMintNft_DefineQuery() {
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

	firstStartDate := getNowUtcAddMin(10)
	firstEndDate := getNowUtcAddMin(25)
	request := &types.MsgMintRentRequest{
		ContractOwner: nftOwner,
		ClassId:       deplotNftResponse.ClassId,
		NftId:         mintNftResponse.NftId,
		Renter:        renter,
		StartDate:     firstStartDate,
		EndDate:       firstEndDate,
	}

	_, err = s.rentKeeper.RentNftMint(s.ctx, request)
	require.NoError(err)
	request.StartDate = getNowUtcAddMin(12)
	request.EndDate = getNowUtcAddMin(24)
	_, err = s.rentKeeper.RentNftMint(s.ctx, request)
	require.Error(err, types.ErrNftOwnerCanRent)

	secondStartDate := getNowUtcAddMin(30)
	secondEndDate := getNowUtcAddMin(45)

	request.StartDate = secondStartDate
	request.EndDate = secondEndDate
	_, err = s.rentKeeper.RentNftMint(s.ctx, request)
	require.NoError(err)

	req := &types.QuerySessionRequest{
		ClassId: mintNftRequest.ClassId,
		NftId:   mintNftRequest.NftId,
	}
	res, err := s.rentKeeper.Sessions(s.ctx, req)
	require.NoError(err)
	require.Equal(2, len(res.NftRent))
	require.Equal(firstStartDate, res.NftRent[0].StartDate)
	require.Equal(firstEndDate, res.NftRent[0].EndDate)
	require.Equal(secondStartDate, res.NftRent[1].StartDate)
	require.Equal(secondEndDate, res.NftRent[1].EndDate)

	thirdStartDate := getNowUtcAddMin(60)
	thirdEndDate := getNowUtcAddMin(65)

	secondRenter := "trm17ulg0elrs4v32962awhaxcwh5qv7sv7vwmeusa"
	request.Renter = secondRenter
	request.StartDate = thirdStartDate
	request.EndDate = thirdEndDate
	_, err = s.rentKeeper.RentNftMint(s.ctx, request)
	require.NoError(err)

	res, err = s.rentKeeper.Sessions(s.ctx, req)
	require.NoError(err)
	require.Equal(3, len(res.NftRent))

	req.Renter = renter
	res, err = s.rentKeeper.Sessions(s.ctx, req)
	require.NoError(err)
	require.Equal(2, len(res.NftRent))

	req.Renter = secondRenter
	res, err = s.rentKeeper.Sessions(s.ctx, req)
	require.NoError(err)
	require.Equal(1, len(res.NftRent))

	accessNftRequest := &types.MsgAccessNftRequest{
		Renter:  renter,
		ClassId: request.ClassId,
		NftId:   request.NftId,
	}

	s.rentKeeper.SetAdditionalMinutesToCurrentDate(1)
	accessRes, err := s.rentKeeper.NftAccess(s.ctx, accessNftRequest)
	require.Equal(false, accessRes.HasAccess)

	s.rentKeeper.SetAdditionalMinutesToCurrentDate(15)
	accessRes, err = s.rentKeeper.NftAccess(s.ctx, accessNftRequest)
	require.Equal(true, accessRes.HasAccess)

	accessNftRequest.Renter = secondRenter
	accessRes, err = s.rentKeeper.NftAccess(s.ctx, accessNftRequest)
	require.Equal(false, accessRes.HasAccess)
}

func getNowUtcAddMin(addMin int32) int64 {
	now := time.Now().Add(time.Minute * time.Duration(addMin)).UTC()
	formatted := now.Format("200601021504")
	d, _ := strconv.ParseInt(formatted, 10, 64)
	return d
}
