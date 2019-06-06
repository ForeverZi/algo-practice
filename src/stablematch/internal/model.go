package internal

import (
	"strings"
	"fmt"
)

type MatchPerson struct{
	Name			string
	CurrentPartner	int
	NextInvite		int
	PrefList		[]int
}

func (person *MatchPerson)IsFree()bool{
	return person.CurrentPartner < 0
}

func (person *MatchPerson)LeftIsPref(left, right int)bool{
	for _, i := range person.PrefList {
		if i == left {
			return true
		}
		if i == right {
			return false
		}
	}
	return false
}

func NewMatchPerson(name string, prefList []int)*MatchPerson{
	person := MatchPerson{
		Name: name,
		CurrentPartner: -1,
		NextInvite: 0,
		PrefList: prefList,
	}
	return &person
}

type MatchList	[]*MatchPerson

func (list *MatchList)GetFreePerson()int{
	for i:=0; i < len(*list); i++ {
		if (*list)[i].CurrentPartner < 0 {
			return i
		}
	}
	return -1;
}

func (list *MatchList)PrintMatch(otherGroup MatchList){
	fmt.Println(strings.Repeat("*", 40))
	for i, personP := range *list {
		if !personP.IsFree() {
			fmt.Printf("Match %v: [%v, %v]\n", i, personP.Name, otherGroup[personP.CurrentPartner].Name)
		}
	}
	fmt.Println(strings.Repeat("*", 40))
}