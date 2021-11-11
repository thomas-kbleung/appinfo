package appinfo

// AppInfo data model.
type AppInfo struct {

	// System specifies either iOS or Android
	System            string `json:"system" `

	// LatestVersion specifies latest version of published app.
	LatestVersion     string `json:"latest_version" `

	// MinVersion specifies minimum version application should run.
	// If the starting application is older than this, 
	// forces user to update the application from AppStore or Google Play.
	MinVersion        string `json:"min_version" `
}