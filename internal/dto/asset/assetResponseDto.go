package asset

import "time"

type AssetResponseDTO struct {
	Name      string    `json:"name"`
	Uid       uint64    `json:"uid"`
	Data      []byte    `json:"data"`
	CreatedAt time.Time `json:"createdAt"`
}
