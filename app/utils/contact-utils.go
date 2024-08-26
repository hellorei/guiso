package utils

import (
	"guiso/app/types"
)

func GetContactInfoArray() []types.ContactInfo {
	return []types.ContactInfo{
		{
			Name: 		"General",
			Email: 		"info@guiso.io",
		},
	}
}