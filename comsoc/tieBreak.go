package comsoc

import (
	"fmt"
	"slices"
)

func TieBreakFactory(orderedAlts []Alternative) func([]Alternative) (Alternative, error) {
	return func(bestAlts []Alternative) (Alternative, error) {
		bestAlt := bestAlts[0]
		for _, alt := range bestAlts[1:] {
			if isPref(alt, bestAlt, orderedAlts) {
				bestAlt = alt
			}
		}
		fmt.Println("Debug - tieBreak - bestAlt : ", bestAlt)
		return bestAlt, nil
	}
}

func SWFFactory(swf func(p Profile) (Count, error), tb func([]Alternative) (Alternative, error)) func(Profile) ([]Alternative, error) {
	return func(p Profile) ([]Alternative, error) {
		count, errSWF := swf(p)
		fmt.Println("Debug - count SWF : ", count)

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
		fmt.Println("Debug - sortedAlts SWF : ", sortedAlts)
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

func SCFOptionFactory(scf func(p Profile, thresholds []int) ([]Alternative, error), tb func([]Alternative) (Alternative, error)) func(Profile, []int) (Alternative, error) {
	return func(p Profile, thresholds []int) (Alternative, error) {
		bestAlts, errSCF := scf(p, thresholds)

		if errSCF != nil {
			return Alternative(-1), errSCF
		}

		return tb(bestAlts)
	}
}

func SWFOptionFactory(swf func(p Profile, thresholds []int) (Count, error), tb func([]Alternative) (Alternative, error)) func(Profile, []int) ([]Alternative, error) {
	return func(p Profile, thresholds []int) ([]Alternative, error) {
		count, errSWF := swf(p, thresholds)

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
