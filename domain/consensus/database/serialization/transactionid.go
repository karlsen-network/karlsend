package serialization

import (
	"github.com/karlsen-network/karlsend/domain/consensus/model/externalapi"
	"github.com/karlsen-network/karlsend/domain/consensus/utils/transactionid"
)

// DbTransactionIDToDomainTransactionID converts DbTransactionId to DomainTransactionID
func DbTransactionIDToDomainTransactionID(dbTransactionID *DbTransactionId) (*externalapi.DomainTransactionID, error) {
	return transactionid.FromBytes(dbTransactionID.TransactionId)
}

// DomainTransactionIDToDbTransactionID converts DomainTransactionID to DbTransactionId
func DomainTransactionIDToDbTransactionID(domainTransactionID *externalapi.DomainTransactionID) *DbTransactionId {
	return &DbTransactionId{TransactionId: domainTransactionID.ByteSlice()}
}
