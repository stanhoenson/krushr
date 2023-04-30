package constants

import "time"

const (
	PublishedStatusName   = "Published"
	UnpublishedStatusName = "Unpublished"
	InactiveStatusName    = "Inactive"
	AdminRoleName         = "Admin"
	CreatorRoleName       = "Creator"
	TitleMaxLength        = 50
)

const (
	LegacyWebsiteContentTypeName = "website"
	LegacyWebsiteContentTypeId   = 1012
	LegacyImageContentTypeName   = "interne foto"
	LegacyImageContentTypeId     = 1025
	LegacyTekstContentTypeName   = "tekst"
	LegacyTekstContentTypeId     = 1024
)

const TokenValidityPeriod = time.Hour * 24 * 7 // 7 days

var DefaultLegacyRouteLanguages = []string{"Dutch"}

var Roles = []string{
	AdminRoleName,
	CreatorRoleName,
}

var Statuses = []string{
	PublishedStatusName, InactiveStatusName, UnpublishedStatusName,
}
