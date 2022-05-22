package model

type Greenbird struct {
	Order   int    `gorm:"primary_key" json:"order"`
	Title   string `gorm:"type:text" json:"title"`
	Content string `gorm:"type:text" json:"content"`
}

//api
type GreenbirdData struct {
	Greenbirds []Greenbird `json:"greenBirds"`
}
