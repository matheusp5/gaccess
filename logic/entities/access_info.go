package entities

type AccessInfo struct {
	Id             uint32 `gorm:"primaryKey"`
	SiteRedirectId uint32
	Ip             string
	City           string
	State          string
	Country        string
}
