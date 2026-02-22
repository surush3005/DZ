package card


func Issue(currency types.Currency, color string, name string) types.Card{
	return types.Card{
		Id:	1000,
		PAN: "5058 XXXX XXXX 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,

		Id:	1001,
		PAN: "5058 XXXX XXXX 9999",
		Balance: -100,
		Currency: currency,
		Color: color,
		Name: name,
		Active: false,


	}
}