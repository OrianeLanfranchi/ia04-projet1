package comsoc

import (
	"errors"
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
	var max int = 0

	for alt, c := range count {
		if c > max {
			max = c
			bestAlts = []Alternative{alt}
		} else if c == max {
			bestAlts = append(bestAlts, alt)
		}
	}
	return bestAlts
}

// vérifie les préférences d'un agent, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois
func checkProfile(prefs []Alternative, alts []Alternative) error {
	//vérifie que profil complet
	if len(prefs) != len(alts) {
		return errors.New("(checkProfile) - prefs et alts n'ont pas la même taille")
	}

	//fmt.Println("(checkProfile) - (prefs) - ", prefs)
	//fmt.Println("(checkProfile) - (alts) - ", alts)

	for _, pref := range prefs {
		if !slices.Contains(alts, pref) {
			//fmt.Println("(checkProfile) - (pref) - ", pref)
			return errors.New("(checkProfile) - Il manque une préférence")
		}
	}

	return nil
}

func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	for _, pref := range prefs {
		return checkProfile(pref, alts)
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
