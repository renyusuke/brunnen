package sortx

import (
	"fmt"
	"ph-gitlab.vipsroom.net/basis/public/helper/randx"
	"strconv"
	"testing"
)

type P struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func TestSortStruct(t *testing.T) {
	var ps []*P

	for i := 0; i < 5; i++ {
		p := &P{
			Age:  int(randx.GetRandNum(2, 0)),
			Name: strconv.FormatInt(randx.GetRandNum(2, 0), 10),
		}
		ps = append(ps, p)
	}

	sortStruct := SortStruct[*P](ps, "Name", 1)
	fmt.Println(sortStruct)
}
