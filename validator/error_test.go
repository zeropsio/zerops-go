package validator

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorList_With(t *testing.T) {

	call := func(getNil bool) error {
		var list ErrorList
		if !getNil {
			list = list.With(fmt.Errorf("error"))
		}
		return list.GetError()
	}

	if call(true) != nil {
		t.Errorf("should be nil")
	}
	if call(false) == nil {
		t.Errorf("should not be nil")
	}
}

func TestErrorListWrapping(t *testing.T) {
	err := do()
	if err == nil {
		t.Error("should not be nil")
	}
}

func do() error {
	err := dbInsert()
	if err != nil {
		return err
	}
	return nil
}

func dbInsert() error {
	err := validation()
	if err != nil {
		return err
	}
	return nil
}

func validation() error {
	var errorList ErrorList
	errorList = errorList.WithPrefix("xxx.", errors.New("aaa"))
	errorList = errorList.WithPrefix("yyy.", errors.New("bbb"))
	return errorList.GetError()
}
