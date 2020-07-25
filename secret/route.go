package secret

//GetSecret ...
func GetSecret() *Secret {
	return &Secret{
		Name:  "Foo",
		Value: "Bar",
	}
}
