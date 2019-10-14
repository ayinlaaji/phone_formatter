package csv

import (
	"fmt"
	"testing"
)

func TestRemoveChar(t *testing.T) {

	samp := "+234 80 3472 849 6"
	expRes := "+2348034728496"
	res := RemoveChar(samp, " ")

	if expRes != res {
		t.Error("Something is not right")
	}
}

func TestCheckPlus(t *testing.T) {
	samp := "+2348034923892"
	hasPlus := CheckPlus(samp)

	if hasPlus == false {
		t.Error("Error")
	}

	samp2 := "2348034923892"

	hasPlus2 := CheckPlus(samp2)

	if hasPlus2 == true {
		t.Error("Error")
	}

}

func TestAddZip(t *testing.T) {

	samp := "08034923892"
	exp := "+2348034923892"
	res, _ := AddZip(samp, "NG")

	if exp != res {
		t.Error("Zip setting not correct")
	}

}

func TestRemoveDelimiters(t *testing.T) {
	samp := "08034923892 ::: 08033445568"
	exp := []string{"08034923892", "08033445568"}

	res := RemoveDelimiter(samp, ":::")

	for ind, item := range res {
		if item != exp[ind] {
			t.Error("Delimiter not removed")
			return
		}
	}

}

func TestContactFormat(t *testing.T) {
	samp := "803492 3892"
	exp := "+2348034923892"
	res := ContactFormat(samp)

	fmt.Println(exp, res)

	if res != exp {
		t.Fail()
	}
}
