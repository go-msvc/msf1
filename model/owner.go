package model

import "github.com/go-msvc/msf/model"

type Owner struct {
	model.Item
	Company
	MerchantReference string `json:"merchant_reference"`
}
