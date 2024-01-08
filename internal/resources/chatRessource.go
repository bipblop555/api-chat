package resources

type Chat struct {
	Sender   int    `gorm:"not_null" validate:"required"`
	Receiver int    `gorm:"not_null" validate:"required"`
	Message  string `gorm:"not_null" validate:"required"`
}
