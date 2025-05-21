package models

type Subscription struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    UserID string `json:"user_id"`
    Plan   string `json:"plan"`
    Status string `json:"status"`
}