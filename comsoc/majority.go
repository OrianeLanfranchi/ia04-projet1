package comsoc

//import "fmt"

func MajoritySWF(p Profile) (count Count, err error) {
	count = make(Count)

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

	for i := range p {
		err = checkProfile(p[i], alts)
		if err != nil {
			return nil, err
		}
		count[Alternative(p[i][0])] += 1
	}

	return count, nil
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := MajoritySWF(p)
	return maxCount(count), err
}
