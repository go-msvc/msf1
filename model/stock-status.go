package model

import "github.com/go-msvc/msf/model"

type StockStatus struct {
	model.Item
	Name string `uniq:"uniqName"`
}
