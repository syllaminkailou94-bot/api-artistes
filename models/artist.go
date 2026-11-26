package models

type Artist struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

type ArtistPatch struct {
	ID    int     `json:"id" gorm:"primaryKey"`
	Name  *string `json:"name"`
	Genre *string `json:"genre"`
}
