package pipelines

type GithubPipeline struct {
	Name        string          `yaml:"name,omitempty"`
	RunName     string          `yaml:"run-name,omitempty,omitempty"`
	Permissions interface{}     `yaml:"permissions,omitempty"`
	Env         OBJECT_TYPE     `yaml:"env,omitempty"`
	Defaults    *GithubDefaults `yaml:"defaults,omitempty"`
	// concurrency: staging_environment
	// concurrency:
	//   group: ${{ github.ref }}
	//   cancel-in-progress: true
	Concurrency interface{}                    `yaml:"concurrency,omitempty"`
	On          interface{}                    `yaml:"on,omitempty"`
	Jobs        map[string]*GithubPipelineJobs `yaml:"jobs,omitempty"`
}
type GithubPipelineConcurrency struct {
	Group            string `yaml:"group,omitempty"`
	CancelInProgress bool   `yaml:"cancel-in-progress,omitempty"`
}
type GithubPipelinePermissions struct {
	Actions            string `yaml:"actions,omitempty"`
	Checks             string `yaml:"checks,omitempty"`
	Contents           string `yaml:"contents,omitempty"`
	Deployments        string `yaml:"deployments,omitempty"`
	IdToken            string `yaml:"id-token,omitempty"`
	Issues             string `yaml:"issues,omitempty"`
	Discussions        string `yaml:"discussions,omitempty"`
	Packages           string `yaml:"packages,omitempty"`
	Pages              string `yaml:"pages,omitempty"`
	PullRequests       string `yaml:"pull-requests,omitempty"`
	RepositoryProjects string `yaml:"repository-projects,omitempty"`
	SecurityEvents     string `yaml:"security-events,omitempty"`
	Statuses           string `yaml:"statuses,omitempty"`
}

type GithubPipelineOnConfigSecret struct {
	Description string `yaml:"description,omitempty"`
	Required    string `yaml:"required,omitempty"`
}

type GithubPipelineOnConfigOutput struct {
	Description string `yaml:"description,omitempty"`
	Value       string `yaml:"value,omitempty"`
}

type GithubPipelineOnConfigInput struct {
	Description string   `yaml:"description,omitempty"`
	Required    bool     `yaml:"required,omitempty"`
	Default     string   `yaml:"default,omitempty"`
	Type        string   `yaml:"type,omitempty"`
	Options     []string `yaml:"options,omitempty"`
	Value       string   `yaml:"value,omitempty"`
}

type GithubPipelineOnConfig struct {
	Types          []string                                 `yaml:"types,omitempty"`
	PathsIgnore    []string                                 `yaml:"path-ignore,omitempty"`
	Paths          []string                                 `yaml:"paths,omitempty"`
	Tags           []string                                 `yaml:"tags,omitempty"`
	TagsIgnore     []string                                 `yaml:"tags-ignore,omitempty"`
	Branches       []string                                 `yaml:"branches,omitempty"`
	BranchesIgnore []string                                 `yaml:"branches-ignore,omitempty"`
	Workflows      []string                                 `yaml:"workflows,omitempty"`
	Secrets        map[string]*GithubPipelineOnConfigSecret `yaml:"secrets,omitempty"`
	Outputs        map[string]*GithubPipelineOnConfigOutput `yaml:"outputs,omitempty"`
	Inputs         map[string]*GithubPipelineOnConfigInput  `yaml:"inputs,omitempty"`
}

type GithubPipelineOnConfigSchedule struct {
	Cron string `yaml:"cron,omitempty"`
}

type GithubPipelineOn struct {
	PullRequestComment       *GithubPipelineOnConfig           `yaml:"pull_request_comment,omitempty"`
	PullRequestReviewComment *GithubPipelineOnConfig           `yaml:"pull_request_review_comment,omitempty"`
	PullRequestTarget        *GithubPipelineOnConfig           `yaml:"pull_request_target,omitempty"`
	BranchProtectionRule     *GithubPipelineOnConfig           `yaml:"branch_protection_rule,omitempty"`
	Label                    *GithubPipelineOnConfig           `yaml:"label,omitempty"`
	CheckSuite               *GithubPipelineOnConfig           `yaml:"check_suite,omitempty"`
	CheckRun                 *GithubPipelineOnConfig           `yaml:"check_run,omitempty"`
	Create                   OBJECT_TYPE                       `yaml:"create,omitempty"`
	Delete                   OBJECT_TYPE                       `yaml:"delete,omitempty"`
	DeploymentStatus         OBJECT_TYPE                       `yaml:"deployment_status,omitempty"`
	Discussion               *GithubPipelineOnConfig           `yaml:"discussion,omitempty"`
	DiscussionComment        *GithubPipelineOnConfig           `yaml:"discussion_comment,omitempty"`
	Fork                     OBJECT_TYPE                       `yaml:"fork,omitempty"`
	Gollum                   OBJECT_TYPE                       `yaml:"gollum,omitempty"`
	PageBuild                OBJECT_TYPE                       `yaml:"page_build,omitempty"`
	Status                   OBJECT_TYPE                       `yaml:"status,omitempty"`
	Public                   OBJECT_TYPE                       `yaml:"public,omitempty"`
	IssueComment             *GithubPipelineOnConfig           `yaml:"issue_comment,omitempty"`
	Issues                   *GithubPipelineOnConfig           `yaml:"issues,omitempty"`
	MergeGroup               *GithubPipelineOnConfig           `yaml:"merge_group,omitempty"`
	Milestone                *GithubPipelineOnConfig           `yaml:"milestone,omitempty"`
	Schedule                 []*GithubPipelineOnConfigSchedule `yaml:"schedule,omitempty"`
	Watch                    *GithubPipelineOnConfig           `yaml:"watch,omitempty"`
	WorkflowRun              *GithubPipelineOnConfig           `yaml:"workflow_run,omitempty"`
	WorkflowDispatch         *GithubPipelineOnConfig           `yaml:"workflow_dispatch,omitempty"`
	Project                  *GithubPipelineOnConfig           `yaml:"project,omitempty"`
	ProjectCard              *GithubPipelineOnConfig           `yaml:"project_card,omitempty"`
	ProjectColumn            *GithubPipelineOnConfig           `yaml:"project_column,omitempty"`
	RepositoryDispatch       *GithubPipelineOnConfig           `yaml:"repository_dispatch,omitempty"`
	PullRequest              *GithubPipelineOnConfig           `yaml:"pull_request,omitempty"`
	Push                     *GithubPipelineOnConfig           `yaml:"push,omitempty"`
	RegistryPackage          *GithubPipelineOnConfig           `yaml:"registry_package,omitempty"`
	Release                  *GithubPipelineOnConfig           `yaml:"release,omitempty"`
}
type GithubPipelineJobs struct {
	Name string                 `yaml:"name,omitempty"`
	Uses string                 `yaml:"uses,omitempty"`
	With map[string]interface{} `yaml:"with,omitempty"`
	// secrets: inherit
	// secrets:
	//   access-token: ${{ secrets.PERSONAL_ACCESS_TOKEN }}
	Secrets         interface{}                           `yaml:"secrets,omitempty"`
	Services        map[string]*GithubPipelineJobsService `yaml:"services,omitempty"`
	Container       *GithubPipelineContainer              `yaml:"container,omitempty"`
	FailFast        bool                                  `yaml:"fail-fast,omitempty"`
	ContinueOnError bool                                  `yaml:"continue-on-error,omitempty"`
	MaxParallel     int                                   `yaml:"max-parallel,omitempty"`
	If              string                                `yaml:"if,omitempty"`
	Environment     *GithubPipelineJobsEnvironment        `yaml:"environment,omitempty"`
	// runs-on: ubuntu-latest|windows-latest|windows-2022|windows-2019|ubuntu-22.04|macos-13|macos-13-xl|macos-latest|macos-12|macos-latest-xl|macos-12-xl|macos-11
	// runs-on:
	//   group: ubuntu-runners
	//   labels: ubuntu-20.04-16core
	RunsOn         interface{}            `yaml:"runs-on,omitempty"`
	Outputs        OBJECT_TYPE            `yaml:"outputs,omitempty"`
	Env            OBJECT_TYPE            `yaml:"env,omitempty"`
	Defaults       GithubPipelineDefaults `yaml:"defaults,omitempty"`
	Needs          []string               `yaml:"needs,omitempty"`
	TimeoutMinutes int                    `yaml:"timeout-minutes,omitempty"`
	// concurrency: staging_environment
	// concurrency:
	//   group: ${{ github.ref }}
	//   cancel-in-progress: true
	Concurrency interface{}            `yaml:"concurrency,omitempty"`
	Strategy    GithubPipelineStrategy `yaml:"strategy,omitempty"`
	Permissions interface{}            `yaml:"permissions,omitempty"`
	Steps       []*GithubPipelineStep  `yaml:"steps,omitempty"`
}

type GithubPipelineJobsEnvironment struct {
	Name string `yaml:"name,omitempty"`
	Url  string `yaml:"url,omitempty"`
}

type GithubPipelineDefaults struct {
	Run *GithubPipelineDefaultsRun `yaml:"run,omitempty"`
}
type GithubPipelineJobsService struct {
	Env         map[string]interface{}              `yaml:"env,omitempty"`
	Volumes     []string                            `yaml:"volumes,omitempty"`
	Options     string                              `yaml:"options,omitempty"`
	Image       string                              `yaml:"image,omitempty"`
	Ports       []string                            `yaml:"ports,omitempty"`
	Credentials *GithubPipelineContainerCredentials `yaml:"credentials,omitempty"`
}

type GithubPipelineDefaultsRun struct {
	Shell            SHELL_TYPE `yaml:"shell,omitempty"`
	WorkingDirectory string     `yaml:"working-directory,omitempty"`
}

type GithubPipelineStrategy struct {
	Matrix map[string]interface{} `yaml:"matrix,omitempty"`
}

type GithubPipelineStep struct {
	Name             string      `yaml:"name,omitempty"`
	ContinueOnError  bool        `yaml:"continue-on-error,omitempty"`
	TimeoutMinutes   int         `yaml:"timeout-minutes,omitempty"`
	Env              OBJECT_TYPE `yaml:"env,omitempty"`
	Run              string      `yaml:"run,omitempty"`
	Id               string      `yaml:"id,omitempty"`
	WorkingDirectory string      `yaml:"working-directory,omitempty"`
	// bash, pwsh, python, sh, cmd, powershell or a custom string with {0} e.g perl {0}
	Shell SHELL_TYPE  `yaml:"shell,omitempty"`
	If    string      `yaml:"if,omitempty"`
	Uses  string      `yaml:"uses,omitempty"`
	With  OBJECT_TYPE `yaml:"with,omitempty"`
}

type GithubPipelineContainer struct {
	Env         OBJECT_TYPE                         `yaml:"env,omitempty"`
	Volumes     []string                            `yaml:"volumes,omitempty"`
	Options     string                              `yaml:"options,omitempty"`
	Image       string                              `yaml:"image,omitempty"`
	Ports       []string                            `yaml:"ports,omitempty"`
	Credentials *GithubPipelineContainerCredentials `yaml:"credentials,omitempty"`
}

type GithubPipelineContainerCredentials struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

type GithubPipelinePermissionsVariants struct {
	PermissionsAsString string
	PermissionsAsStruct *GithubPermissions
}

type GithubPermissions struct {
	Actions            string `yaml:"actions,omitempty"`
	Checks             string `yaml:"checks,omitempty"`
	Contents           string `yaml:"contents,omitempty"`
	Deployments        string `yaml:"deployments,omitempty"`
	IDTokens           string `yaml:"id-tokens,omitempty"`
	Issues             string `yaml:"issues,omitempty"`
	Discussions        string `yaml:"discussions,omitempty"`
	Packages           string `yaml:"packages,omitempty"`
	PullRequests       string `yaml:"pull-requests,omitempty"`
	RepositoryProjects string `yaml:"repository-projects,omitempty"`
	SecurityEvents     string `yaml:"security-events,omitempty"`
	Statuses           string `yaml:"statuses,omitempty"`
}

type GithubRun struct {
	Shell            SHELL_TYPE `yaml:"shell,omitempty"`
	WorkingDirectory string     `yaml:"working-directory,omitempty"`
}

type SHELL_TYPE = string

var BASH_TYPES = []SHELL_TYPE{
	"bash", "pwsh", "python", "sh", "cmd", "powershell",
}

type EVENT_NAMES = string

var EVENTS = []EVENT_NAMES{
	"pull_request_comment",
	"pull_request_review_comment",
	"pull_request_target",
	"branch_protection_rule",
	"label",
	"check_run",
	"check_suite",
	"create",
	"delete",
	"deployment_status",
	"discussion",
	"discussion_comment",
	"fork",
	"gollum",
	"issue_comment",
	"issues",
	"merge_group",
	"milestone",
	"schedule",
	"workflow_run",
	"watch",
	"workflow_dispatch",
	"page_build",
	"status",
	"project",
	"project_card",
	"project_column",
	"public",
	"repository_dispatch",
	"pull_request",
	"push",
	"registry_package",
	"release",
}

type EVENT_TYPES = []EVENT_NAMES

var EVENTS_WITH_TYPES = map[EVENT_NAMES]EVENT_TYPES{
	"release": EVENT_TYPES{
		"published",
		"unpublished",
		"created",
		"edited",
		"deleted",
		"prereleased",
		"released",
	},
	"registry_package": EVENT_TYPES{
		"published",
		"updated",
	},
	"repository_dispatch": EVENT_TYPES{},
	"project_column": EVENT_TYPES{
		"created", "updated", "moved", "deleted",
	},
	"project_card": EVENT_TYPES{
		"created", "moved", "converted", "edited", "deleted",
	},
	"project": EVENT_TYPES{
		"created", "closed", "reopened", "edited", "deleted",
	},
	"watch": EVENT_TYPES{
		"starred",
	},
	"workflow_run": EVENT_TYPES{
		"completed", "requested", "in_progress",
	},
	"milestone": EVENT_TYPES{
		"created", "closed", "opened", "edited", "deleted",
	},
	"merge_group": EVENT_TYPES{
		"checks_requseted",
	},
	"issues": EVENT_TYPES{
		"opened", "edited", "deleted", "transferred", "pinned", "unpinned", "closed", "reopened", "assigned", "unassigned", "labeled", "unlabeled", "locked", "unlocked", "milestoned", "demilestoned",
	},
	"discussion_comment": EVENT_TYPES{
		"created", "edited", "deleted",
	},
	"check_suite": EVENT_TYPES{
		"completed",
	},
	"check_run": EVENT_TYPES{
		"rerequested", "completed", "created", "requested_action",
	},
	"label": EVENT_TYPES{
		"created",
	},
	"branch_protection_rule": EVENT_TYPES{
		"created", "deleted", "edited",
	},
	"pull_request_target": EVENT_TYPES{
		"assigned", "unassigned", "labeled", "unlabeled", "opened", "edited", "closed", "reopened", "synchronize", "converted_to_draft", "ready_for_review", "locked", "unlocked", "review_requested", "review_request_removed", "auto_merge_enabled", "auto_merge_disabled",
	},
	"pull_request_review_comment": EVENT_TYPES{
		"created", "edited", "deleted",
	},
	"pull_request_comment": EVENT_TYPES{
		"submitted", "edited", "dismissed",
	},
}

type GithubDefaults struct {
	Run *GithubRun `yaml:"run,omitempty"`
}

type GithubConcurrency struct {
	Group            string `yaml:"group,omitempty"`
	CancelInProgress bool   `yaml:"cancel-in-progress,omitempty"`
}
