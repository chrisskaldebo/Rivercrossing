package state //Package for testing the representation of the state of the river world

import "testing"

func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn hs ---\\ \\__/ _________________/---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
func TestPutInBoat(t *testing.T) {
	wanted := "kylling"
	state := PutInBoat()
	if state != wanted {
		t.Errorf("Feil, fikk #{state}, ønsket #{wanted}")
	}
}
func TestCrossRiver(t *testing.T) {
	wanted := "west"
	state := CrossRiver()
	if state != wanted {
		t.Errorf("Feil, fikk #{state}, ønsket #{wanted}.")
	}
}
func TestPutInBoat2(t *testing.T) {
	wanted := "HS"
	state := PutInBoat2()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
func TestViewState2(t *testing.T) {
	wanted := "[Rev Korn            ---V \\_Kylling HS_/ _____________________Ø---]"
	state := ViewState2()
	if state != wanted {
		t.Errorf("feil, fikk #{state}, ønsket #{wanted}.")
	}
}
