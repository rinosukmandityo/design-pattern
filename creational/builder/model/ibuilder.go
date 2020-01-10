package model

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func GetBuilder(builderType string) iBuilder {
	switch builderType {
	case "igloo":
		return newIglooBuilder()
	default:
		return newNormalBuilder()
	}
}
