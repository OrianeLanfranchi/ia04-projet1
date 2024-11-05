package comsoc

func STVSWF(p Profile) (count Count, err error) {
	err = checkProfileFromProfile(p)

	if err != nil {
		return nil, err
	}
	//initialisation du décompte
	count = make(Count)
	for _, alt := range p[0] {
		count[alt] = 0
	}

	//décompte des votes pour chaque tour
	for round := range p[0] {
		// résultats du tour courant
		resultRound := make(Count)

		//décompte pour le tour courant
		for _, pref := range p {
			resultRound[pref[0]]++
		}

		worstAlts := minCount(resultRound)
		resultRound[worstAlts[0]] = round

		//on retire la pire alternative des préférences
		for i, pref := range p {
			pref[rank(worstAlts[0], pref)] = pref[len(pref)-1]
			pref = pref[:len(pref)-1]
			p[i] = pref
		}
	}
	return count, nil
}

func STVSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := STVSWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), err
}
