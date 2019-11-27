package models

// PluginSetting for plugin
type PluginSetting struct {
	Type            string `json:"type" bson:"type"`
	RepoName        string `json:"repo_name" bson:"repo_name"`
	Active          bool   `json:"active" bson:"active"`
	Description     string `json:"description" bson:"description"`
	BuildFolderName string `json:"build_folder_name" bson:"build_folder_name"`
}

// TODO validate repo name
func (p PluginSetting) BeforeCreate() {
	return
}

type ToggleActiveSetting struct {
	RepoName string `json:"repo_name"`
	Active   bool   `json:"active"`
}
