package gocd

import "github.com/hashicorp/go-multierror"

// Client interface that exposes all the API methods supported by the underlying Client
type Client interface {
	// Agents API
	GetAllAgents() ([]*Agent, error)
	GetAgent(uuid string) (*Agent, error)
	UpdateAgent(uuid string, agent *Agent) (*Agent, error)
	DisableAgent(uuid string) error
	EnableAgent(uuid string) error
	DeleteAgent(uuid string) error
	AgentRunJobHistory(uuid string, offset int) ([]*JobHistory, error)

	// Pipeline Groups API
	GetPipelineGroups() ([]*PipelineGroup, error)

	// Pipeline API
	GetPipelineStatus(pipelineName string) (*PipelineStatus, error)
	CreatePipeline(pipelineConfig PipelineConfig) (*ApiResponse, *multierror.Error)
	DeletePipeline(pipelineName string) error

	// Jobs API
	GetScheduledJobs() ([]*ScheduledJob, error)
	GetJobHistory(pipeline, stage, job string, offset int) ([]*JobHistory, error)

	// Environment Config API
	GetAllEnvironmentConfigs() ([]*EnvironmentConfig, error)
	GetEnvironmentConfig(name string) (*EnvironmentConfig, error)

	// Repositories
	GetAllRepositories() (*AllPackageRepositories, *multierror.Error)
	GetRepository() (*PackageRepository, *multierror.Error)
}

//TODO: fix this back