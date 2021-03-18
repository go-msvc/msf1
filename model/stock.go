package model

import "github.com/go-msvc/msf/model"

type Stock struct {
	model.Item
	Inventory
	Location
}
