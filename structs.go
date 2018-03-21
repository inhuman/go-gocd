package go_gocd

import "encoding/json"

type PipelineStatus struct {
	PausedCause string `json:"pausedCause"`
	PausedBy    string `json:"pausedBy"`
	Paused      bool   `json:"paused"`
	Schedulable bool   `json:"schedulable"`
	Locked      bool   `json:"locked"`
}

type PipelineConfig struct {
	Group    string   `json:"group"`
	Pipeline Pipeline `json:"pipeline"`
}

type Pipeline struct {
	LabelTemplate        string                `json:"label_template"`
	LockBehavior         string                `json:"lock_behavior"`
	Name                 string                `json:"name"`
	Template             interface{}           `json:"template"`
	Parameters           []Parameter           `json:"parameters"`
	Materials            []PipelineMaterial    `json:"materials"`
	Stages               []Stage               `json:"stages"`
	EnvironmentVariables []EnvironmentVariable `json:"environment_variables"`
	Origin               Origin                `json:"origin"`
	Links                Links                 `json:"_links,omitempty"`
}

type Stage struct {
	Name                  string                `json:"name"`
	FetchMaterials        bool                  `json:"fetch_materials"`
	CleanWorkingDirectory bool                  `json:"clean_working_directory"`
	NeverCleanupArtifacts bool                  `json:"never_cleanup_artifacts"`
	Approval              Approval              `json:"approval"`
	EnvironmentVariables  []EnvironmentVariable `json:"environment_variables"`
	Jobs                  []Job                 `json:"jobs"`
}

type PipelineMaterial struct {
	Type   string `json:"type"`
	Errors map[string][]json.RawMessage
	Attributes struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Branch      string `json:"branch"`
		Destination string `json:"destination"`
		//Filter          Filter `json:"filter,omitempty"`
		InvertFilter    bool   `json:"invert_filter"`
		AutoUpdate      bool   `json:"auto_update"`
		SubmoduleFolder string `json:"submodule_folder"`
		ShallowClone    bool   `json:"shallow_clone"`
	} `json:"attributes"`
}

type ApiResponse struct {
	Message string `json:"message"`
	Data struct {
		Errors               map[string][]json.RawMessage
		Materials            []PipelineMaterial
		EnvironmentVariables []EnvironmentVariable `json:"environment_variables"`
	}
}

type EnvironmentVariable struct {
	Secure         bool   `json:"secure"`
	Name           string `json:"name"`
	Value          string `json:"value,omitempty"`
	EncryptedValue string `json:"encrypted_value,omitempty"`
	Errors struct {
		ValueForDisplay []json.RawMessage `json:"valueForDisplay"`
	} `json:"errors,omitempty"`
}

type Job struct {
	Name                 string                `json:"name"`
	RunInstanceCount     interface{}           `json:"run_instance_count"`
	Timeout              int                   `json:"timeout"`
	EnvironmentVariables []EnvironmentVariable `json:"environment_variables"`
	Resources            []string              `json:"resources"`
	Tasks                []Task                `json:"tasks"`
	Tabs                 []Tab                 `json:"tabs"`
	Artifacts            []Artifact            `json:"artifacts"`
	Properties           []Property            `json:"properties"`
	ElasticProfileId     string                `json:"elastic_profile_id"`
}

type Filter struct {
	Ignore []string `json:"ignore"`
}

type Tab struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Task struct {
	Type string `json:"type"`
	Attributes struct {
		RunIf            []string `json:"run_if"`
		Command          string   `json:"command"`
		WorkingDirectory string   `json:"working_directory"`
		OnCancel         *Task    `json:"on_cancel"`
	} `json:"attributes"`
}

type Property struct {
	Name   string `json:"name"`
	Source string `json:"source"`
	Xpath  string `json:"xpath"`
}

type Artifact struct {
	Type        string `json:"type"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Origin struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Approval struct {
	Type          string        `json:"type"`
	Authorization Authorization `json:"authorization"`
}

type Authorization struct {
	Roles []string `json:"roles"`
	Users []string `json:"users"`
}

type Timer struct {
	Spec          string
	OnlyOnChanges bool
}

type Links struct {
	Self struct {
		Href string `json:"href"`
	} `json:"self"`
	Doc struct {
		Href string `json:"href"`
	} `json:"doc"`
}

type PackageRepository struct {
	Links          Links           `json:"_links,omitempty"`
	Id             string          `json:"id"`
	Name           string          `json:"name"`
	PluginMetadata PluginMetadata  `json:"plugin_metadata,omitempty"`
	Configuration  []Configuration `json:"configuration,omitempty"`
}

type Repository struct {
	Links          Links           `json:"_links,omitempty"`
	RepoId         string          `json:"repo_id"`
	Name           string          `json:"name"`
	PluginMetadata PluginMetadata  `json:"plugin_metadata,omitempty"`
	Configuration  []Configuration `json:"configuration,omitempty"`
}

type Configuration struct {
	Key            string `json:"key"`
	Value          string `json:"value,omitempty"`
	EncryptedValue string `json:"encrypted_value,omitempty"`
}

type PluginMetadata struct {
	Id      string `json:"id"`
	Version string `json:"version"`
}

type AllPackageRepositories struct {
	Links Links `json:"_links"`
	Embedded struct {
		Repositories []PackageRepository `json:"package_repositories"`
	} `json:"_embedded"`
}

type Package struct {
	Name          string            `json:"name"`
	Id            string            `json:"id"`
	AutoUpdate    bool              `json:"auto_update"`
	PackageRepo   PackageRepository `json:"package_repo"`
	Configuration []Configuration   `json:"configuration"`
}
