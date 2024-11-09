package comsoc

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) int {
	for i := range prefs {
		if prefs[i] == alt {
			return i
		}
	}

	return -1
}

// renvoie vrai ssi alt1 est préférée à alt2
func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	return rank(alt1, prefs) < rank(alt2, prefs)
}

// renvoie les meilleures alternatives pour un décompte donné
func maxCount(count Count) (bestAlts []Alternative) {
	var max float64 = math.Inf(-1)

	for alt, c := range count {
		if float64(c) > max {
			max = float64(c)
			bestAlts = []Alternative{alt}
		} else if float64(c) == max {
			bestAlts = append(bestAlts, alt)
		}
	}
	return bestAlts
}

// renvoie les pires alternatives pour un décompte donné
func minCount(count Count) (worstAlts []Alternative) {
	var min float64 = math.Inf(1)

	for alt, c := range count {
		if float64(c) < min {
			min = float64(c)
			worstAlts = []Alternative{alt}
		} else if float64(c) == min {
			worstAlts = append(worstAlts, alt)
		}
	}
	return worstAlts
}

// vérifie les préférences d'un agent, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois
func CheckProfile(prefs []Alternative, alts []Alternative) error {
	//vérifie que profil complet
	if len(prefs) != len(alts) {
		fmt.Println("(checkProfile) - len(prefs) - ", len(prefs))
		fmt.Println("(checkProfile) - len(alts) - ", len(alts))
		return errors.New("(checkProfile) - prefs et alts n'ont pas la même taille")
	}

	fmt.Println("(checkProfile) - (prefs) - ", prefs)
	fmt.Println("(checkProfile) - (alts) - ", alts)

	for _, pref := range prefs {
		if !slices.Contains(alts, pref) {
			fmt.Println("(checkProfile) - (pref) - ", pref)
			return errors.New("(checkProfile) - Il manque une préférence")
		}
	}

	//il faudrait aussi vérifier qu'une alternative n'est pas présente en double
	isUnique := CheckUniquePreferences(prefs)

	if !isUnique {
		fmt.Println("(checkProfile) - (prefs) - ", prefs)
		return errors.New("(checkProfile) - La même alternative est présente plusieurs fois dans les préférences")
	}

	return nil
}

func CheckUniquePreferences(alts []Alternative) bool {
	unique := make(map[Alternative]int, len(alts))

	for _, alt := range alts {
		unique[alt] = 0
	}

	for _, alt := range alts {
		unique[alt]++
		if unique[alt] > 1 {
			return false
		}
	}

	return true
}

func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	for _, pref := range prefs {
		return CheckProfile(pref, alts)
	}
	return nil
}

func checkProfileFromProfile(prof Profile) (err error) {
	alts := make([]Alternative, 0)
	for i := 1; i <= len(prof[0]); i++ {
		alts = append(alts, Alternative(i))
	}
	err = checkProfileAlternative(prof, alts)
	return
}
