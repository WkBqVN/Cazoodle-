package model

type Client struct {
	ID         int
	Survey_ids []int `gorm:"surveys_id"`
}
