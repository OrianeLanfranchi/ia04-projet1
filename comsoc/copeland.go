package comsoc

func CopelandSWF(p Profile) (count Count, err error) {
	err = checkProfileFromProfile(p)

	if err != nil {
		return nil, err
	}
	//initialisation du décompte
	count = make(Count)
	for _, alt := range p[0] {
		count[alt] = 0
	}

	for alt1 := 1; alt1 < len(p[0])+1; alt1++ {
		for alt2 := alt1 + 1; alt2 < len(p[0])+1; alt2++ {
			alt1Preffered, alt2Preffered := 0, 0

			for _, pref := range p {
				if isPref(Alternative(alt1), Alternative(alt2), pref) {
					alt1Preffered++
				} else {
					alt2Preffered++
				}
			}

			if alt1Preffered > alt2Preffered {
				count[Alternative(alt1)]++
				count[Alternative(alt2)]--
			} else if alt1Preffered < alt2Preffered {
				count[Alternative(alt1)]--
				count[Alternative(alt2)]++
			}

		}
	}

	return count, nil
}

func CopelandSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := CopelandSWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), err
}
