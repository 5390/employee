package models

type Employee struct {
	Id         int     `db:"id" json:"id"`
	Name       *string `db:"name" json:"name"`
	Address    string  `db:"address" json:"addres"`
	Department *string `db:"department" json:"department"`
	Skills     *string `db:"skills" json:"skills"`
	IsDeleted  int     `db:"is_deleted" json:"isDeleted"`
}

type Search struct {
	SearchKey string `json:"searchKey"`
}
