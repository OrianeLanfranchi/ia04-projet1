package comsoc

import "fmt"

func STVSWF(p Profile) (count Count, err error) {
	pCopy := make(Profile, len(p))

	for i := range p {
		pCopy[i] = append(pCopy[i], p[i]...)
	}

	fmt.Println("(STVSWF) - p (beginning) - ", p)
	fmt.Println("(STVSWF) - pCopy (beginning) - ", pCopy)

	err = checkProfileFromProfile(pCopy)

	if err != nil {
		return nil, err
	}
	//initialisation du décompte
	count = make(Count)
	for _, alt := range pCopy[0] {
		count[alt] = 0
	}

	fmt.Println("(STVSWF) - count (init) - ", count)

	//décompte des votes pour chaque tour
	for round := range pCopy[0] {
		// résultats du tour courant
		resultRound := make(Count)
		for _, alt := range pCopy[0] {
			resultRound[alt] = 0
		}

		//décompte pour le tour courant
		for _, pref := range pCopy {
			resultRound[pref[0]]++
		}
		fmt.Println("(STVSWF) - resultRound : ", resultRound)

		worstAlts := minCount(resultRound)
		fmt.Println("(STVSWF) - worstAlts : ", worstAlts)
		count[worstAlts[0]] = round + 1

		//on retire la pire alternative des préférences
		for i, pref := range pCopy {
			pref[rank(worstAlts[0], pref)] = pref[len(pref)-1]
			pref = pref[:len(pref)-1]
			pCopy[i] = pref
		}
	}
	fmt.Println(count)
	fmt.Println("(STVSWF) - p (end) - ", p)
	fmt.Println("(STVSWF) - pCopy (end) - ", pCopy)
	return count, nil
}

func STVSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := STVSWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), err
}
