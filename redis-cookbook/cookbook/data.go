package cookbook

import (
	"nipun.io/config"
	"nipun.io/entity"
)

var Users []entity.User

func init() {
	Users = []entity.User{
		{
			Name:              "Apple",
			Email:             "apple@example.xyz",
			Contact:           "1234567890",
			Age:               23,
			Wallet:            1000,
			TransactionListId: "apple-transaction-list12345",
			Address: entity.Address{
				Place:    "#1,Ist Floor, BTM Layout",
				Locality: "BTM",
				District: "Bangalore",
				State:    "Karnataka",
				Country:  "India",
			},
		},
		{
			Name:              "Orange",
			Email:             "orange@example.xyz",
			Contact:           "1234567890",
			Age:               23,
			Wallet:            1000,
			TransactionListId: "orange-transaction-list12345",
			Address: entity.Address{
				Place:    "#3,Ist Floor, Indranagar",
				Locality: "Domlur",
				District: "Bangalore",
				State:    "Karnataka",
				Country:  "India",
			},
		},
		{
			Name:              "Grape",
			Email:             "grape@example.xyz",
			Contact:           "1234567890",
			Age:               23,
			Wallet:            1000,
			TransactionListId: "grape-transaction-list12345",
			Address: entity.Address{
				Place:    "#1,Ist Floor, Diamond district",
				Locality: "Domlur",
				District: "Bangalore",
				State:    "Karnataka",
				Country:  "India",
			},
		},
		{
			Name:              "Banana",
			Email:             "banana@example.xyz",
			Contact:           "1234567890",
			Age:               23,
			Wallet:            1000,
			TransactionListId: "apple-transaction-list12345",
			Address: entity.Address{
				Place:    "#1,Ist Floor, BTM Layout",
				Locality: "BTM",
				District: "Bangalore",
				State:    "Karnataka",
				Country:  "India",
			},
		},
		{
			Name:              "Mango",
			Email:             "mango@example.xyz",
			Contact:           "1234567890",
			Age:               23,
			Wallet:            1000,
			TransactionListId: "apple-transaction-list12345",
			Address: entity.Address{
				Place:    "#1,Ist Floor, BTM Layout",
				Locality: "Indranagar",
				District: "Bangalore",
				State:    "Karnataka",
				Country:  "India",
			},
		},
		{
			Name:              "Pineapple",
			Email:             "pineapple@example.xyz",
			Contact:           "1234567890",
			Age:               23,
			Wallet:            1000,
			TransactionListId: "apple-transaction-list12345",
			Address: entity.Address{
				Place:    "#1,Ist Floor, BTM Layout",
				Locality: "BTM",
				District: "Bangalore",
				State:    "Karnataka",
				Country:  "India",
			},
		},
	}
}

var odd = config.Odd
var even = config.Even
var prime = config.Prime
var perfectSquare = config.PerfectSquare

var numbers = []entity.Number {
	{
	Value:1,
	Categories : []string{odd, perfectSquare},
},
	{
		Value:2,
		Categories : []string{even, prime},
	},
	{
		Value:3,
		Categories : []string{odd, prime},
	},
	{
		Value:4,
		Categories : []string{even, perfectSquare},
	},	{
		Value:5,
		Categories : []string{odd, prime},
	},	{
		Value:6,
		Categories : []string{even},
	},	{
		Value:7,
		Categories : []string{odd, prime},
	},	{
		Value:8,
		Categories : []string{even},
	},	{
		Value:9,
		Categories : []string{odd, prime, perfectSquare},
	},	{
		Value:10,
		Categories : []string{even},
	},	{
		Value:11,
		Categories : []string{odd, prime},
	},
	{
		Value:12,
		Categories : []string{even},
	},
}

func getUserData() []entity.User {
	return Users
}

func getNumberData() []entity.Number {
	return numbers
}

func getFileNames() []string {
	return []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt"}
}