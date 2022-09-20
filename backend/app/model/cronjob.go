package model

type Cronjob struct {
	BaseModel

	Name     string `gorm:"type:varchar(64);not null" json:"name"`
	Type     string `gorm:"type:varchar(64);not null" json:"type"`
	SpecType string `gorm:"type:varchar(64);not null" json:"specType"`
	Spec     string `gorm:"type:varchar(64);not null" json:"spec"`
	Week     uint64 `gorm:"type:decimal" json:"week"`
	Day      uint64 `gorm:"type:decimal" json:"day"`
	Hour     uint64 `gorm:"type:decimal" json:"hour"`
	Minute   uint64 `gorm:"type:decimal" json:"minute"`

	Script         string `gorm:"longtext" json:"script"`
	Website        string `gorm:"type:varchar(64)" json:"website"`
	Database       string `gorm:"type:varchar(64)" json:"database"`
	URL            string `gorm:"type:varchar(256)" json:"url"`
	SourceDir      string `gorm:"type:varchar(256)" json:"sourceDir"`
	TargetDirID    uint64 `gorm:"type:decimal" json:"targetDirID"`
	ExclusionRules string `gorm:"longtext" json:"exclusionRules"`
	RetainCopies   uint64 `gorm:"type:decimal" json:"retainCopies"`

	Status string `gorm:"type:varchar(64)" json:"status"`
}
