package event

import (
	"walls-user-service/internal/core/domain/entity"
)

type OtpValidatedEventData struct {
	UserReference string        `json:"user_reference" bson:"user_reference"`
	OtpType       string        `json:"otp_type" bson:"otp_type" validate:"required"`
	Contact       string        `json:"contact" bson:"contact"`
	Device        entity.Device `json:"device" bson:"device"`
}
