package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/internal/models"
)

func (d *DAO) InsertOneStickerGroup(ctx context.Context, name string) (*models.Sticker, error) {
	sticker := &models.Sticker{
		StickerName: name,
	}

	if err := sticker.InsertOne(ctx, d.engine); err != nil {
		return nil, err
	}
	return sticker, nil
}

func (d *DAO) InsertOneStickerGroupWithResources(ctx context.Context, model *models.Sticker) (*models.Sticker, error) {
	if err := model.InsertOne(ctx, d.engine); err != nil {
		return nil, err
	}
	return model, nil
}

func (d *DAO) InsertStickerListIntoGroup(ctx context.Context, sticker *models.Sticker, paths []string) error {
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

func (d *DAO) UpdateOneStickerGroup(ctx context.Context, sticker *models.Sticker) error {
	return sticker.UpdateOne(ctx, d.engine)
}

func (d *DAO) GetStickerGroupList(ctx context.Context) ([]*models.Sticker, error) {
	return (&models.Sticker{}).FindAll(ctx, d.engine)
}
