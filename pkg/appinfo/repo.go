package appinfo

// Repository manages data access.
type Repository interface {

	// Get app info by system
	Get(system string) (a *AppInfo, err error)
}
