package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/karlsen-network/karlsend/v2/cmd/karlsenwallet/daemon/server"
	"github.com/karlsen-network/karlsend/v2/cmd/karlsenwallet/keys"
	"github.com/karlsen-network/karlsend/v2/cmd/karlsenwallet/libkarlsenwallet/serialization"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/utils/consensushashing"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/utils/constants"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/utils/txscript"
	"github.com/karlsen-network/karlsend/v2/util/txmass"
	"github.com/pkg/errors"
)

func parse(conf *parseConfig) error {
	if conf.Transaction == "" && conf.TransactionFile == "" {
		return errors.Errorf("Either --transaction or --transaction-file is required")
	}
	if conf.Transaction != "" && conf.TransactionFile != "" {
		return errors.Errorf("Both --transaction and --transaction-file cannot be passed at the same time")
	}

	keysFile, err := keys.ReadKeysFile(conf.NetParams(), conf.KeysFile)
	if err != nil {
		return err
	}

	transactionHex := conf.Transaction
	if conf.TransactionFile != "" {
		transactionHexBytes, err := ioutil.ReadFile(conf.TransactionFile)
		if err != nil {
			return errors.Wrapf(err, "Could not read hex from %s", conf.TransactionFile)
		}
		transactionHex = strings.TrimSpace(string(transactionHexBytes))
	}

	transactions, err := server.DecodeTransactionsFromHex(transactionHex)
	if err != nil {
		return err
	}

	txMassCalculator := txmass.NewCalculator(conf.NetParams().MassPerTxByte, conf.NetParams().MassPerScriptPubKeyByte, conf.NetParams().MassPerSigOp)
	for i, transaction := range transactions {

		partiallySignedTransaction, err := serialization.DeserializePartiallySignedTransaction(transaction)
		if err != nil {
			return err
		}

		fmt.Printf("Transaction #%d ID: \t%s\n", i+1, consensushashing.TransactionID(partiallySignedTransaction.Tx))
		fmt.Println()

		allInputSompi := uint64(0)
		for index, input := range partiallySignedTransaction.Tx.Inputs {
			partiallySignedInput := partiallySignedTransaction.PartiallySignedInputs[index]

			if conf.Verbose {
				fmt.Printf("Input %d: \tOutpoint: %s:%d \tAmount: %.2f Karlsen\n", index, input.PreviousOutpoint.TransactionID,
					input.PreviousOutpoint.Index, float64(partiallySignedInput.PrevOutput.Value)/float64(constants.SompiPerKarlsen))
			}

			allInputSompi += partiallySignedInput.PrevOutput.Value
		}
		if conf.Verbose {
			fmt.Println()
		}

		allOutputSompi := uint64(0)
		for index, output := range partiallySignedTransaction.Tx.Outputs {
			scriptPublicKeyType, scriptPublicKeyAddress, err := txscript.ExtractScriptPubKeyAddress(output.ScriptPublicKey, conf.ActiveNetParams)
			if err != nil {
				return err
			}

			addressString := scriptPublicKeyAddress.EncodeAddress()
			if scriptPublicKeyType == txscript.NonStandardTy {
				scriptPublicKeyHex := hex.EncodeToString(output.ScriptPublicKey.Script)
				addressString = fmt.Sprintf("<Non-standard transaction script public key: %s>", scriptPublicKeyHex)
			}

			fmt.Printf("Output %d: \tRecipient: %s \tAmount: %.2f Karlsen\n",
				index, addressString, float64(output.Value)/float64(constants.SompiPerKarlsen))

			allOutputSompi += output.Value
		}
		fmt.Println()

		fee := allInputSompi - allOutputSompi
		fmt.Printf("Fee:\t%d Sompi (%f KLS)\n", fee, float64(fee)/float64(constants.SompiPerKarlsen))
		mass, err := server.EstimateMassAfterSignatures(partiallySignedTransaction, keysFile.ECDSA, keysFile.MinimumSignatures, txMassCalculator)
		if err != nil {
			return err
		}

		fmt.Printf("Mass: %d grams\n", mass)
		feeRate := float64(fee) / float64(mass)
		fmt.Printf("Fee rate: %.2f Sompi/Gram\n", feeRate)
	}

	return nil
}
