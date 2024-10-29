package comsoc

//import "fmt"

func MajoritySWF(p Profile) (count Count, err error) {
	err = checkProfileFromProfile(p)
	if err != nil {
		return nil, err
	}
	//initialisation du décompte
	count = make(Count)
	for _, alt := range p[0] {
		count[alt] = 0
	}
	//décompte des votes
	for _, pref := range p {
		count[pref[0]]++
	}
	return
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := MajoritySWF(p)
	return maxCount(count), err
}
