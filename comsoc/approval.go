package comsoc

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	count = make(Count)

	for _, alt := range p[0] {
		count[alt] = 0
	}

	for index, alt := range p {
		for i := 0; i < thresholds[index]; i++ {
			count[alt[i]] += 1
		}
	}
	return count, nil
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	count, err := ApprovalSWF(p, thresholds)

	if err != nil {
		return nil, err
	}

	return maxCount(count), nil
}
