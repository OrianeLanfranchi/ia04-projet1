package comsoc

import "fmt"

func MajoritySWF(p Profile) (count Count, err error) {
	err = checkProfileFromProfile(p)

	if err != nil {
		return nil, err
	}
	fmt.Println("(STVSWF) - p[0] (init) - ", p)

	//initialisation du décompte
	count = make(Count)
	for _, alt := range p[0] {
		count[alt] = 0
	}

	fmt.Println("(MajoritySWF) - count (init) - ", count)

	//décompte des votes
	for _, pref := range p {
		count[pref[0]]++
	}
	return count, nil
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := MajoritySWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), err
}
