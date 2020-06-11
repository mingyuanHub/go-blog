package core

import (
	"strconv"
	"fmt"
)
func StrToInt(s string) (i int){
	var err error
	i, err = strconv.Atoi(s)
	if (err != nil) {
		fmt.Println(err)
	}
	return
}