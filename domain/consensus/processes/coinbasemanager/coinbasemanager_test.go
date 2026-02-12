package coinbasemanager

import (
	"math"
	"strconv"
	"testing"

	"github.com/karlsen-network/karlsend/v2/domain/consensus/model/externalapi"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/utils/constants"
	"github.com/karlsen-network/karlsend/v2/domain/dagconfig"
)

func TestCalcDeflationaryPeriodBlockSubsidy(t *testing.T) {
	const secondsPerMonth = 2629800
	const secondsPerHalving = secondsPerMonth * 12
	const deflationaryPhaseDaaScore = secondsPerMonth * 6
	const deflationaryPhaseBaseSubsidy = 44 * constants.SompiPerKarlsen
	coinbaseManagerInterface := New(
		nil,
		0,
		0,
		0,
		&externalapi.DomainHash{},
		deflationaryPhaseDaaScore,
		deflationaryPhaseBaseSubsidy,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil)
	coinbaseManagerInstance := coinbaseManagerInterface.(*coinbaseManager)

	tests := []struct {
		name                 string
		blockDaaScore        uint64
		expectedBlockSubsidy uint64
	}{
		{
			name:                 "start of deflationary phase",
			blockDaaScore:        deflationaryPhaseDaaScore,
			expectedBlockSubsidy: deflationaryPhaseBaseSubsidy,
		},
		{
			name:                 "after 1 year",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / 1.4)),
		},
		{
			name:                 "after 2 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*2,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / math.Pow(1.4, 2))),
		},
		{
			name:                 "after 5 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*5,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / math.Pow(1.4, 5))),
		},
		{
			name:                 "after 32 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*32,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / math.Pow(1.4, 32))),
		},
		{
			name:                 "after 64 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*64,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / math.Pow(1.4, 64))),
		},
		{
			name:                 "just before subsidy depleted",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*65,
			expectedBlockSubsidy: 1,
		},
		{
			name:                 "after subsidy depleted",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*66,
			expectedBlockSubsidy: 0,
		},
	}

	for _, test := range tests {
		blockSubsidy := coinbaseManagerInstance.calcDeflationaryPeriodBlockSubsidy(test.blockDaaScore)
		if blockSubsidy != test.expectedBlockSubsidy {
			t.Errorf("TestCalcDeflationaryPeriodBlockSubsidy: test '%s' failed. Want: %d, got: %d",
				test.name, test.expectedBlockSubsidy, blockSubsidy)
		}
	}
}

func TestBuildSubsidyTable(t *testing.T) {
	deflationaryPhaseBaseSubsidy := dagconfig.MainnetParams.DeflationaryPhaseBaseSubsidy
	if deflationaryPhaseBaseSubsidy != 44*constants.SompiPerKarlsen {
		t.Errorf("TestBuildSubsidyTable: table generation function was not updated to reflect "+
			"the new base subsidy %d. Please fix the constant above and replace subsidyByDeflationaryMonthTable "+
			"in coinbasemanager.go with the printed table", deflationaryPhaseBaseSubsidy)
	}
	coinbaseManagerInterface := New(
		nil,
		0,
		0,
		0,
		&externalapi.DomainHash{},
		0,
		deflationaryPhaseBaseSubsidy,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil)
	coinbaseManagerInstance := coinbaseManagerInterface.(*coinbaseManager)

	var subsidyTable []uint64
	for M := uint64(0); ; M++ {
		subsidy := coinbaseManagerInstance.calcDeflationaryPeriodBlockSubsidyFloatCalc(M)
		subsidyTable = append(subsidyTable, subsidy)
		if subsidy == 0 {
			break
		}
	}

	tableStr := "\n{\t"
	for i := 0; i < len(subsidyTable); i++ {
		tableStr += strconv.FormatUint(subsidyTable[i], 10) + ", "
		if (i+1)%25 == 0 {
			tableStr += "\n\t"
		}
	}
	tableStr += "\n}"
	t.Logf("%s", tableStr)
}
