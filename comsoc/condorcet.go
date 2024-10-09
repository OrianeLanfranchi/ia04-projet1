package comsoc

func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	//fmt.Println("(CondorcetWinner) - entered")
	count := make(Count)

	err = checkProfileFromProfile(p)
	if err != nil {
		return nil, err
	}

	//on suppose que le premier élément du profil est correctement formé
	var alts = p[0]

	//fmt.Println("(MajoritySWF) - (p[0]) - ", p[0])
	for alt := range p[0] {
		count[Alternative(alt)] = 0
		//fmt.Println("(MajoritySWF) - (count[Alternative(alt)]) - ", count[Alternative(alt)], " - (Alternative(alt)) - ", Alternative(alt))
	}

	//on compare les alternatives deux à deux
	for _, i := range alts {
		victories := 0

		for _, j := range alts {

			if i != j {
				//fmt.Println(i, " VS ", j, " = ")
				iPreffered := 0
				jPreffered := 0

				//on parcourt le profil et on compte le nombre de fois où i est préféré à j
				for _, rankedAlts := range p {
					if isPref(i, j, rankedAlts) {
						iPreffered += 1
					} else {
						jPreffered += 1
					}
				}

				if iPreffered > jPreffered {
					//fmt.Println(i, " wins")
					victories += 1
				}

			}
		}

		if victories == len(alts)-1 { //victoire contre toutes les alternatives sauf elle-même
			//fmt.Println(i, " is winner with : ", victories, "victories")
			bestAlts = append(bestAlts, i)
			return bestAlts, nil
		}
	}

	return nil, nil
}
