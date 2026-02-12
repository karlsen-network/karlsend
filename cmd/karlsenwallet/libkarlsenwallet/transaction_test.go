package libkarlsenwallet_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/karlsen-network/karlsend/v2/cmd/karlsenwallet/libkarlsenwallet/serialization"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/utils/constants"

	"github.com/karlsen-network/karlsend/v2/cmd/karlsenwallet/libkarlsenwallet"
	"github.com/karlsen-network/karlsend/v2/domain/consensus"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/model/externalapi"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/utils/consensushashing"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/utils/testutils"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/utils/txscript"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/utils/utxo"
	"github.com/karlsen-network/karlsend/v2/util"
)

func forSchnorrAndECDSA(t *testing.T, testFunc func(t *testing.T, ecdsa bool)) {
	t.Run("schnorr", func(t *testing.T) {
		testFunc(t, false)
	})

	t.Run("ecdsa", func(t *testing.T) {
		testFunc(t, true)
	})
}

func createUnsignedTransactionSerialized(
	extendedPublicKeys []string,
	minimumSignatures uint32,
	payments []*libkarlsenwallet.Payment,
	selectedUTXOs []*libkarlsenwallet.UTXO) ([]byte, error) {

	tx, err := libkarlsenwallet.CreateUnsignedTransaction(extendedPublicKeys, minimumSignatures, payments, selectedUTXOs)
	if err != nil {
		return nil, err
	}

	return serialization.SerializePartiallySignedTransaction(tx)
}

func TestMultisig(t *testing.T) {
	testutils.ForAllPaths(t, func(t *testing.T, version uint32) {
		testutils.ForAllNets(t, true, func(t *testing.T, consensusConfig *consensus.Config) {
			params := &consensusConfig.Params
			forSchnorrAndECDSA(t, func(t *testing.T, ecdsa bool) {
				consensusConfig.BlockCoinbaseMaturity = 0
				tc, teardown, err := consensus.NewFactory().NewTestConsensus(consensusConfig, "TestMultisig")
				if err != nil {
					t.Fatalf("Error setting up tc: %+v", err)
				}
				defer teardown(false)

				const numKeys = 3
				mnemonics := make([]string, numKeys)
				publicKeys := make([]string, numKeys)
				for i := 0; i < numKeys; i++ {
					var err error
					mnemonics[i], err = libkarlsenwallet.CreateMnemonic()
					if err != nil {
						t.Fatalf("CreateMnemonic: %+v", err)
					}

					publicKeys[i], err = libkarlsenwallet.MasterPublicKeyFromMnemonic(&consensusConfig.Params, mnemonics[i], true, version)
					if err != nil {
						t.Fatalf("MasterPublicKeyFromMnemonic: %+v", err)
					}
				}

				const minimumSignatures = 2
				path := "m/1/2/3"
				address, err := libkarlsenwallet.Address(params, publicKeys, minimumSignatures, path, ecdsa)
				if err != nil {
					t.Fatalf("Address: %+v", err)
				}

				if _, ok := address.(*util.AddressScriptHash); !ok {
					t.Fatalf("The address is of unexpected type")
				}

				scriptPublicKey, err := txscript.PayToAddrScript(address)
				if err != nil {
					t.Fatalf("PayToAddrScript: %+v", err)
				}

				coinbaseData := &externalapi.DomainCoinbaseData{
					ScriptPublicKey: scriptPublicKey,
					ExtraData:       nil,
				}

				fundingBlockHash, _, err := tc.AddBlock([]*externalapi.DomainHash{consensusConfig.GenesisHash}, coinbaseData, nil)
				if err != nil {
					t.Fatalf("AddBlock: %+v", err)
				}

				block1Hash, _, err := tc.AddBlock([]*externalapi.DomainHash{fundingBlockHash}, nil, nil)
				if err != nil {
					t.Fatalf("AddBlock: %+v", err)
				}

				block1, _, err := tc.GetBlock(block1Hash)
				if err != nil {
					t.Fatalf("GetBlock: %+v", err)
				}

				block1Tx := block1.Transactions[0]
				block1TxOut := block1Tx.Outputs[0]
				selectedUTXOs := []*libkarlsenwallet.UTXO{
					{
						Outpoint: &externalapi.DomainOutpoint{
							TransactionID: *consensushashing.TransactionID(block1.Transactions[0]),
							Index:         0,
						},
						UTXOEntry:      utxo.NewUTXOEntry(block1TxOut.Value, block1TxOut.ScriptPublicKey, true, 0),
						DerivationPath: path,
					},
				}

				unsignedTransaction, err := createUnsignedTransactionSerialized(publicKeys, minimumSignatures,
					[]*libkarlsenwallet.Payment{{
						Address: address,
						Amount:  10,
					}}, selectedUTXOs)
				if err != nil {
					t.Fatalf("CreateUnsignedTransactions: %+v", err)
				}

				isFullySigned, err := libkarlsenwallet.IsTransactionFullySigned(unsignedTransaction)
				if err != nil {
					t.Fatalf("IsTransactionFullySigned: %+v", err)
				}

				if isFullySigned {
					t.Fatalf("Transaction is not expected to be signed")
				}

				_, err = libkarlsenwallet.ExtractTransaction(unsignedTransaction, ecdsa)
				if err == nil || !strings.Contains(err.Error(), fmt.Sprintf("missing %d signatures", minimumSignatures)) {
					t.Fatal("Unexpectedly succeed to extract a valid transaction out of unsigned transaction")
				}

				signedTxStep1, err := libkarlsenwallet.Sign(params, mnemonics[:1], unsignedTransaction, ecdsa, version)
				if err != nil {
					t.Fatalf("Sign: %+v", err)
				}

				isFullySigned, err = libkarlsenwallet.IsTransactionFullySigned(signedTxStep1)
				if err != nil {
					t.Fatalf("IsTransactionFullySigned: %+v", err)
				}

				if isFullySigned {
					t.Fatalf("Transaction is not expected to be fully signed")
				}

				signedTxStep2, err := libkarlsenwallet.Sign(params, mnemonics[1:2], signedTxStep1, ecdsa, version)
				if err != nil {
					t.Fatalf("Sign: %+v", err)
				}

				extractedSignedTxStep2, err := libkarlsenwallet.ExtractTransaction(signedTxStep2, ecdsa)
				if err != nil {
					t.Fatalf("ExtractTransaction: %+v", err)
				}

				signedTxOneStep, err := libkarlsenwallet.Sign(params, mnemonics[:2], unsignedTransaction, ecdsa, version)
				if err != nil {
					t.Fatalf("Sign: %+v", err)
				}

				extractedSignedTxOneStep, err := libkarlsenwallet.ExtractTransaction(signedTxOneStep, ecdsa)
				if err != nil {
					t.Fatalf("ExtractTransaction: %+v", err)
				}

				// We check IDs instead of comparing the actual transactions because the actual transactions have different
				// signature scripts due to non deterministic signature scheme.
				if !consensushashing.TransactionID(extractedSignedTxStep2).Equal(consensushashing.TransactionID(extractedSignedTxOneStep)) {
					t.Fatalf("Expected extractedSignedTxOneStep and extractedSignedTxStep2 IDs to be equal")
				}

				_, virtualChangeSet, err := tc.AddBlock([]*externalapi.DomainHash{block1Hash}, nil, []*externalapi.DomainTransaction{extractedSignedTxStep2})
				if err != nil {
					t.Fatalf("AddBlock: %+v", err)
				}

				addedUTXO := &externalapi.DomainOutpoint{
					TransactionID: *consensushashing.TransactionID(extractedSignedTxStep2),
					Index:         0,
				}
				if !virtualChangeSet.VirtualUTXODiff.ToAdd().Contains(addedUTXO) {
					t.Fatalf("Transaction wasn't accepted in the DAG")
				}
			})
		})
	})
}

func TestP2PK(t *testing.T) {
	testutils.ForAllPaths(t, func(t *testing.T, version uint32) {
		testutils.ForAllNets(t, true, func(t *testing.T, consensusConfig *consensus.Config) {
			params := &consensusConfig.Params
			forSchnorrAndECDSA(t, func(t *testing.T, ecdsa bool) {
				consensusConfig.BlockCoinbaseMaturity = 0
				tc, teardown, err := consensus.NewFactory().NewTestConsensus(consensusConfig, "TestMultisig")
				if err != nil {
					t.Fatalf("Error setting up tc: %+v", err)
				}
				defer teardown(false)

				const numKeys = 1
				mnemonics := make([]string, numKeys)
				publicKeys := make([]string, numKeys)
				for i := 0; i < numKeys; i++ {
					var err error
					mnemonics[i], err = libkarlsenwallet.CreateMnemonic()
					if err != nil {
						t.Fatalf("CreateMnemonic: %+v", err)
					}

					publicKeys[i], err = libkarlsenwallet.MasterPublicKeyFromMnemonic(&consensusConfig.Params, mnemonics[i], false, version)
					if err != nil {
						t.Fatalf("MasterPublicKeyFromMnemonic: %+v", err)
					}
				}

				const minimumSignatures = 1
				path := "m/1/2/3"
				address, err := libkarlsenwallet.Address(params, publicKeys, minimumSignatures, path, ecdsa)
				if err != nil {
					t.Fatalf("Address: %+v", err)
				}

				if ecdsa {
					if _, ok := address.(*util.AddressPublicKeyECDSA); !ok {
						t.Fatalf("The address is of unexpected type")
					}
				} else {
					if _, ok := address.(*util.AddressPublicKey); !ok {
						t.Fatalf("The address is of unexpected type")
					}
				}

				scriptPublicKey, err := txscript.PayToAddrScript(address)
				if err != nil {
					t.Fatalf("PayToAddrScript: %+v", err)
				}

				coinbaseData := &externalapi.DomainCoinbaseData{
					ScriptPublicKey: scriptPublicKey,
					ExtraData:       nil,
				}

				fundingBlockHash, _, err := tc.AddBlock([]*externalapi.DomainHash{consensusConfig.GenesisHash}, coinbaseData, nil)
				if err != nil {
					t.Fatalf("AddBlock: %+v", err)
				}

				block1Hash, _, err := tc.AddBlock([]*externalapi.DomainHash{fundingBlockHash}, nil, nil)
				if err != nil {
					t.Fatalf("AddBlock: %+v", err)
				}

				block1, _, err := tc.GetBlock(block1Hash)
				if err != nil {
					t.Fatalf("GetBlock: %+v", err)
				}

				block1Tx := block1.Transactions[0]
				block1TxOut := block1Tx.Outputs[0]
				selectedUTXOs := []*libkarlsenwallet.UTXO{
					{
						Outpoint: &externalapi.DomainOutpoint{
							TransactionID: *consensushashing.TransactionID(block1.Transactions[0]),
							Index:         0,
						},
						UTXOEntry:      utxo.NewUTXOEntry(block1TxOut.Value, block1TxOut.ScriptPublicKey, true, 0),
						DerivationPath: path,
					},
				}

				unsignedTransaction, err := createUnsignedTransactionSerialized(publicKeys, minimumSignatures,
					[]*libkarlsenwallet.Payment{{
						Address: address,
						Amount:  10,
					}}, selectedUTXOs)
				if err != nil {
					t.Fatalf("CreateUnsignedTransactions: %+v", err)
				}

				isFullySigned, err := libkarlsenwallet.IsTransactionFullySigned(unsignedTransaction)
				if err != nil {
					t.Fatalf("IsTransactionFullySigned: %+v", err)
				}

				if isFullySigned {
					t.Fatalf("Transaction is not expected to be signed")
				}

				_, err = libkarlsenwallet.ExtractTransaction(unsignedTransaction, ecdsa)
				if err == nil || !strings.Contains(err.Error(), "missing signature") {
					t.Fatal("Unexpectedly succeed to extract a valid transaction out of unsigned transaction")
				}

				signedTx, err := libkarlsenwallet.Sign(params, mnemonics, unsignedTransaction, ecdsa, version)
				if err != nil {
					t.Fatalf("Sign: %+v", err)
				}

				tx, err := libkarlsenwallet.ExtractTransaction(signedTx, ecdsa)
				if err != nil {
					t.Fatalf("ExtractTransaction: %+v", err)
				}

				_, virtualChangeSet, err := tc.AddBlock([]*externalapi.DomainHash{block1Hash}, nil, []*externalapi.DomainTransaction{tx})
				if err != nil {
					t.Fatalf("AddBlock: %+v", err)
				}

				addedUTXO := &externalapi.DomainOutpoint{
					TransactionID: *consensushashing.TransactionID(tx),
					Index:         0,
				}
				if !virtualChangeSet.VirtualUTXODiff.ToAdd().Contains(addedUTXO) {
					t.Fatalf("Transaction wasn't accepted in the DAG")
				}
			})
		})
	})
}

func TestMaxSompi(t *testing.T) {
	testutils.ForAllPaths(t, func(t *testing.T, version uint32) {
		testutils.ForAllNets(t, true, func(t *testing.T, consensusConfig *consensus.Config) {
			params := &consensusConfig.Params
			cfg := *consensusConfig
			cfg.BlockCoinbaseMaturity = 0
			cfg.PreDeflationaryPhaseBaseSubsidy = 20e6 * constants.SompiPerKarlsen
			tc, teardown, err := consensus.NewFactory().NewTestConsensus(&cfg, "TestMaxSompi")
			if err != nil {
				t.Fatalf("Error setting up tc: %+v", err)
			}
			defer teardown(false)

			const numKeys = 1
			mnemonics := make([]string, numKeys)
			publicKeys := make([]string, numKeys)
			for i := 0; i < numKeys; i++ {
				var err error
				mnemonics[i], err = libkarlsenwallet.CreateMnemonic()
				if err != nil {
					t.Fatalf("CreateMnemonic: %+v", err)
				}

				publicKeys[i], err = libkarlsenwallet.MasterPublicKeyFromMnemonic(&cfg.Params, mnemonics[i], false, version)
				if err != nil {
					t.Fatalf("MasterPublicKeyFromMnemonic: %+v", err)
				}
			}

			const minimumSignatures = 1
			path := "m/1/2/3"
			address, err := libkarlsenwallet.Address(params, publicKeys, minimumSignatures, path, false)
			if err != nil {
				t.Fatalf("Address: %+v", err)
			}

			scriptPublicKey, err := txscript.PayToAddrScript(address)
			if err != nil {
				t.Fatalf("PayToAddrScript: %+v", err)
			}

			coinbaseData := &externalapi.DomainCoinbaseData{
				ScriptPublicKey: scriptPublicKey,
				ExtraData:       nil,
			}

			fundingBlock1Hash, _, err := tc.AddBlock([]*externalapi.DomainHash{cfg.GenesisHash}, coinbaseData, nil)
			if err != nil {
				t.Fatalf("AddBlock: %+v", err)
			}

			fundingBlock2Hash, _, err := tc.AddBlock([]*externalapi.DomainHash{fundingBlock1Hash}, coinbaseData, nil)
			if err != nil {
				t.Fatalf("AddBlock: %+v", err)
			}

			fundingBlock3Hash, _, err := tc.AddBlock([]*externalapi.DomainHash{fundingBlock2Hash}, coinbaseData, nil)
			if err != nil {
				t.Fatalf("AddBlock: %+v", err)
			}

			fundingBlock4Hash, _, err := tc.AddBlock([]*externalapi.DomainHash{fundingBlock3Hash}, coinbaseData, nil)
			if err != nil {
				t.Fatalf("AddBlock: %+v", err)
			}

			fundingBlock2, _, err := tc.GetBlock(fundingBlock2Hash)
			if err != nil {
				t.Fatalf("GetBlock: %+v", err)
			}

			fundingBlock3, _, err := tc.GetBlock(fundingBlock3Hash)
			if err != nil {
				t.Fatalf("GetBlock: %+v", err)
			}

			fundingBlock4, _, err := tc.GetBlock(fundingBlock4Hash)
			if err != nil {
				t.Fatalf("GetBlock: %+v", err)
			}

			block1Hash, _, err := tc.AddBlock([]*externalapi.DomainHash{fundingBlock4Hash}, nil, nil)
			if err != nil {
				t.Fatalf("AddBlock: %+v", err)
			}

			block1, _, err := tc.GetBlock(block1Hash)
			if err != nil {
				t.Fatalf("GetBlock: %+v", err)
			}

			txOut1 := fundingBlock2.Transactions[0].Outputs[0]
			txOut2 := fundingBlock3.Transactions[0].Outputs[0]
			txOut3 := fundingBlock4.Transactions[0].Outputs[0]
			txOut4 := block1.Transactions[0].Outputs[0]
			selectedUTXOsForTxWithLargeInputAmount := []*libkarlsenwallet.UTXO{
				{
					Outpoint: &externalapi.DomainOutpoint{
						TransactionID: *consensushashing.TransactionID(fundingBlock2.Transactions[0]),
						Index:         0,
					},
					UTXOEntry:      utxo.NewUTXOEntry(txOut1.Value, txOut1.ScriptPublicKey, true, 0),
					DerivationPath: path,
				},
				{
					Outpoint: &externalapi.DomainOutpoint{
						TransactionID: *consensushashing.TransactionID(fundingBlock3.Transactions[0]),
						Index:         0,
					},
					UTXOEntry:      utxo.NewUTXOEntry(txOut2.Value, txOut2.ScriptPublicKey, true, 0),
					DerivationPath: path,
				},
			}

			unsignedTxWithLargeInputAmount, err := createUnsignedTransactionSerialized(publicKeys, minimumSignatures,
				[]*libkarlsenwallet.Payment{{
					Address: address,
					Amount:  10,
				}}, selectedUTXOsForTxWithLargeInputAmount)
			if err != nil {
				t.Fatalf("CreateUnsignedTransactions: %+v", err)
			}

			signedTxWithLargeInputAmount, err := libkarlsenwallet.Sign(params, mnemonics, unsignedTxWithLargeInputAmount, false, version)
			if err != nil {
				t.Fatalf("Sign: %+v", err)
			}

			txWithLargeInputAmount, err := libkarlsenwallet.ExtractTransaction(signedTxWithLargeInputAmount, false)
			if err != nil {
				t.Fatalf("ExtractTransaction: %+v", err)
			}

			_, virtualChangeSet, err := tc.AddBlock([]*externalapi.DomainHash{block1Hash}, nil, []*externalapi.DomainTransaction{txWithLargeInputAmount})
			if err != nil {
				t.Fatalf("AddBlock: %+v", err)
			}

			addedUTXO1 := &externalapi.DomainOutpoint{
				TransactionID: *consensushashing.TransactionID(txWithLargeInputAmount),
				Index:         0,
			}
			if !virtualChangeSet.VirtualUTXODiff.ToAdd().Contains(addedUTXO1) {
				t.Fatalf("Transaction wasn't accepted in the DAG")
			}

			selectedUTXOsForTxWithLargeInputAndOutputAmount := []*libkarlsenwallet.UTXO{
				{
					Outpoint: &externalapi.DomainOutpoint{
						TransactionID: *consensushashing.TransactionID(fundingBlock4.Transactions[0]),
						Index:         0,
					},
					UTXOEntry:      utxo.NewUTXOEntry(txOut3.Value, txOut3.ScriptPublicKey, true, 0),
					DerivationPath: path,
				},
				{
					Outpoint: &externalapi.DomainOutpoint{
						TransactionID: *consensushashing.TransactionID(block1.Transactions[0]),
						Index:         0,
					},
					UTXOEntry:      utxo.NewUTXOEntry(txOut4.Value, txOut4.ScriptPublicKey, true, 0),
					DerivationPath: path,
				},
			}

			unsignedTxWithLargeInputAndOutputAmount, err := createUnsignedTransactionSerialized(publicKeys, minimumSignatures,
				[]*libkarlsenwallet.Payment{{
					Address: address,
					Amount:  22e6 * constants.SompiPerKarlsen,
				}}, selectedUTXOsForTxWithLargeInputAndOutputAmount)
			if err != nil {
				t.Fatalf("CreateUnsignedTransactions: %+v", err)
			}

			signedTxWithLargeInputAndOutputAmount, err := libkarlsenwallet.Sign(params, mnemonics, unsignedTxWithLargeInputAndOutputAmount, false, version)
			if err != nil {
				t.Fatalf("Sign: %+v", err)
			}

			txWithLargeInputAndOutputAmount, err := libkarlsenwallet.ExtractTransaction(signedTxWithLargeInputAndOutputAmount, false)
			if err != nil {
				t.Fatalf("ExtractTransaction: %+v", err)
			}

			// We're creating a new longer chain so we can double spend txWithLargeInputAmount
			newChainRoot, _, err := tc.AddBlock([]*externalapi.DomainHash{block1Hash}, nil, nil)
			if err != nil {
				t.Fatalf("AddBlock: %+v", err)
			}

			_, virtualChangeSet, err = tc.AddBlock([]*externalapi.DomainHash{newChainRoot}, nil, []*externalapi.DomainTransaction{txWithLargeInputAndOutputAmount})
			if err != nil {
				t.Fatalf("AddBlock: %+v", err)
			}

			addedUTXO2 := &externalapi.DomainOutpoint{
				TransactionID: *consensushashing.TransactionID(txWithLargeInputAndOutputAmount),
				Index:         0,
			}

			if !virtualChangeSet.VirtualUTXODiff.ToAdd().Contains(addedUTXO2) {
				t.Fatalf("txWithLargeInputAndOutputAmount weren't accepted in the DAG")
			}
		})
	})
}
