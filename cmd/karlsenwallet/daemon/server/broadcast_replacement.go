package server

import (
	"context"
	"time"

	"github.com/karlsen-network/karlsend/v2/app/appmessage"
	"github.com/karlsen-network/karlsend/v2/cmd/karlsenwallet/daemon/pb"
	"github.com/karlsen-network/karlsend/v2/cmd/karlsenwallet/libkarlsenwallet"
	"github.com/karlsen-network/karlsend/v2/cmd/karlsenwallet/libkarlsenwallet/serialization"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/model/externalapi"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/utils/consensushashing"
	"github.com/karlsen-network/karlsend/v2/infrastructure/network/rpcclient"
	"github.com/pkg/errors"
)

func (s *server) BroadcastReplacement(_ context.Context, request *pb.BroadcastRequest) (*pb.BroadcastResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	txIDs, err := s.broadcastReplacement(request.Transactions, request.IsDomain)
	if err != nil {
		return nil, err
	}

	return &pb.BroadcastResponse{TxIDs: txIDs}, nil
}

// broadcastReplacement assumes that all transactions depend on the first one
func (s *server) broadcastReplacement(transactions [][]byte, isDomain bool) ([]string, error) {

	txIDs := make([]string, len(transactions))
	var tx *externalapi.DomainTransaction
	var err error

	for i, transaction := range transactions {

		if isDomain {
			tx, err = serialization.DeserializeDomainTransaction(transaction)
			if err != nil {
				return nil, err
			}
		} else if !isDomain { //default in proto3 is false
			tx, err = libkarlsenwallet.ExtractTransaction(transaction, s.keysFile.ECDSA)
			if err != nil {
				return nil, err
			}
		}

		// Once the first transaction is added to the mempool, the transactions that depend
		// on the replaced transaction will be removed, so there's no need to submit them
		// as RBF transactions.
		if i == 0 {
			txIDs[i], err = sendTransactionRBF(s.rpcClient, tx)
			if err != nil {
				return nil, err
			}
		} else {
			txIDs[i], err = sendTransaction(s.rpcClient, tx)
			if err != nil {
				return nil, err
			}

		}

		for _, input := range tx.Inputs {
			s.usedOutpoints[input.PreviousOutpoint] = time.Now()
		}
	}

	s.forceSync()
	return txIDs, nil
}

func sendTransactionRBF(client *rpcclient.RPCClient, tx *externalapi.DomainTransaction) (string, error) {
	submitTransactionResponse, err := client.SubmitTransactionReplacement(appmessage.DomainTransactionToRPCTransaction(tx), consensushashing.TransactionID(tx).String())
	if err != nil {
		return "", errors.Wrapf(err, "error submitting transaction replacement")
	}
	return submitTransactionResponse.TransactionID, nil
}
