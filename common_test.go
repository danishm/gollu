package gollu

import (
	"testing"
	"time"
)

func TestDaysToExpiry(t *testing.T) {
	ticket := LLLULoginResponseAuthTicket{
		Token:    "gummy token value",
		Expires:  1747496939,
		Duration: 15552000000,
	}

	from := time.Date(2024, time.November, 18, 17, 00, 00, 00, time.UTC)
	days := ticket.DaysToExpiry(from)
	var expected int32 = 179

	if days != expected {
		t.Errorf("got %d, expected %d", days, expected)
	}
}
