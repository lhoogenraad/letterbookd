package tools

import(
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails {
	"leon": {
		AuthToken: "123ABC",
		Username: "leon",
	},
}

var mockCoinDetails = map[string]CoinDetails {
	"leon": {
		Coins: 10000,
		Username: "Leon",
	},
}


func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (db *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 4)

	var clientData = CoinDetails{}

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}


func (d *mockDB) SetupDatabase() error {
	return nil
}
