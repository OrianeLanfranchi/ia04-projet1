package comsoc

func CopelandSWF(p Profile) (count Count, err error) {
	count = make(Count)
	err = checkProfileFromProfile(p)

	if err != nil {
		return nil, err
	}
	//initialisation du décompte
	count = make(Count)
	for _, alt := range p[0] {
		count[alt] = 0
	}

	return count, nil
}
