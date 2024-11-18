package gollu

import (
	"strings"
	"time"
)

const (
	LLUTimestampFormat = "1/2/2006 3:04:05 PM"
)

type LLLULoginResponseAuthTicket struct {
	Token    string
	Expires  int64
	Duration int64
}

func (ticket *LLLULoginResponseAuthTicket) DaysToExpiry() int32 {
	now := time.Now()
	ticketExpiry := time.Unix(ticket.Expires, 0)
	difference := ticketExpiry.Sub(now)
	return rune(difference.Hours() / 24)
}

type LLUTimestamp time.Time

// UnmarshalJSON implements a custom date unmarshaler for the time format used by LibreLinkup
func (lts *LLUTimestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(LLUTimestampFormat, s)
	if err != nil {
		return err
	}
	*lts = LLUTimestamp(t)
	return nil
}

func (lts *LLUTimestamp) String() string {
	t := time.Time(*lts)
	return t.Format(LLUTimestampFormat)
}
