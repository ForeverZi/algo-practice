package main

import(
	"flag"
	"monkandmonster/internal"
)

var (
	initMonkNum		int
	initMonsterNum	int
)

func init(){
	flag.IntVar(&initMonkNum, "monk", 3, "please enter initial monk number")
	flag.IntVar(&initMonsterNum, "monster", 3, "please enter initial monster number")
}

func main(){
	flag.Parse()
	leftState := internal.BankState{
		MonkNum: initMonkNum,
		MonsterNum: initMonsterNum,
		HasBoat: true,
	}
	var initState internal.BothSides
	initState[0] = leftState
	stack := internal.StateStack{initState}
	searchState(stack)
}

func searchState(stack internal.StateStack){
	back := stack.Back()
	if back.IsFinal() {
		stack.PrintPath()
		return
	}
	for _, action := range internal.BOAT_ACTIONS {
		newState := changeState(back, action)
		if newState != nil && stack.Index(newState)<0 {
			stack.Push(newState)
			searchState(stack)
			stack.Pop()
		}
	}
}

func changeState(currentState *internal.BothSides, actionParam [2]int)*internal.BothSides{
	newState, err := currentState.Transfer(actionParam[0], actionParam[1])
	if err != nil {
		return nil
	}
	if !newState.IsValid() {
		return nil
	}
	return newState
}