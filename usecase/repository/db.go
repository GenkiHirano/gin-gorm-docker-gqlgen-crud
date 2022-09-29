package usecase

import "gorm.io/gorm"

// ファイル構造整理する (Repositoryでネストさせるか・させないか)
type DBRepository interface {
    Begin() *gorm.DB
    Connect() *gorm.DB
}
