package entities

type Module struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Icon   string `json:"icon"`
	Active bool   `json:"active" gorm:"type:bool not null;default:false"`
}
