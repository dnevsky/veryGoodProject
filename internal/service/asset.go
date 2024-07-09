package service

import (
	"context"
	assetDto "github.com/dnevsky/veryGoodProject/internal/dto/asset"
	"github.com/dnevsky/veryGoodProject/internal/models"
	"github.com/dnevsky/veryGoodProject/internal/repository"
)

type Asset interface {
	UploadAsset(ctx context.Context, dto assetDto.UploadAssetDTO) (asset *models.Asset, err error)
	GetAsset(ctx context.Context, dto assetDto.GetAssetDTO) (assetResp *assetDto.AssetResponseDTO, err error)
	GetAssets(ctx context.Context, dto assetDto.GetAssetsDTO) (assetsResp []assetDto.AssetResponseDTO, count int64, err error)
	DeleteAsset(ctx context.Context, dto assetDto.DeleteAssetDTO) (err error)
}

type AssetService struct {
	AssetRepository repository.AssetRepository
}

func NewAssetService(
	assetRepo repository.AssetRepository,
) *AssetService {
	return &AssetService{
		AssetRepository: assetRepo,
	}
}

func (s *AssetService) UploadAsset(ctx context.Context, dto assetDto.UploadAssetDTO) (asset *models.Asset, err error) {
	return s.AssetRepository.UploadAsset(ctx, dto)
}

func (s *AssetService) GetAsset(ctx context.Context, dto assetDto.GetAssetDTO) (assetResp *assetDto.AssetResponseDTO, err error) {
	asset, err := s.AssetRepository.GetAsset(ctx, dto)
	if err != nil {
		return nil, err
	}

	assetResp = &assetDto.AssetResponseDTO{
		Name:      asset.Name,
		Uid:       asset.Uid,
		Data:      asset.Data,
		CreatedAt: asset.CreatedAt,
	}

	return assetResp, nil
}

func (s *AssetService) GetAssets(ctx context.Context, dto assetDto.GetAssetsDTO) (assetsResp []assetDto.AssetResponseDTO, count int64, err error) {
	dto.Offset = dto.PaginationDTO.Limit * (dto.PaginationDTO.Page - 1)
	assets, count, err := s.AssetRepository.GetAssets(ctx, dto)
	if err != nil {
		return nil, 0, err
	}

	for _, asset := range assets {
		assetsResp = append(assetsResp, assetDto.AssetResponseDTO{
			Name:      asset.Name,
			Uid:       asset.Uid,
			Data:      asset.Data,
			CreatedAt: asset.CreatedAt,
		})
	}

	return assetsResp, count, nil
}

func (s *AssetService) DeleteAsset(ctx context.Context, dto assetDto.DeleteAssetDTO) (err error) {
	return s.AssetRepository.DeleteAsset(ctx, dto)
}
