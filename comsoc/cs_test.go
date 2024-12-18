// version 2.0.0

package comsoc

import "testing"

func TestRank(t *testing.T) {
	prefs := []Alternative{4, 3, 2, 1, 0}
	var alt Alternative = 0
	x := rank(alt, prefs)

	if x != 4 {
		t.Errorf("Le rang est %d et devrait être 4", x)
	}
}

func TestIsPref(t *testing.T) {
	prefs := []Alternative{4, 3, 2, 1, 0}

	if !isPref(3, 1, prefs) {
		t.Errorf("3 devrait être préféré à 1")
	}
}

func TestMaxCount(t *testing.T) {
	count := map[Alternative]int{

		0: 4,
		1: 4,
		2: 2,
		3: 2,
		4: 0,
	}
	result := maxCount(count)

	if result[0] != 0 || result[1] != 1 {
		t.Errorf("les plus haut sont %d et %d au lieu de 0 et 1", result[0], result[1])
	}
}

func TestCheckProfile(t *testing.T) {
	prefs := Preferences{1, 3, 2}
	alts := []Alternative{1, 2, 3}

	err := CheckProfile(prefs, alts)
	if err != nil {
		t.Error(err.Error())
	}

}

func TestCheckProfileAlterative(t *testing.T) {
	prefs := Profile{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}
	alts := []Alternative{1, 2, 3}

	err := checkProfileAlternative(prefs, alts)
	if err != nil {
		t.Error(err.Error())
	}

}

func TestBordaSWF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res, _ := BordaSWF(prefs)

	if res[1] != 4 {
		t.Errorf("error, result for 1 should be 4, %d computed", res[1])
	}
	if res[2] != 3 {
		t.Errorf("error, result for 2 should be 3, %d computed", res[2])
	}
	if res[3] != 2 {
		t.Errorf("error, result for 3 should be 2, %d computed", res[3])
	}
}

func TestBordaSCF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res, err := BordaSCF(prefs)

	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 || res[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}
}

func TestMajoritySWF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res, _ := MajoritySWF(prefs)

	if res[1] != 2 {
		t.Errorf("error, result for 1 should be 2, %d computed", res[1])
	}
	if res[2] != 0 {
		t.Errorf("error, result for 2 should be 0, %d computed", res[2])
	}
	if res[3] != 1 {
		t.Errorf("error, result for 3 should be 1, %d computed", res[3])
	}
}

func TestMajoritySCF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	res, err := MajoritySCF(prefs)

	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 || res[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}
}

func TestApprovalSWF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{2, 3, 1},
	}
	thresholds := []int{2, 1, 2}

	res, _ := ApprovalSWF(prefs, thresholds)

	if res[1] != 2 {
		t.Errorf("error, result for 1 should be 2, %d computed", res[1])
	}
	if res[2] != 2 {
		t.Errorf("error, result for 2 should be 2, %d computed", res[2])
	}
	if res[3] != 1 {
		t.Errorf("error, result for 3 should be 1, %d computed", res[3])
	}
}

func TestApprovalSCF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 3, 2},
		{1, 2, 3},
		{2, 1, 3},
	}
	thresholds := []int{2, 1, 2}

	res, err := ApprovalSCF(prefs, thresholds)

	if err != nil {
		t.Error(err)
	}
	if len(res) != 1 || res[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}
}

func TestCondorcetWinner(t *testing.T) {
	prefs1 := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	prefs2 := [][]Alternative{
		{1, 2, 3},
		{2, 3, 1},
		{3, 1, 2},
	}

	res1, _ := CondorcetWinner(prefs1)
	res2, _ := CondorcetWinner(prefs2)

	if len(res1) == 0 || res1[0] != 1 {
		t.Errorf("error, 1 should be the only best alternative for prefs1")
	}
	if len(res2) != 0 {
		t.Errorf("no best alternative for prefs2")
	}
}

func TestSTVSWF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{2, 3, 1},
	}

	res, _ := STVSWF(prefs)

	if res[1] != 3 {
		t.Errorf("error, result for 1 should be 1, %d computed", res[1])
	}
	if res[2] != 2 {
		t.Errorf("error, result for 2 should be 2, %d computed", res[2])
	}
	if res[3] != 1 {
		t.Errorf("error, result for 3 should be 3, %d computed", res[3])
	}

	res, _ = STVSWF(prefs)
	if res[1] != 3 {
		t.Errorf("error, result for 1 should be 1, %d computed", res[1])
	}
	if res[2] != 2 {
		t.Errorf("error, result for 2 should be 2, %d computed", res[2])
	}
	if res[3] != 1 {
		t.Errorf("error, result for 3 should be 3, %d computed", res[3])
	}
}

func TestSTVSCF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 3, 2},
		{1, 2, 3},
		{2, 1, 3},
	}
	res, err := STVSCF(prefs)

	if err != nil {
		t.Error(err)
	}
	if len(res) != 1 || res[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}
}

func TestCopelanSFC(t *testing.T) {
	prefs := [][]Alternative{
		{1, 3, 2},
		{1, 2, 3},
		{2, 1, 3},
	}
	res, err := CopelandSCF(prefs)

	if err != nil {
		t.Error(err)
	}
	if len(res) != 1 || res[0] != 1 {
		t.Errorf("error, 1 should be the only best Alternative")
	}
}

func TestCopelandSWF(t *testing.T) {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 3, 2},
		{2, 3, 1},
	}

	res, _ := CopelandSWF(prefs)

	if res[1] != 2 {
		t.Errorf("error, result for 1 should be 1, %d computed", res[1])
	}
	if res[2] != 0 {
		t.Errorf("error, result for 2 should be 2, %d computed", res[2])
	}
	if res[3] != -2 {
		t.Errorf("error, result for 3 should be 3, %d computed", res[3])
	}
}

func TestUniqueAlts(t *testing.T) {
	pref1 := []Alternative{1, 2, 3}
	pref2 := []Alternative{1, 2, 4, 4, 5}
	res1 := CheckUniquePreferences(pref1)
	res2 := CheckUniquePreferences(pref2)

	if res1 != true {
		t.Errorf("error, result for res1 should be True, False computed")
	}

	if res2 != false {
		t.Errorf("error, result for res1 should be False, True computed")
	}
}
