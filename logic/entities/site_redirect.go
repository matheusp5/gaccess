package entities

type SiteRedirect struct {
	Id       uint32 `gorm:"primaryKey"`
	Code     string
	Redirect string
	Accesses uint32
}
