package main

import (
	"flag"
	"strings"
	"strconv"
	"errors"
	"fmt"
	// "encoding/json"
)

type Bucket struct{
	Cap		int
	Vol		int
}

type BucketGroup []Bucket

func (group BucketGroup)HashKey()string{
	if len(group) < 1 {
		return ""
	}
	key := strconv.Itoa(group[0].Vol)
	for i := 1; i < len(group); i++ {
		key += "-" + strconv.Itoa(group[i].Vol)
	}
	return key
}

func (group *BucketGroup)String()string{
	var b strings.Builder
	for i,bucket := range *group {
		if i > 0 {
			b.WriteString(";")
		}
		b.WriteString(fmt.Sprintf("%v,%v", bucket.Cap, bucket.Vol))
	}
	return b.String()
}

func (group *BucketGroup)Set(str string)error{
	pairs := strings.Split(str, ";")
	for _, pairStr := range pairs {
		pair := strings.Split(pairStr, ",")
		if len(pair) != 2{
			return errors.New("param not valid, please enter group like '8,8;5,0;3,0'")
		}
		cap, err := strconv.Atoi(pair[0])
		if err != nil {
			return err
		}
		vol, err := strconv.Atoi(pair[1])
		if err != nil {
			return nil
		}
		*group = append(*group, Bucket{
			Cap: cap,
			Vol: vol,
		})
	}
	return nil
}

var (
	initGroup BucketGroup
	finalGroup BucketGroup
)

func init(){
	flag.Var(&initGroup, "init", "please set init state")
	flag.Var(&finalGroup, "final", "please set final state")
}

func main(){
	flag.Parse()
	records := make(map[string]string)
	records[initGroup.HashKey()] = ""
	search(initGroup, records)
	// bytes, _ := json.MarshalIndent(records, "", "  ")
	// fmt.Println(string(bytes))
	fpath, err := getFinalPath(records)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(strings.Join(fpath, "->"))
	}
}

func getFinalPath(records map[string]string)([]string, error){
	last, ok := records[finalGroup.HashKey()]
	if !ok {
		return nil, errors.New("can't find final state")
	}
	fpath := []string{finalGroup.HashKey()}
	for last != "" {
		llast, ok := records[last]
		if !ok {
			return nil, errors.New("lost state:"+last)
		}
		fpath = append(fpath, last)
		last = llast
	}
	for i,j := 0, len(fpath)-1; i < j; i,j = i+1,j-1 {
		fpath[i], fpath[j] = fpath[j], fpath[i]
	}
	return fpath, nil
}

func search(group BucketGroup, records map[string]string){
	for from := 0; from < len(group); from++ {
		for to := 0; to < len(group); to++ {
			pourVol, valid := checkAction(group, from, to)
			if valid {
				nextGroup := pour(group, from, to, pourVol)
				key := nextGroup.HashKey()
				// fmt.Println(group.HashKey(), "->", key)
				_, visited := records[key]
				if !visited {
					records[key] = group.HashKey()
					if checkFinal(nextGroup) {
						return
					}
					search(nextGroup, records)
				}
			}
		} 
	}
}

func checkFinal(group BucketGroup)bool{
	return group.HashKey() == finalGroup.HashKey()
}

func pour(group BucketGroup, from, to, pourVol int)BucketGroup{
	var next BucketGroup
	for i, bucket := range group {
		if from == i {
			bucket.Vol -= pourVol
		}
		if to == i {
			bucket.Vol += pourVol
		}
		next = append(next, bucket)
	}
	return next
}

func checkAction(group BucketGroup, from, to int)(pourVol int, valid bool){
	if from == to {
		return 0, false
	}
	if from < 0 || from >= len(group) {
		return 0, false
	}
	if to < 0 || to >= len(group) {
		return 0, false
	}
	if group[from].Vol == 0 {
		return 0, false
	}
	toRemain := group[to].Cap - group[to].Vol
	if toRemain == 0 {
		return 0, false
	}
	if group[from].Vol >= toRemain {
		return toRemain, true
	}else{
		return group[from].Vol, true
	}
}