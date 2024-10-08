package comsoc

//import "fmt"

func BordaSWF(p Profile) (count Count, err error) {
	count = make(Count)

	//on suppose que le premier élément du profil est correctement formé
	var alts = p[0]
	for alt := range p[0] {
		count[Alternative(alt)] = 0
	}

	for i := range p {
		//fmt.Println("(BordaSWF) - len(p[i]) - ", len(p[i]))
		err = checkProfile(p[i], alts)
		if err != nil {
			break
		}
		for j := range p[i] {
			count[Alternative(p[i][j])] += len(p[i]) - j - 1
			//fmt.Println("(BordaSWF) - Alternative(p[i][0]) - ", Alternative(p[i][0]))
			//fmt.Println("(BordaSWF) - len(p[i]) - j - ", len(p[i])-j)
		}
	}

	return count, err
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := BordaSWF(p)
	return maxCount(count), err
}
