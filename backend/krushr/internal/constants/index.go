package constants

const (
	PublishedStatusName   = "Published"
	UnpublishedStatusName = "Unpublished"
	InactiveStatusName    = "Inactive"
	AdminRoleName         = "Admin"
	CreatorRoleName       = "Creator"
	TitleMaxLength        = 50
)

var Roles = []string{
	AdminRoleName,
	CreatorRoleName,
}

var Statuses = []string{
	PublishedStatusName, InactiveStatusName, UnpublishedStatusName,
}
