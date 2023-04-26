package constants

const (
	PublishedStatusName   = "Published"
	UnpublishedStatusName = "Unpublished"
	InactiveStatusName    = "Inactive"
	AdminRoleName         = "Admin"
	CreatorRoleName       = "Creator"
	TitleMaxLength        = 50
)

var DefaultLegacyRouteLanguages = []string{"Dutch"}

var Roles = []string{
	AdminRoleName,
	CreatorRoleName,
}

var Statuses = []string{
	PublishedStatusName, InactiveStatusName, UnpublishedStatusName,
}
