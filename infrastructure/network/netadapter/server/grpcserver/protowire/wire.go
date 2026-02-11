package protowire

import (
	"github.com/karlsen-network/karlsend/v2/app/appmessage"
	"github.com/pkg/errors"
)

type converter interface {
	toAppMessage() (appmessage.Message, error)
}

// ToAppMessage converts a KarlsendMessage to its appmessage.Message representation
func (x *KarlsendMessage) ToAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KarlsendMessage is nil")
	}
	converter, ok := x.Payload.(converter)
	if !ok {
		return nil, errors.Errorf("received invalid message")
	}
	appMessage, err := converter.toAppMessage()
	if err != nil {
		return nil, err
	}
	return appMessage, nil
}

// FromAppMessage creates a KarlsendMessage from a appmessage.Message
func FromAppMessage(message appmessage.Message) (*KarlsendMessage, error) {
	payload, err := toPayload(message)
	if err != nil {
		return nil, err
	}
	return &KarlsendMessage{
		Payload: payload,
	}, nil
}

func toPayload(message appmessage.Message) (isKarlsendMessage_Payload, error) {
	p2pPayload, err := toP2PPayload(message)
	if err != nil {
		return nil, err
	}
	if p2pPayload != nil {
		return p2pPayload, nil
	}

	rpcPayload, err := toRPCPayload(message)
	if err != nil {
		return nil, err
	}
	if rpcPayload != nil {
		return rpcPayload, nil
	}

	return nil, errors.Errorf("unknown message type %T", message)
}

func toP2PPayload(message appmessage.Message) (isKarlsendMessage_Payload, error) {
	switch message := message.(type) {
	case *appmessage.MsgAddresses:
		payload := new(KarlsendMessage_Addresses)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlock:
		payload := new(KarlsendMessage_Block)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestBlockLocator:
		payload := new(KarlsendMessage_RequestBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockLocator:
		payload := new(KarlsendMessage_BlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestAddresses:
		payload := new(KarlsendMessage_RequestAddresses)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestIBDBlocks:
		payload := new(KarlsendMessage_RequestIBDBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextHeaders:
		payload := new(KarlsendMessage_RequestNextHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDoneHeaders:
		payload := new(KarlsendMessage_DoneHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestRelayBlocks:
		payload := new(KarlsendMessage_RequestRelayBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestTransactions:
		payload := new(KarlsendMessage_RequestTransactions)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTransactionNotFound:
		payload := new(KarlsendMessage_TransactionNotFound)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgInvRelayBlock:
		payload := new(KarlsendMessage_InvRelayBlock)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgInvTransaction:
		payload := new(KarlsendMessage_InvTransactions)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPing:
		payload := new(KarlsendMessage_Ping)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPong:
		payload := new(KarlsendMessage_Pong)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTx:
		payload := new(KarlsendMessage_Transaction)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgVerAck:
		payload := new(KarlsendMessage_Verack)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgVersion:
		payload := new(KarlsendMessage_Version)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgReject:
		payload := new(KarlsendMessage_Reject)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointUTXOSet:
		payload := new(KarlsendMessage_RequestPruningPointUTXOSet)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPointUTXOSetChunk:
		payload := new(KarlsendMessage_PruningPointUtxoSetChunk)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgUnexpectedPruningPoint:
		payload := new(KarlsendMessage_UnexpectedPruningPoint)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocator:
		payload := new(KarlsendMessage_IbdBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocatorHighestHash:
		payload := new(KarlsendMessage_IbdBlockLocatorHighestHash)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocatorHighestHashNotFound:
		payload := new(KarlsendMessage_IbdBlockLocatorHighestHashNotFound)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BlockHeadersMessage:
		payload := new(KarlsendMessage_BlockHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextPruningPointUTXOSetChunk:
		payload := new(KarlsendMessage_RequestNextPruningPointUtxoSetChunk)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDonePruningPointUTXOSetChunks:
		payload := new(KarlsendMessage_DonePruningPointUtxoSetChunks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockWithTrustedData:
		payload := new(KarlsendMessage_BlockWithTrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointAndItsAnticone:
		payload := new(KarlsendMessage_RequestPruningPointAndItsAnticone)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDoneBlocksWithTrustedData:
		payload := new(KarlsendMessage_DoneBlocksWithTrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlock:
		payload := new(KarlsendMessage_IbdBlock)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestHeaders:
		payload := new(KarlsendMessage_RequestHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPoints:
		payload := new(KarlsendMessage_PruningPoints)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointProof:
		payload := new(KarlsendMessage_RequestPruningPointProof)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPointProof:
		payload := new(KarlsendMessage_PruningPointProof)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgReady:
		payload := new(KarlsendMessage_Ready)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTrustedData:
		payload := new(KarlsendMessage_TrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockWithTrustedDataV4:
		payload := new(KarlsendMessage_BlockWithTrustedDataV4)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks:
		payload := new(KarlsendMessage_RequestNextPruningPointAndItsAnticoneBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestIBDChainBlockLocator:
		payload := new(KarlsendMessage_RequestIBDChainBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDChainBlockLocator:
		payload := new(KarlsendMessage_IbdChainBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestAnticone:
		payload := new(KarlsendMessage_RequestAnticone)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	default:
		return nil, nil
	}
}

func toRPCPayload(message appmessage.Message) (isKarlsendMessage_Payload, error) {
	switch message := message.(type) {
	case *appmessage.GetCurrentNetworkRequestMessage:
		payload := new(KarlsendMessage_GetCurrentNetworkRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCurrentNetworkResponseMessage:
		payload := new(KarlsendMessage_GetCurrentNetworkResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockRequestMessage:
		payload := new(KarlsendMessage_SubmitBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockResponseMessage:
		payload := new(KarlsendMessage_SubmitBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateRequestMessage:
		payload := new(KarlsendMessage_GetBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateResponseMessage:
		payload := new(KarlsendMessage_GetBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyBlockAddedRequestMessage:
		payload := new(KarlsendMessage_NotifyBlockAddedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyBlockAddedResponseMessage:
		payload := new(KarlsendMessage_NotifyBlockAddedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BlockAddedNotificationMessage:
		payload := new(KarlsendMessage_BlockAddedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetPeerAddressesRequestMessage:
		payload := new(KarlsendMessage_GetPeerAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetPeerAddressesResponseMessage:
		payload := new(KarlsendMessage_GetPeerAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSelectedTipHashRequestMessage:
		payload := new(KarlsendMessage_GetSelectedTipHashRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSelectedTipHashResponseMessage:
		payload := new(KarlsendMessage_GetSelectedTipHashResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntryRequestMessage:
		payload := new(KarlsendMessage_GetMempoolEntryRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntryResponseMessage:
		payload := new(KarlsendMessage_GetMempoolEntryResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetConnectedPeerInfoRequestMessage:
		payload := new(KarlsendMessage_GetConnectedPeerInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetConnectedPeerInfoResponseMessage:
		payload := new(KarlsendMessage_GetConnectedPeerInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.AddPeerRequestMessage:
		payload := new(KarlsendMessage_AddPeerRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.AddPeerResponseMessage:
		payload := new(KarlsendMessage_AddPeerResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionRequestMessage:
		payload := new(KarlsendMessage_SubmitTransactionRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionResponseMessage:
		payload := new(KarlsendMessage_SubmitTransactionResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentChainChangedRequestMessage:
		payload := new(KarlsendMessage_NotifyVirtualSelectedParentChainChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentChainChangedResponseMessage:
		payload := new(KarlsendMessage_NotifyVirtualSelectedParentChainChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualSelectedParentChainChangedNotificationMessage:
		payload := new(KarlsendMessage_VirtualSelectedParentChainChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockRequestMessage:
		payload := new(KarlsendMessage_GetBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockResponseMessage:
		payload := new(KarlsendMessage_GetBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSubnetworkRequestMessage:
		payload := new(KarlsendMessage_GetSubnetworkRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSubnetworkResponseMessage:
		payload := new(KarlsendMessage_GetSubnetworkResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentChainFromBlockRequestMessage:
		payload := new(KarlsendMessage_GetVirtualSelectedParentChainFromBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentChainFromBlockResponseMessage:
		payload := new(KarlsendMessage_GetVirtualSelectedParentChainFromBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlocksRequestMessage:
		payload := new(KarlsendMessage_GetBlocksRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlocksResponseMessage:
		payload := new(KarlsendMessage_GetBlocksResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockCountRequestMessage:
		payload := new(KarlsendMessage_GetBlockCountRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockCountResponseMessage:
		payload := new(KarlsendMessage_GetBlockCountResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoRequestMessage:
		payload := new(KarlsendMessage_GetBlockDagInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoResponseMessage:
		payload := new(KarlsendMessage_GetBlockDagInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ResolveFinalityConflictRequestMessage:
		payload := new(KarlsendMessage_ResolveFinalityConflictRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ResolveFinalityConflictResponseMessage:
		payload := new(KarlsendMessage_ResolveFinalityConflictResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyFinalityConflictsRequestMessage:
		payload := new(KarlsendMessage_NotifyFinalityConflictsRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyFinalityConflictsResponseMessage:
		payload := new(KarlsendMessage_NotifyFinalityConflictsResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.FinalityConflictNotificationMessage:
		payload := new(KarlsendMessage_FinalityConflictNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.FinalityConflictResolvedNotificationMessage:
		payload := new(KarlsendMessage_FinalityConflictResolvedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesRequestMessage:
		payload := new(KarlsendMessage_GetMempoolEntriesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesResponseMessage:
		payload := new(KarlsendMessage_GetMempoolEntriesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ShutDownRequestMessage:
		payload := new(KarlsendMessage_ShutDownRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ShutDownResponseMessage:
		payload := new(KarlsendMessage_ShutDownResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetHeadersRequestMessage:
		payload := new(KarlsendMessage_GetHeadersRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetHeadersResponseMessage:
		payload := new(KarlsendMessage_GetHeadersResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyUTXOsChangedRequestMessage:
		payload := new(KarlsendMessage_NotifyUtxosChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyUTXOsChangedResponseMessage:
		payload := new(KarlsendMessage_NotifyUtxosChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UTXOsChangedNotificationMessage:
		payload := new(KarlsendMessage_UtxosChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingUTXOsChangedRequestMessage:
		payload := new(KarlsendMessage_StopNotifyingUtxosChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingUTXOsChangedResponseMessage:
		payload := new(KarlsendMessage_StopNotifyingUtxosChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetUTXOsByAddressesRequestMessage:
		payload := new(KarlsendMessage_GetUtxosByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetUTXOsByAddressesResponseMessage:
		payload := new(KarlsendMessage_GetUtxosByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressRequestMessage:
		payload := new(KarlsendMessage_GetBalanceByAddressRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressResponseMessage:
		payload := new(KarlsendMessage_GetBalanceByAddressResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentBlueScoreRequestMessage:
		payload := new(KarlsendMessage_GetVirtualSelectedParentBlueScoreRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentBlueScoreResponseMessage:
		payload := new(KarlsendMessage_GetVirtualSelectedParentBlueScoreResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage:
		payload := new(KarlsendMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage:
		payload := new(KarlsendMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage:
		payload := new(KarlsendMessage_VirtualSelectedParentBlueScoreChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BanRequestMessage:
		payload := new(KarlsendMessage_BanRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BanResponseMessage:
		payload := new(KarlsendMessage_BanResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UnbanRequestMessage:
		payload := new(KarlsendMessage_UnbanRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UnbanResponseMessage:
		payload := new(KarlsendMessage_UnbanResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetInfoRequestMessage:
		payload := new(KarlsendMessage_GetInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetInfoResponseMessage:
		payload := new(KarlsendMessage_GetInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyPruningPointUTXOSetOverrideRequestMessage:
		payload := new(KarlsendMessage_NotifyPruningPointUTXOSetOverrideRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyPruningPointUTXOSetOverrideResponseMessage:
		payload := new(KarlsendMessage_NotifyPruningPointUTXOSetOverrideResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.PruningPointUTXOSetOverrideNotificationMessage:
		payload := new(KarlsendMessage_PruningPointUTXOSetOverrideNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingPruningPointUTXOSetOverrideRequestMessage:
		payload := new(KarlsendMessage_StopNotifyingPruningPointUTXOSetOverrideRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondRequestMessage:
		payload := new(KarlsendMessage_EstimateNetworkHashesPerSecondRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondResponseMessage:
		payload := new(KarlsendMessage_EstimateNetworkHashesPerSecondResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualDaaScoreChangedRequestMessage:
		payload := new(KarlsendMessage_NotifyVirtualDaaScoreChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualDaaScoreChangedResponseMessage:
		payload := new(KarlsendMessage_NotifyVirtualDaaScoreChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualDaaScoreChangedNotificationMessage:
		payload := new(KarlsendMessage_VirtualDaaScoreChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalancesByAddressesRequestMessage:
		payload := new(KarlsendMessage_GetBalancesByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalancesByAddressesResponseMessage:
		payload := new(KarlsendMessage_GetBalancesByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateRequestMessage:
		payload := new(KarlsendMessage_NotifyNewBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateResponseMessage:
		payload := new(KarlsendMessage_NotifyNewBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NewBlockTemplateNotificationMessage:
		payload := new(KarlsendMessage_NewBlockTemplateNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesByAddressesRequestMessage:
		payload := new(KarlsendMessage_GetMempoolEntriesByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesByAddressesResponseMessage:
		payload := new(KarlsendMessage_GetMempoolEntriesByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCoinSupplyRequestMessage:
		payload := new(KarlsendMessage_GetCoinSupplyRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCoinSupplyResponseMessage:
		payload := new(KarlsendMessage_GetCoinSupplyResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetFeeEstimateRequestMessage:
		payload := new(KarlsendMessage_GetFeeEstimateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionReplacementRequestMessage:
		payload := new(KarlsendMessage_SubmitTransactionReplacementRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionReplacementResponseMessage:
		payload := new(KarlsendMessage_SubmitTransactionReplacementResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	default:
		return nil, nil
	}
}
