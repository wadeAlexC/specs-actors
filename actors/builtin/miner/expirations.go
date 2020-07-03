package miner

import (
	"sort"

	"github.com/filecoin-project/specs-actors/actors/abi"
)

type sectorEpochSet struct {
	epoch   abi.ChainEpoch
	sectors []uint64
}

func groupSectorsByExpiration(sectors []*SectorOnChainInfo) []sectorEpochSet {
	sectorsByExpiration := make(map[abi.ChainEpoch][]uint64)

	for _, sector := range sectors {
		sectorsByExpiration[sector.Expiration] = append(sectorsByExpiration[sector.Expiration], uint64(sector.SectorNumber))
	}

	sectorEpochSets := make([]sectorEpochSet, 0, len(sectorsByExpiration))
	for expiration, sectors := range sectorsByExpiration {
		sectorEpochSets = append(sectorEpochSets, sectorEpochSet{expiration, sectors})
	}
	sort.Slice(sectorEpochSets, func(i, j int) bool {
		return sectorEpochSets[i].epoch < sectorEpochSets[j].epoch
	})
	return sectorEpochSets
}
