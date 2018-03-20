package main

import "github.com/hashicorp/go-multierror"

// Client interface that exposes all the API methods supported by the underlying Client
type Client interface {

	// Agents
	GetAllAgents() ([]*Agent, error)
	GetAgent(uuid string) (*Agent, error)
	UpdateAgent(uuid string, agent *Agent) (*Agent, error)
	DisableAgent(uuid string) error
	EnableAgent(uuid string) error
	DeleteAgent(uuid string) error
	AgentRunJobHistory(uuid string, offset int) ([]*JobHistory, error)

	// Pipeline Groups
	GetPipelineGroups() ([]*PipelineGroup, error)

	// Pipeline
	GetPipelineStatus(pipelineName string) (*PipelineStatus, error)
	CreatePipeline(pipelineConfig PipelineConfig) (*ApiResponse, *multierror.Error)
	DeletePipeline(pipelineName string) error

	// Jobs
	GetScheduledJobs() ([]*ScheduledJob, error)
	GetJobHistory(pipeline, stage, job string, offset int) ([]*JobHistory, error)

	// Environment Config
	GetAllEnvironmentConfigs() ([]*EnvironmentConfig, error)
	GetEnvironmentConfig(name string) (*EnvironmentConfig, error)

	// Repositories
	GetAllRepositories() (*AllPackageRepositories, *multierror.Error)
	GetRepository(id string) (*PackageRepository, *multierror.Error)

	// Packages
	CreatePackage(pkg Package) (*Package, *ApiResponse, *multierror.Error)
	DeletePackage(id string) (*ApiResponse, *multierror.Error)
}

