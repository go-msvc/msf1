package model

import "github.com/go-msvc/msf/model"

type Inventory struct {
	model.Item
	Owner
	ProductReference string
}
