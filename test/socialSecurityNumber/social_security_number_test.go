package socialSecurityNumber

import (
	"testing"
	"subLease/src/server/socialSecurityNumber"
	"time"
	"fmt"
)

func TestFromTenDigitString(t *testing.T) {
	ssnString := "8906011111"

	expectedSsn := socialSecurityNumber.Create(time.Date(1989, time.June, 1, 0, 0, 0, 0, time.Local), "111", 1)
	actualSsn, err := socialSecurityNumber.FromString(ssnString)
	if err != nil {
		t.Error(err.Error())
	}

	if !actualSsn.Equal(expectedSsn) {
		t.Error(fmt.Sprintf("Social secutiry number string was not parsed correctly. Expected: %v but was %v", expectedSsn, actualSsn))
	}
}

func TestFromTwelveDigitString(t *testing.T) {
	ssnString := "198906011111"

	expectedSsn := socialSecurityNumber.Create(time.Date(1989, time.June, 1, 0, 0, 0, 0, time.Local), "111", 1)
	actualSsn, err := socialSecurityNumber.FromString(ssnString)
	if err != nil {
		t.Error(err.Error())
	}

	if !actualSsn.Equal(expectedSsn) {
		t.Error(fmt.Sprintf("Social secutiry number string was not parsed correctly. Expected: %v but was %v", expectedSsn, actualSsn))
	}
}
