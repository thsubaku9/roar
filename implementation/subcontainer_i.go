package main

import (
	roarInterfaces "roar/interfaces"
	roar "roar/types"
)

func CreateNewSubContainer() roarInterfaces.SubContainer {
	sc := roar.CreateSarr()

	return &sc
}
