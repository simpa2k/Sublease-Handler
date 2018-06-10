package socialSecurityNumber

import (
	"time"
)

type SocialSecurityNumber struct {
	Century int
	Year int
	Month time.Month
	Day int
	BirthNumber string
	Control int
}

func Create(birthDate time.Time, birthNumber string, control int) SocialSecurityNumber {
	return SocialSecurityNumber {
		Century: birthDate.Year() / 100,
		Year: birthDate.Year() % 100,
		Month: birthDate.Month(),
		Day: birthDate.Day(),
		BirthNumber: birthNumber,
		Control: control,
	}
}

func (ssn *SocialSecurityNumber) Equal(other SocialSecurityNumber) bool {
	return ssn.Century == other.Century &&
		ssn.Year == other.Year &&
		ssn.Month == other.Month &&
		ssn.Day == other.Day &&
		ssn.BirthNumber == other.BirthNumber &&
		ssn.Control == other.Control
}