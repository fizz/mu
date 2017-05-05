package common

import (
	"io"
	"time"
)

// Context defines the context object passed around
type Context struct {
	Config          Config
	StackManager    StackManager
	ClusterManager  ClusterManager
	ElbManager      ElbManager
	RdsManager      RdsManager
	ParamManager    ParamManager
	PipelineManager PipelineManager
	LogsManager     LogsManager
	DockerManager   DockerManager
	DockerOut       io.Writer
	TaskManager     TaskManager
}

// Config defines the structure of the yml file for the mu config
type Config struct {
	Environments []Environment `yaml:"environments,omitempty"`
	Service      Service       `yaml:"service,omitempty"`
	Basedir      string        `yaml:"-"`
	Repo         struct {
		Name     string
		OrgName  string
		Slug     string
		Revision string
		Provider string
	} `yaml:"-"`
	Templates map[string]interface{} `yaml:"templates,omitempty"`
}

// Environment defines the structure of the yml file for an environment
type Environment struct {
	Name         string `yaml:"name,omitempty"`
	Loadbalancer struct {
		HostedZone  string `yaml:"hostedzone,omitempty"`
		Name        string `yaml:"name,omitempty"`
		Certificate string `yaml:"certificate,omitempty"`
		Internal    bool   `yaml:"internal,omitempty"`
	} `yaml:"loadbalancer,omitempty"`
	Cluster struct {
		InstanceType      string `yaml:"instanceType,omitempty"`
		ImageID           string `yaml:"imageId,omitempty"`
		InstanceTenancy   string `yaml:"instanceTenancy,omitempty"`
		DesiredCapacity   int    `yaml:"desiredCapacity,omitempty"`
		MaxSize           int    `yaml:"maxSize,omitempty"`
		KeyName           string `yaml:"keyName,omitempty"`
		SSHAllow          string `yaml:"sshAllow,omitempty"`
		ScaleOutThreshold int    `yaml:"scaleOutThreshold,omitempty"`
		ScaleInThreshold  int    `yaml:"scaleInThreshold,omitempty"`
		HTTPProxy         string `yaml:"httpProxy,omitempty"`
	} `yaml:"cluster,omitempty"`
	Discovery struct {
		Provider      string            `yaml:"provider,omitempty"`
		Configuration map[string]string `yaml:"configuration,omitempty"`
	} `yaml:"discovery,omitempty"`
	VpcTarget struct {
		VpcID        string   `yaml:"vpcId,omitempty"`
		EcsSubnetIds []string `yaml:"ecsSubnetIds,omitempty"`
		ElbSubnetIds []string `yaml:"elbSubnetIds,omitempty"`
	} `yaml:"vpcTarget,omitempty"`
}

// Service defines the structure of the yml file for a service
type Service struct {
	Name            string                 `yaml:"name,omitempty"`
	DesiredCount    int                    `yaml:"desiredCount,omitempty"`
	Dockerfile      string                 `yaml:"dockerfile,omitempty"`
	ImageRepository string                 `yaml:"imageRepository,omitempty"`
	Port            int                    `yaml:"port,omitempty"`
	HealthEndpoint  string                 `yaml:"healthEndpoint,omitempty"`
	CPU             int                    `yaml:"cpu,omitempty"`
	Memory          int                    `yaml:"memory,omitempty"`
	Environment     map[string]interface{} `yaml:"environment,omitempty"`
	PathPatterns    []string               `yaml:"pathPatterns,omitempty"`
	Priority        int                    `yaml:"priority,omitempty"`
	Pipeline        Pipeline               `yaml:"pipeline,omitempty"`
	Database        Database               `yaml:"database,omitempty"`
}

// Database definition
type Database struct {
	Name              string `yaml:"name,omitempty"`
	InstanceClass     string `yaml:"instanceClass,omitempty"`
	Engine            string `yaml:"engine,omitempty"`
	IamAuthentication bool   `yaml:"iamAuthentication,omitempty"`
	MasterUsername    string `yaml:"masterUsername,omitempty"`
	AllocatedStorage  string `yaml:"allocatedStorage,omitempty"`
}

// Pipeline definition
type Pipeline struct {
	Source struct {
		Provider string `yaml:"provider,omitempty"`
		Repo     string `yaml:"repo,omitempty"`
		Branch   string `yaml:"branch,omitempty"`
	} `yaml:"source,omitempty"`
	Build struct {
		Type        string `yaml:"type,omitempty"`
		ComputeType string `yaml:"computeType,omitempty"`
		Image       string `yaml:"image,omitempty"`
	} `yaml:"build,omitempty"`
	Acceptance struct {
		Environment string `yaml:"environment,omitempty"`
		Type        string `yaml:"type,omitempty"`
		ComputeType string `yaml:"computeType,omitempty"`
		Image       string `yaml:"image,omitempty"`
	} `yaml:"acceptance,omitempty"`
	Production struct {
		Environment string `yaml:"environment,omitempty"`
	} `yaml:"production,omitempty"`
	MuBaseurl string `yaml:"muBaseurl,omitempty"`
	MuVersion string `yaml:"muVersion,omitempty"`
}

// Stack summary
type Stack struct {
	ID             string
	Name           string
	Status         string
	StatusReason   string
	LastUpdateTime time.Time
	Tags           map[string]string
	Outputs        map[string]string
	Parameters     map[string]string
}

// StackType describes supported stack types
type StackType string

// List of valid stack types
const (
	StackTypeVpc      StackType = "vpc"
	StackTypeTarget             = "target"
	StackTypeCluster            = "cluster"
	StackTypeConsul             = "consul"
	StackTypeRepo               = "repo"
	StackTypeService            = "service"
	StackTypePipeline           = "pipeline"
	StackTypeDatabase           = "database"
	StackTypeBucket             = "bucket"
)

// Constants for available command names and options
const (
	EnvSubCmdCount          = 5
	SingleAliasIndex        = 0
	SvcSubCmdCount          = 6
	SvcShowFormatFlagIndex  = 0
	SvcLogFlagCount         = 3
	EnvLogFollowFlagIndex   = 0
	EnvLogDurationFlagIndex = 1
	SvcLogServiceFlagIndex  = 0
	SvcLogFollowFlagIndex   = 1
	SvcLogDurationFlagIndex = 2
	ShowFlagCount           = 1
	ExeArgsSvcIndex         = 1
	ExeArgsCmdIndex         = 2
	EnvLogsFlagCount        = 2
	SvcPushTagFlagIndex     = 0
	SvcDeployTagFlagIndex   = 0
	SvcUndeploySvcFlagIndex = 1
	DefaultLogDurationValue = 1 * time.Minute
	SvcCmd                  = "service"
	SvcAlias                = "svc"
	SvcUsage                = "options for managing services"
	SvcShowUsage            = "[<service>]"
	SvcLogUsage             = "show service logs"
	SvcLogArgUsage          = "<environment> [<filter>...]"
	SvcLogServiceFlagUsage  = "service name to view logs for"
	SvcPushTagFlagUsage     = "tag to push"
	SvcDeployTagFlagUsage   = "tag to deploy"
	TagFlagName             = "tag, t"
	EnvCmd                  = "environment"
	EnvAlias                = "env"
	EnvUsage                = "options for managing environments"
	EnvArgUsage             = "<environment>"
	Tag                     = "tag"
	UpsertCmd               = "upsert"
	UpsertAlias             = "up"
	UpsertUsage             = "create/update an environment"
	ListCmd                 = "list"
	TerminateCmd            = "terminate"
	TerminateAlias          = "term"
	TerminateUsage          = "terminate an environment"
	ListAlias               = "ls"
	ListUsage               = "list environments"
	ShowCmd                 = "show"
	ShowCmdUsage            = "show environment details"
	ExeCmd                  = "exec"
	ExeUsage                = "execute a command in environment"
	ExeArgs                 = "<environment> <service> <command>"
	LogsCmd                 = "logs"
	LogsArgs                = "<environment> [<filter>...]"
	LogsUsage               = "show environment logs"
	Format                  = "format"
	FormatFlag              = "format, f"
	FormatFlagUsage         = "output format, either 'json' or 'cli' (default: cli)"
	FormatFlagDefault       = "cli"
	Follow                  = "follow"
	FollowFlag              = "follow, f"
	ServiceFlag             = "service, s"
	FollowUsage             = "follow logs for latest changes"
	SearchDuration          = "search-duration"
	SearchDurationUsage     = "duration to go into the past for searching (e.g. 5m for 5 minutes)"
	SearchDurationFlag      = "search-duration, t"
	PushCmd                 = "push"
	SvcPushCmdUsage         = "push service to repository"
	DeployCmd               = "deploy"
	SvcDeployCmdUsage       = "deploy service to environment"
	UndeployCmd             = "undeploy"
	SvcUndeployCmdUsage     = "undeploy service from environment"
	SvcUndeployArgsUsage    = "<environment> [<service>]"
)

// Constants to prevent multiple updates when making changes.
const (
	Zero                       = 0
	ECSRunTaskDefaultCount     = 1
	Empty                      = ""
	Space                      = " "
	Spaces                     = "   "
	DefaultVersion             = "0.0.0-local"
	ECSServiceNameParameterKey = "ServiceName"
	ECSTaskDefinitionOutputKey = "MicroserviceTaskDefinition"
	NoEnvValidation            = "environment must be provided"
	NoSvcValidation            = "service must be provided"
	EmptySvcValidation         = "service must not be an empty string"
	NoCmdValidation            = "command must be provided"
	EmptyCmdValidation         = "command must not be an empty string"
	EnvCmdTaskExecutingLog     = "Executing Command '%s' for environment '%s' ..."
	EnvCmdTaskResultLog        = "Result of Command '%s' for environment '%s' ...\n'%s"
	EnvCmdTaskErrorLog         = "The following error has occurred executing the command:  '%v'"
	EcsConnectionLog           = "Connecting to ECS service"
)

// Constants used during testing
const (
	EnvAliasCount    = 1
	SvcAliasCount    = 1
	FailExitCode     = 1
	Test             = "test"
	TestEnv          = "fooenv"
	TestSvc          = "foosvc"
	TestCmd          = "foocmd"
	Help             = "help"
	GetStackName     = "GetStack"
	RunTaskName      = "RunTask"
	NameMessage      = "Name should match"
	UsageMessage     = "Usage should match"
	AliasLenMessage  = "Aliases len should match"
	AliasMessage     = "Aliases should match"
	ArgsUsageMessage = "ArgsUsage should match"
	SubCmdLenMessage = "Subcommands len should match"
	FlagLenMessage   = "Flag len should match"
	FlagMessage      = "Flag should match"
)
