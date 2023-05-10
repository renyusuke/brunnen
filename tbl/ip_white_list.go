package tbl

type IpWhiteList struct {
	Id        int64  `gorm:"primaryKey;unique;not null;type:int(250)"`
	Ip        string `gorm:"type:varchar(60)"`
	Mark      string `gorm:"type:varchar(50)"`
	Type      int64  `gorm:"type:int(12)"`
	Status    int64  `gorm:"type:int(12)"`
	CreatedAt int64  `gorm:"type:bigint(20)"`
	UpdatedAt int64  `gorm:"type:bigint(12)"`
}
