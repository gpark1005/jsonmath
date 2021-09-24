package data

import (
	"errors"
	"github.com/google/uuid"
	"strconv"
)

type OuterData struct {
	Id            string
	Nums          []int
	NumberOne     int
	NumberTwo     int
	Sum           int
	StringVersion InnerData
}

type InnerData struct {
	NumOne string
	NumTwo string
	Answer string
}

func (od *OuterData) DoConversion() error {

	if len(od.Nums) < 2 {
		return errors.New("sum field must have 2 numbers")
	}
	od.NumberOne = od.Nums[0]
	od.NumberTwo = od.Nums[1]
	od.Sum = od.NumberOne + od.NumberTwo

	od.StringVersion.NumOne = strconv.Itoa(od.NumberOne)
	od.StringVersion.NumTwo = strconv.Itoa(od.NumberTwo)
	od.StringVersion.Answer = strconv.Itoa(od.Sum)
	return nil
}

func (od *OuterData) SetId() {
	od.Id = uuid.New().String()
}
