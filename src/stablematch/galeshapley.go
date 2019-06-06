package main

import (
	"stablematch/internal"
)

func main(){
	men := internal.MatchList{
		internal.NewMatchPerson("A", []int{1,0,2}),
		internal.NewMatchPerson("B", []int{1,2,0}),
		internal.NewMatchPerson("C", []int{2,1,0}),
	}
	women := internal.MatchList{
		internal.NewMatchPerson("a", []int{2,1,0}),
		internal.NewMatchPerson("b", []int{1,0,2}),
		internal.NewMatchPerson("c", []int{1,2,0}),
	}
	galeshapley(men, women)
	men.PrintMatch(women)
}

func galeshapley(men, women internal.MatchList){
	index := men.GetFreePerson()
	for index>=0 {
		man := men[index]
		woman := women[man.NextInvite]
		if woman.IsFree() {
			woman.CurrentPartner = index
			man.CurrentPartner = man.NextInvite
		}else if !woman.IsFree() && woman.LeftIsPref(index, woman.CurrentPartner) {
			oldPartner := men[woman.CurrentPartner]
			oldPartner.CurrentPartner = -1
			woman.CurrentPartner = index
			man.CurrentPartner = man.NextInvite
		}
		man.NextInvite++
		index = men.GetFreePerson()
	}
}