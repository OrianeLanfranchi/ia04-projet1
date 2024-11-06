package comsoc

//import "fmt"

func BordaSWF(p Profile) (count Count, err error) {
	count = make(Count)

	err = checkProfileFromProfile(p)
	if err != nil {
		return nil, err
	}

	for i := range p {
		//fmt.Println("(BordaSWF) - len(p[i]) - ", len(p[i]))
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
	if err != nil {
		return nil, err
	}

	return maxCount(count), err
}
