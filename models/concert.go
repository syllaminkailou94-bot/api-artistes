package models

type Concert struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	ArtistID int    `json:"artist_id"`
	City     string `json:"city"`
	Venue    string `json:"venue"`
	Date     string `json:"date"`
}
type ConcertPatch struct {
	ArtistID *int    `json:"artist_id"`
	City     *string `json:"city"`
	Venue    *string `json:"venue"`
	Date     *string `json:"date"`
}
