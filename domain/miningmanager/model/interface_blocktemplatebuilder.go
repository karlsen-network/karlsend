package model

import (
	consensusexternalapi "github.com/karlsen-network/karlsend/domain/consensus/model/externalapi"
)

// BlockTemplateBuilder builds block templates for miners to consume
type BlockTemplateBuilder interface {
	BuildBlockTemplate(coinbaseData *consensusexternalapi.DomainCoinbaseData) (*consensusexternalapi.DomainBlockTemplate, error)
	ModifyBlockTemplate(newCoinbaseData *consensusexternalapi.DomainCoinbaseData,
		blockTemplateToModify *consensusexternalapi.DomainBlockTemplate) (*consensusexternalapi.DomainBlockTemplate, error)
}
