package model

type ImageModel struct {
	PriID
	UID
	Path     string `json:"path" gorm:"column:path;not null;type:varchar(120);comment:资源地址"`
	Filename string `json:"filename" gorm:"column:filename;not null;type:varchar(120);comment:文件名"`
	MimeType string `json:"mime_type" gorm:"column:mime_type;not null;type:varchar(30);comment:文件类型"`
	FileMd5  string `json:"file_md5" gorm:"column:file_md5;not null;type:char(32);comment:文件md5签名"`
	Size     int32  `json:"size" gorm:"column:size;not null;comment:文件大小"`
	Create
}
