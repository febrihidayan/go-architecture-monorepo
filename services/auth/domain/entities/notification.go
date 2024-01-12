package entities

type NotificationSends struct {
	UserId       string
	TemplateName string
	Data         string
	Services     []string
	PathEmail    string
}

const (
	TemplateTypeWelcome = "welcome"
)

const (
	NotificationTypeEmail = "email"
)
