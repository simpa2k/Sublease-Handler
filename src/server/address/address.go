package address

type Address struct {
	Street string
	Number int
	ZipCode string // The datatype is debatable but the zipcode is not really an integer, despite technically being numeric. It would be overkill to represent it as a separate type.
	City string
}

func Create(street string, number int, zipCode string, city string) Address {
	return Address {
		Street: street,
		Number: number,
		ZipCode: zipCode,
		City: city,
	}
}

func (a *Address) Equal(other *Address) bool {
	return a.Street == other.Street && a.Number == other.Number && a.ZipCode == other.ZipCode && a.City == other.City
}
