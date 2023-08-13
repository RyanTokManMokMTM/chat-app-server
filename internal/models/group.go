package models

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Uuid        string `gorm:"type:varchar(64);not null;unique_index:idx_uuid"`
	GroupName   string `gorm:"type:varchar(64);not null"`
	GroupAvatar string
	GroupLead   uint `gorm:"not null;index"`

	LeadInfo  User   `gorm:"foreignKey:GroupLead;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UsersInfo []User `gorm:"many2many:users_groups;foreignKey:Id;joinForeignKey:GroupId"`
	CommonField
}

func (g *Group) BeforeCreate(tx *gorm.DB) error {
	g.Uuid = uuid.New().String()
	//g.GroupAvatar = "/defaultGroup.jpg"
	return nil
}

func (g *Group) TableName() string {
	return "groups_info"
}

func (g *Group) InsertOne(ctx context.Context, db *gorm.DB) error {
	//TODO: Add Group Lead to the group member list?
	return db.WithContext(ctx).Debug().Transaction(func(tx *gorm.DB) error {
		db := tx.WithContext(ctx).Debug().Create(&g)
		err := db.Error
		if err != nil {
			return err
		}

		if db.RowsAffected == 0 {
			return errors.New("db row affected 0")
		}

		return tx.WithContext(ctx).Debug().Create(&UserGroup{GroupId: g.Id, UserId: g.GroupLead}).Error
	})
}

func (g *Group) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().First(&g).Error
}
func (g *Group) FindOneByUUID(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("uuid = ?", g.Uuid).First(&g).Error
}

func (g *Group) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("id = ?", g.Id).Delete(&g).Error
}

func (g *Group) UpdateOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(g).Where("id = ?", g.Id).Update("GroupName", g.GroupName).Error
}

func (g *Group) UpdateOneAvatar(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(g).Where("id = ?", g.Id).Update("GroupAvatar", g.GroupAvatar).Error
}

func (g *Group) SearchGroup(ctx context.Context, db *gorm.DB, query string) ([]*Group, error) {
	var groups []*Group
	if err := db.WithContext(ctx).Debug().Model(g).Where("group_name like ?", "%"+query+"%").Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}
