package gollu

import (
	"testing"
)

func TestDaysToExpiry(t *testing.T) {
	ticket := LLLULoginResponseAuthTicket{
		Token:    "gummy token value",
		Expires:  1747496939,
		Duration: 15552000000,
	}

	days := ticket.DaysToExpiry()
	var expected int32 = 179

	if days != expected {
		t.Errorf("got %d, expected %d", days, expected)
	}
}
