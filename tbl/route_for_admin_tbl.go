package tbl

type RouteForAdmin struct {
	Id    int64  `gorm:"type:int(50);primaryKey;unique;not null"`
	Sign  int64  `gorm:"type:int(50)"`
	Title string `gorm:"type:varchar(12);comment:'標題 default 中文'"`
	Type  int64  `gorm:"type:int(12);comment:'1:主條目 2:頁面 3:按鍵'"`
	Path  string `gorm:"type:varchar(40);unique;comment:'路徑'"`
}
