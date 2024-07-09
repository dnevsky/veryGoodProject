package postgresDB

import (
	"context"
	"errors"
	"fmt"
	assetDto "github.com/dnevsky/veryGoodProject/internal/dto/asset"
	"github.com/dnevsky/veryGoodProject/internal/models"
	repoErrors "github.com/dnevsky/veryGoodProject/internal/repository/errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AssetRepository struct {
	db *pgxpool.Pool
}

func NewAssetRepository(db *pgxpool.Pool) *AssetRepository {
	return &AssetRepository{db: db}
}

func (r *AssetRepository) UploadAsset(ctx context.Context, dto assetDto.UploadAssetDTO) (asset *models.Asset, err error) {
	asset = &models.Asset{}

	if err := r.checkExists(ctx, dto.Name, dto.XUserID); err != nil {
		return nil, repoErrors.ErrAlreadyExists
	}

	query := `INSERT INTO assets (name, uid, data) VALUES ($1, $2, $3) RETURNING name, uid, data, created_at`

	err = r.db.
		QueryRow(ctx, query, dto.Name, dto.ServiceDTO.XUserID, dto.Body).
		Scan(&asset.Name, &asset.Uid, &asset.Data, &asset.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create asset: %w", err)
	}

	return asset, nil
}

func (r *AssetRepository) GetAsset(ctx context.Context, dto assetDto.GetAssetDTO) (asset *models.Asset, err error) {
	asset = &models.Asset{}
	query := `SELECT name, uid, data, created_at FROM assets WHERE name=$1 AND uid=$2`

	err = r.db.
		QueryRow(ctx, query, dto.Name, dto.ServiceDTO.XUserID).
		Scan(&asset.Name, &asset.Uid, &asset.Data, &asset.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repoErrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get asset: %w", err)
	}

	return asset, nil
}

func (r *AssetRepository) GetAssets(ctx context.Context, dto assetDto.GetAssetsDTO) (assets []models.Asset, count int64, err error) {
	query := `SELECT name, uid, data, created_at FROM assets WHERE uid=$1 LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(ctx, query, dto.ServiceDTO.XUserID, dto.PaginationDTO.Limit, dto.PaginationDTO.Offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get assets: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var asset models.Asset
		err = rows.Scan(&asset.Name, &asset.Uid, &asset.Data, &asset.CreatedAt)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan assets: %w", err)
		}
		assets = append(assets, asset)
	}

	queryCount := `SELECT COUNT(*) FROM assets WHERE uid=$1`
	err = r.db.
		QueryRow(ctx, queryCount, dto.ServiceDTO.XUserID).
		Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get assets count: %w", err)
	}

	return assets, count, nil
}

func (r *AssetRepository) DeleteAsset(ctx context.Context, dto assetDto.DeleteAssetDTO) (err error) {
	query := `DELETE FROM assets WHERE name=$1 AND uid=$2`

	_, err = r.db.Exec(ctx, query, dto.Name, dto.ServiceDTO.XUserID)
	if err != nil {
		return fmt.Errorf("failed to delete asset: %w", err)
	}

	return nil
}

func (r *AssetRepository) checkExists(ctx context.Context, name string, uid uint64) error {
	var count int
	query := `SELECT COUNT(*) FROM assets WHERE name=$1 AND uid=$2`
	err := r.db.QueryRow(ctx, query, name, uid).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("asset already exists")
	}

	return nil
}
