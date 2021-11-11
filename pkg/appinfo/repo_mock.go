package appinfo

type mockRepo struct {
}

// NewMockRepository instantiate new mock-up repository for demo.
func NewMockRepository() Repository {
	return &mockRepo{
	}
}

// Get mock appinfo data. 
func (repo *mockRepo) Get(system string) (a *AppInfo, err error) {

	switch system {
	case "iOS": a = &AppInfo{
		System: system,
		LatestVersion: "1.0",
		MinVersion: "0.1",
	}
	case "Android": a = &AppInfo{
		System: system,
		LatestVersion: "1.1",
		MinVersion: "0.1",
	}
	}
	return
}