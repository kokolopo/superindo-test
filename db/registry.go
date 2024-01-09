package db

import "superindo-test/domain"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: domain.User{}},
		{Model: domain.ProductCategory{}},
		{Model: domain.ProductVariant{}},
		{Model: domain.Product{}},
		{Model: domain.Transaction{}},
		{Model: domain.TransactionDetail{}},
	}
}
