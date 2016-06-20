package models

type Type struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetTypes() (*[]Type, error) {
	types := []Type{}
	err := db.Find(&types, &Type{}).Error
	return &types, err
}

func GetTypeByName(name string) (*Type, error) {
	typeInst := Type{}
	err := db.Where("name = ?", name).Find(&typeInst, Type{}).Error
	return &typeInst, err
}