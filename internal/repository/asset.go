package repository

import (
	"context"
	assetDto "github.com/dnevsky/veryGoodProject/internal/dto/asset"
	"github.com/dnevsky/veryGoodProject/internal/models"
)

type AssetRepository interface {
	UploadAsset(ctx context.Context, dto assetDto.UploadAssetDTO) (asset *models.Asset, err error)
	GetAsset(ctx context.Context, dto assetDto.GetAssetDTO) (asset *models.Asset, err error)
	GetAssets(ctx context.Context, dto assetDto.GetAssetsDTO) (assets []models.Asset, count int64, err error)
	DeleteAsset(ctx context.Context, dto assetDto.DeleteAssetDTO) (err error)
}
