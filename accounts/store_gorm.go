package accounts

import (
	"github.com/flow-hydraulics/flow-wallet-api/datastore"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	db.AutoMigrate(&Account{})
	return &GormStore{db}
}

func (s *GormStore) Accounts(o datastore.ListOptions) (aa []Account, err error) {
	err = s.db.
		Order("created_at desc").
		Limit(o.Limit).
		Offset(o.Offset).
		Find(&aa).Error
	return
}

func (s *GormStore) Account(address string) (a Account, err error) {
	err = s.db.First(&a, "address = ?", address).Error
	return
}

func (s *GormStore) InsertAccount(a *Account) error {
	return s.db.Create(a).Error
}
