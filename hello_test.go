package starter_project_coderetreat

import "testing"

func TestPass(t *testing.T) {

}

func TestFail(t *testing.T) {

	if !(true == false) {
		t.Fail()
	}
}

