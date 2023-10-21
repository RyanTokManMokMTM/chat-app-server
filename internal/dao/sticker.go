package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
)

func (d *DAO) InsertOneStickerGroup(ctx context.Context, name string) (*models.Sticker, error) {
	sticker := &models.Sticker{
		SickerName: name,
	}

	if err := sticker.InsertOne(ctx, d.engine); err != nil {
		return nil, err
	}
	return sticker, nil
}

func (d *DAO) InsertOneStickerIntoGroup(ctx context.Context, sticker *models.Sticker, paths []string) error {
	return sticker.InsertResources(ctx, d.engine, paths)
}

func (d *DAO) FindOneStickerGroupByStickerUUID(ctx context.Context, uuid string) (*models.Sticker, error) {
	sticker := &models.Sticker{
		Uuid: uuid,
	}

	if err := sticker.FindOneByUuid(ctx, d.engine); err != nil {
		return nil, err
	}
	return sticker, nil
}
