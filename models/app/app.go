package app

type Exportable interface {
	Export() App
}

type Type string

const (
	FlatpakApp Type = "flatpak"
	SnapApp    Type = "snap"
	ExternApp  Type = "extern"
)

type App struct {
	AppType Type `json:"app_type" bson:"app_type"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	Version string `json:"version" bson:"version"`

	StatsAvailable bool `json:"stats_available" bson:"stats_available"`

	Stats Stats `json:"stats" bson:"stats"`

	IconURL     string    `json:"icon_url" bson:"icon_url"`
	HeaderURL   string    `json:"header_url" bson:"header_url"`
	Screenshots []string  `json:"screenshots" bson:"screenshots"`
	Publisher   Publisher `json:"publisher" bson:"publisher"`

	Package string `json:"package" bson:"package"`
}
