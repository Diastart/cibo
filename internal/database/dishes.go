package database

type Dish struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ImageURL    string  `json:"img"`
	Calorie     int     `json:"calorie"`
	Runtime     int     `json:"runtime"`
	Like        float64 `json:"like"`
	Dislike     float64 `json:"dislike"`
	Nationality string  `json:"nationality"`
}