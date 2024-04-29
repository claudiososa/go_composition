package types

type Profile struct {
	Name    string
	Surname string
	Address string
	Phone   string
}

func (profile *Profile) GetName() string {
	return profile.Name
}

func (profile *Profile) GetSurname() string {
	return profile.Surname
}

func (profile *Profile) GetAddress() string {
	return profile.Address
}

func (profile *Profile) GetPhone() string {
	return profile.Phone
}

func (profile *Profile) SetName(Name string) *Profile {
	profile.Name = Name
	return profile
}

func (profile *Profile) SetSurname(Surname string) *Profile {
	profile.Surname = Surname
	return profile
}

func (profile *Profile) SetAddress(Address string) *Profile {
	profile.Address = Address
	return profile
}

func (profile *Profile) SetPhone(Phone string) *Profile {
	profile.Phone = Phone
	return profile
}
