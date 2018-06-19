package socialSecurityNumber

import (
	"time"
	"github.com/pkg/errors"
	"strconv"
)

type SocialSecurityNumber struct {
	Century     int
	Year        int
	Month       time.Month
	Day         int
	BirthNumber string
	Control     int
}

func Create(birthDate time.Time, birthNumber string, control int) SocialSecurityNumber {
	return SocialSecurityNumber{
		Century:     birthDate.Year() / 100,
		Year:        birthDate.Year() % 100,
		Month:       birthDate.Month(),
		Day:         birthDate.Day(),
		BirthNumber: birthNumber,
		Control:     control,
	}
}

func FromString(ssn string) (SocialSecurityNumber, error) {
	var birthDate time.Time
	var birthNumber string
	if len(ssn) == 10 {
		date, err := time.Parse("060102", ssn[0:6])
		if err != nil {
			return SocialSecurityNumber{}, err
		}

		birthDate = date
		birthNumber = ssn[7:10]

	} else if len(ssn) == 12 {
		date, err := time.Parse("20060102", ssn[0:8])
		if err != nil {
			return SocialSecurityNumber{}, err
		}

		birthDate = date
		birthNumber = ssn[9:12]
	} else {
		return SocialSecurityNumber{}, errors.New("Social security number must be either 10 or 12 digits")
	}

	control, err := strconv.Atoi(string(ssn[len(ssn) - 1]))
	if err != nil {
		return SocialSecurityNumber{}, err
	}

	return Create(birthDate, birthNumber, control), nil
}

func (ssn *SocialSecurityNumber) Equal(other SocialSecurityNumber) bool {
	return ssn.Century == other.Century &&
		ssn.Year == other.Year &&
		ssn.Month == other.Month &&
		ssn.Day == other.Day &&
		ssn.BirthNumber == other.BirthNumber &&
		ssn.Control == other.Control
}
