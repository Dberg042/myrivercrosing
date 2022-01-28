// Package for testing the represention of the state of the riverworld
package state

import "testing"

// Test function implemented in line with the Golang's testing tool
func TestViewState(t *testing.T) {
	wanted := "[  ---\\    _________________\\  / / human sheep wolf corn]"
	//Parameter 1 should take int 0-3
	//Parameter 1 == 0 Human puts wolf on the boat
	//Parameter 2 == 1 Human puts corn on the boat
	//Parameter 1 == 2 Human crosses alone the river
	//Parameter 1 == 3 Human puts sheep on the boat
	//If you find right condition, human will put of the item on the left
	//and carry to right. 
	//For 1st item enter 0 for paremeter 2
	//For 2nd item enter 1 for parameter 2
	//Parameter 3 is third choice, if you choose right human act logicly and
	//finalize the process. You should choose 0 or 1. 
	//Parameter 3 = 0 means human put item which passed first the river and
	//takes back to right side.
	// Parameter 3 = 1 means human turn right side alone.
	state := ViewState(3, 0, 1)
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
