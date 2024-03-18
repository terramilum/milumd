package keeper_test

import (
	"strconv"
	"time"

	"github.com/terramirum/mirumd/x/rental/types"
)

func (s *TestSuite) TestRentMintNft_DefineSession() {
	require := s.Require()
	nftOwner := "mirum1z60lkcwptx4yuykdp9fqcmr56qgeqdh8h4zz3g"
	renter := "mirum1wyyz7tw8m0z4n6elkc6fx4rwtm2jsll72eh452"

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

	_, err = s.rentKeeper.RentNftMint(s.ctx, request)
	require.NoError(err)
}

func (s *TestSuite) TestRentMintNft_DefineQuery() {
	require := s.Require()
	nftOwner := "mirum1z60lkcwptx4yuykdp9fqcmr56qgeqdh8h4zz3g"
	renter := "mirum1wyyz7tw8m0z4n6elkc6fx4rwtm2jsll72eh452"
	rents := s.rentKeeper.GetAllSessionOfNft(s.ctx)
	require.Equal(0, len(rents))

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
	require.Equal(2, len(res.SessionDetail))
	require.Equal(firstStartDate, res.SessionDetail[0].NftRent.StartDate)
	require.Equal(firstEndDate, res.SessionDetail[0].NftRent.EndDate)
	require.Equal(secondStartDate, res.SessionDetail[1].NftRent.StartDate)
	require.Equal(secondEndDate, res.SessionDetail[1].NftRent.EndDate)

	thirdStartDate := getNowUtcAddMin(60)
	thirdEndDate := getNowUtcAddMin(65)

	secondRenter := "mirum17ulg0elrs4v32962awhaxcwh5qv7sv7vwmeusa"
	request.Renter = secondRenter
	request.StartDate = thirdStartDate
	request.EndDate = thirdEndDate
	_, err = s.rentKeeper.RentNftMint(s.ctx, request)
	require.NoError(err)

	res, err = s.rentKeeper.Sessions(s.ctx, req)
	require.NoError(err)
	require.Equal(3, len(res.SessionDetail))

	req.Renter = renter
	res, err = s.rentKeeper.Sessions(s.ctx, req)
	require.NoError(err)
	require.Equal(2, len(res.SessionDetail))

	req.Renter = secondRenter
	res, err = s.rentKeeper.Sessions(s.ctx, req)
	require.NoError(err)
	require.Equal(1, len(res.SessionDetail))

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

	renterReq := &types.QuerySessionRequest{
		Renter: renter,
	}
	res, err = s.rentKeeper.Sessions(s.ctx, renterReq)
	require.NoError(err)
	require.Equal(2, len(res.SessionDetail))

	sendTo := "mirumum16upgs0enps8mf6phjrj0pt8p2t8lp032ach8uh"
	transferedSessionId := res.SessionDetail[0].NftRent.SessionId
	sendRequest := &types.MsgSendSessionRequest{
		FromRenter: renter,
		ToRenter:   sendTo,
		ClassId:    request.ClassId,
		NftId:      request.NftId,
		SessionId:  transferedSessionId,
	}
	_, err = s.rentKeeper.SendSession(s.ctx, sendRequest)
	require.NoError(err)

	res, err = s.rentKeeper.Sessions(s.ctx, renterReq)
	require.NoError(err)
	require.Equal(1, len(res.SessionDetail))

	renterReq.SessionId = transferedSessionId
	res, err = s.rentKeeper.Sessions(s.ctx, renterReq)
	require.NoError(err)
	require.Equal(0, len(res.SessionDetail))

	renterReq = &types.QuerySessionRequest{
		Renter: sendTo,
	}
	res, err = s.rentKeeper.Sessions(s.ctx, renterReq)
	require.NoError(err)
	require.Equal(1, len(res.SessionDetail))
	require.Equal(res.SessionDetail[0].NftRent.SessionId, transferedSessionId)

	rents = s.rentKeeper.GetAllSessionOfNft(s.ctx)
	owners := s.rentKeeper.GetAllClassIdsOwners(s.ctx)
	require.Equal(len(rents), 3)
	require.Equal(len(owners), 2)
}

func getNowUtcAddMin(addMin int32) int64 {
	now := time.Now().Add(time.Minute * time.Duration(addMin)).UTC()
	formatted := now.Format("200601021504")
	d, _ := strconv.ParseInt(formatted, 10, 64)
	return d
}
