package comsoc

import (
	"errors"
	"slices"
)

func TieBreakFactory(orderedAlts []Alternative) func([]Alternative) (Alternative, error) {
	return func(alts []Alternative) (Alternative, error) {
		if len(orderedAlts) == 0 {
			return -1, errors.New("(TieBreakFactory) - orderedAlts is null")
		}

		if len(alts) == 0 {
			return -1, errors.New("(TieBreakFactory) - alts is null")
		}

		//on renvoie le premier élément que orderedAlts a en commun avec alts
		for i := range orderedAlts {
			if slices.Contains(alts, orderedAlts[i]) {
				return orderedAlts[i], nil
			}
		}

		return Alternative(-1), errors.New("(TieBreakFactory) - no common element between orderedAlts and alts")
	}
}

func SWFFactory(swf func(p Profile) (Count, error), tb func([]Alternative) (Alternative, error)) func(Profile) ([]Alternative, error) {
	return func(p Profile) ([]Alternative, error) {
		count, errSWF := swf(p)

		if errSWF != nil {
			return nil, errSWF
		}

		var sortedAlts []Alternative

		for len(count) > 0 {
			//récupération des meilleurs alternatives
			bestAlts := maxCount(count)

			for i := range bestAlts {
				delete(count, bestAlts[i])
			}

			for len(bestAlts) > 0 {
				bestAlt, errTB := tb(bestAlts)
				if errTB != nil {
					return nil, errTB
				}
				sortedAlts = append(sortedAlts, bestAlt)
				indice := rank(bestAlt, bestAlts)
				bestAlts = slices.Delete(bestAlts, indice, indice+1)
			}
		}
		return sortedAlts, nil
	}
}

func SCFFactory(scf func(p Profile) ([]Alternative, error), tb func([]Alternative) (Alternative, error)) func(Profile) (Alternative, error) {
	return func(p Profile) (Alternative, error) {
		bestAlts, errSCF := scf(p)

		if errSCF != nil {
			return Alternative(-1), errSCF
		}

		return tb(bestAlts)
	}
}
