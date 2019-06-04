package internal

import(
	"errors"
	"fmt"
	"strings"
)

var (
	BOAT_ACTIONS = [5][2]int{
		{1, 0},
		{0, 1},
		{2, 0},
		{0, 2},
		{1, 1},
	}
)

type BankState struct{
	MonkNum		int
	MonsterNum	int
	HasBoat		bool
}

type BothSides [2]BankState

func (bs *BothSides)IsFinal()bool{
	return bs[0].MonkNum==0 && bs[0].MonsterNum==0;
}

func (bs *BothSides)BoatInLeft()bool{
	return bs[0].HasBoat
}

func (bs *BothSides)String()string{
	return fmt.Sprintf("Monk:%v,Monster:%v || Monk:%v,Monster:%v BoatInLeft:%v\n", bs[0].MonkNum, bs[0].MonsterNum, bs[1].MonkNum, bs[1].MonsterNum, bs.BoatInLeft())
}

func (bs *BothSides)IsValid()bool{
	for i := 0; i < 2; i++ {
		if bs[i].MonkNum > 0 && bs[i].MonsterNum > bs[i].MonkNum {
			return false
		}
	}
	return true
}

func (bs *BothSides)Transfer(monkNum, monsterNum int)(*BothSides, error){
	
	if monkNum + monsterNum < 1 {
		return nil, errors.New("boat should haa at least one being")
	}
	if monkNum + monsterNum > 2 {
		return nil, errors.New("boat is overloaded")
	}
	var isLeft bool = bs.BoatInLeft()
	var boatSideState, otherSideState BankState
	if isLeft {
		boatSideState = bs[0]
		otherSideState = bs[1]
	}else{
		otherSideState = bs[0]
		boatSideState = bs[1]
	}
	if boatSideState.MonkNum < monkNum || boatSideState.MonsterNum < monsterNum {
		return nil, errors.New("not enough monk or monster")
	}
	boatSideState.MonkNum -= monkNum
	otherSideState.MonkNum +=  monkNum
	boatSideState.MonsterNum -= monsterNum
	otherSideState.MonsterNum += monsterNum
	boatSideState.HasBoat = false
	otherSideState.HasBoat = true
	var newState *BothSides
	if isLeft {
		newState = &BothSides{
			boatSideState,
			otherSideState,
		}
	}else{
		newState = &BothSides{
			otherSideState,
			boatSideState,
		}
	}
	return newState, nil
}

type StateStack []BothSides

func (stack *StateStack)Len()int{
	return len(*stack)
}

func (stack *StateStack)Push(stateP *BothSides){
	*stack = append(*stack, *stateP)
}

func (stack *StateStack)Pop()*BothSides{
	if stack.Len() < 1{
		return nil 
	}
	state := (*stack)[stack.Len()-1]
	*stack = (*stack)[0:stack.Len()-1]
	return &state
}

func (stack *StateStack)Back()*BothSides{
	return &(*stack)[stack.Len()-1]
}

func (stack *StateStack)Front()*BothSides{
	if stack.Len() < 1 {
		return nil
	}
	return &(*stack)[0]
}

func (stack *StateStack)Index(stateP *BothSides)int{
	if stack.Len() < 1{
		return -1 
	}
	for i := 0; i < stack.Len(); i++ {
		if (*stack)[i].String() == stateP.String() {
			return i
		}
	}
	return -1
}

func (stack *StateStack)PrintPath(){
	fmt.Println(strings.Repeat("=", 10), "path begin", strings.Repeat("=", 10)) 
	for i:=0; i < stack.Len(); i++ {
		fmt.Println((*stack)[i].String())
	}
	fmt.Println(strings.Repeat("*", 20))
}