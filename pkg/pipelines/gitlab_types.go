package pipelines

type GitlabPipeline struct {
	BeforeScript  []string                 `yaml:"before_script,omitempty"`
	AfterScript   []string                 `yaml:"after_script,omitempty"`
	Artifacts     *GitlabPipelineArtifacts `yaml:"artifacts,omitempty"`
	Cache         *GitlabPipelineCache     `yaml:"cache,omitempty"`
	Hooks         map[string][]string      `yaml:"hooks,omitempty"`
	Image         string                   `yaml:"image,omitempty"`
	Interruptible bool                     `yaml:"interruptible,omitempty"`
	Retry         int                      `yaml:"retry,omitempty"`
	Services      []*GitlabPipelineService `yaml:"services,omitempty"`
	Tags          []string                 `yaml:"tags,omitempty"`
	Timeout       string                   `yaml:"timeout,omitempty"`

	Include []*GitlabPipelineInclude `yaml:"include,omitempty"`

	Stages    []string                `yaml:"stages,omitempty"`
	Variables OBJECT_TYPE             `yaml:"variables,omitempty"`
	Workflow  *GitlabPipelineWorkflow `yaml:"workflow,omitempty"`
	Pages     *GitlabPipelineJob      `yaml:"pages,omitempty"`
}

type GitlabPipelineInclude struct {
	Remote  string      `yaml:"remote,omitempty"`
	Project string      `yaml:"project,omitempty"`
	Ref     string      `yaml:"ref,omitempty"`
	File    interface{} `yaml:"file,omitempty"`
}

type GitlabPipelineWorkflowRule struct {
	If        string        `yaml:"if,omitempty"`
	Variables OBJECT_TYPE   `yaml:"variables,omitempty"`
	When      WORKFLOW_WHEN `yaml:"when,omitempty"`
}

type GitlabPipelineWorkflow struct {
	Name  string `yaml:"name,omitempty"`
	Rules []*GitlabPipelineWorkflowRule
}

type GitlabPipelineDefault struct {
	BeforeScript  []string                 `yaml:"before_script,omitempty"`
	AfterScript   []string                 `yaml:"after_script,omitempty"`
	Artifacts     *GitlabPipelineArtifacts `yaml:"artifacts,omitempty"`
	Cache         *GitlabPipelineCache     `yaml:"cache,omitempty"`
	Hooks         map[string][]string      `yaml:"hooks,omitempty"`
	Image         string                   `yaml:"image,omitempty"`
	Interruptible bool                     `yaml:"interruptible,omitempty"`
	Retry         int                      `yaml:"retry,omitempty"`
	Services      []*GitlabPipelineService `yaml:"services,omitempty"`
	Tags          []string                 `yaml:"tags,omitempty"`
	Timeout       string                   `yaml:"timeout,omitempty"`
	Include       []*GitlabPipelineInclude `yaml:"include,omitempty"`
}

type GitlabPipelineArtifacts struct {
	Paths []string `yaml:"paths,omitempty"`
}

type GitlabPipelineService struct {
	Name       string   `yaml:"name,omitempty"`
	Alias      string   `yaml:"alias,omitempty"`
	Entrypoint []string `yaml:"entrypoint,omitempty"`
	Command    []string `yaml:"command,omitempty"`
}

type GitlabPipelineCache struct {
	Key   string   `yaml:"key,omitempty"`
	Paths []string `yaml:"paths,omitempty"`
}

type GitlabPipelineJob struct {
	Stage         string                                `yaml:"stage"`
	Parallel      int                                   `yaml:"parallel,omitempty"`
	AllowFailure  interface{}                           `yaml:"allow_failure,omitempty"`
	AfterScript   []string                              `yaml:"after_script,omitempty"`
	BeforeScript  []string                              `yaml:"before_script,omitempty"`
	Scripts       []string                              `yaml:"scripts,omitempty"`
	Artifacts     *GitlabPipelineJobArtifact            `yaml:"artifacts,omitempty"`
	Cache         *GitlabPipelineJobCache               `yaml:"cache,omitempty"`
	Coverage      string                                `yaml:"coverage,omitempty"`
	Script        string                                `yaml:"script,omitempty"`
	Environment   *GitlabPipelineJobEnvironment         `yaml:"environment,omitempty"`
	Hooks         *GitlabPipelineJobHook                `yaml:"hooks,omitempty"`
	IdTokens      map[string]*GitlabPipelineIdToken     `yaml:"id_tokens,omitempty"`
	Image         *GitlabPipelineJobImage               `yaml:"image,omitempty"`
	Inherit       *GitlabPipelineJobInherit             `yaml:"inherit,omitempty"`
	Interruptible bool                                  `yaml:"interruptible,omitempty"`
	Release       *GitlabPipelineJobRelease             `yaml:"release,omitempty"`
	ResourceGroup string                                `yaml:"resource_group,omitempty"`
	Needs         []*GitlabPipelineJobNeed              `yaml:"needs,omitempty"`
	Rules         []*GitlabPipelineJobRule              `yaml:"rules,omitempty"`
	Secrets       map[string]*GitlabPipelineJobSecret   `yaml:"secrets,omitempty"`
	Tags          []string                              `yaml:"tags,omitempty"`
	Timeout       string                                `yaml:"timeout,omitempty"`
	Trigger       *GitlabPipelineJobTrigger             `yaml:"trigger,omitempty"`
	When          JOB_WHEN                              `yaml:"when,omitempty"`
	Variables     map[string]*GitlabPipelineJobVariable `yaml:"variables,omitempty"`
	Services      []*GitlabPipelineJobServices          `yaml:"services,omitempty"`
	Retry         *GitlabPipelineJobRetry               `yaml:"retry,omitempty"`
	Only          *GitlabPipelineJobOnly                `yaml:"only,omitempty"`
	Except        []string                              `yaml:"except,omitempty"`
}

type GitlabPipelineJobVariable struct {
	Options     []string    `yaml:"options,omitempty"`
	Expand      bool        `yaml:"expand,omitempty"`
	Value       interface{} `yaml:"value,omitempty"`
	Description string      `yaml:"description,omitempty"`
}

type GitlabPipelineJobRetry struct {
	Max  int              `yaml:"max,omitempty"`
	When []JOB_RETRY_WHEN `yaml:"when,omitempty"`
}

type GitlabPipelineJobOnly struct {
	Refs       []string `yaml:"refs,omitempty"`
	Variables  []string `yaml:"variables,omitempty"`
	Kubernetes string   `yaml:"kubernetes,omitempty"`
	Changes    []string `yaml:"changes,omitempty"`
}

type GitlabPipelineJobServices struct {
	Name       string                `yaml:"name,omitempty"`
	PullPolicy []SERVICE_PULL_POLICY `yaml:"pull_policy,omitempty"`
	EntryPoint []string              `yaml:"entrypoint,omitempty"`
	Command    []string              `yaml:"command,omitempty"`
}

type GitlabPipelineJobTrigger struct {
	Include  string                           `yaml:"include,omitempty"`
	Project  string                           `yaml:"project,omitempty"`
	Branch   string                           `yaml:"branch,omitempty"`
	Strategy string                           `yaml:"strategy,omitempty"`
	Forward  *GitlabPipelineJobTriggerForward `yaml:"forward,omitempty"`
}

type GitlabPipelineJobTriggerForward struct {
	PipelineVariables bool `yaml:"pipeline_variable,omitempty"`
	YamlVariables     bool `yaml:"yaml_variables,omitempty"`
}

type GitlabPipelineJobSecret struct {
	File  bool                          `yaml:"file,omitempty"`
	Token string                        `yaml:"token,omitempty"`
	Vault *GitlabPipelineJobSecretVault `yaml:"vault,omitempty"`
}

type GitlabPipelineJobSecretVault struct {
	Engine *GitlabPipelineJobSecretVaultEngine `yaml:"engine,omitempty"`
	Path   string                              `yaml:"path,omitempty"`
	Field  string                              `yaml:"field,omitempty"`
}

type GitlabPipelineJobSecretVaultEngine struct {
	Name string `yaml:"name,omitempty"`
	Path string `yaml:"path,omitempty"`
}

type GitlabPipelineJobRule struct {
	AllowFailure bool                          `yaml:"allow_failure,omitempty"`
	Exists       []string                      `yaml:"exists,omitempty"`
	If           string                        `yaml:"if,omitempty"`
	Needs        []string                      `yaml:"needs,omitempty"`
	Changes      *GitlabPipelineJobRuleChanges `yaml:"changes,omitempty"`
	Variables    OBJECT_TYPE                   `yaml:"variables,omitempty"`
	When         JOB_RULES_WHEN                `yaml:"when,omitempty"`
}
type GitlabPipelineJobRuleChanges struct {
	Paths     []string `yaml:"paths,omitempty"`
	CompareTo string   `yaml:"compare_to,omitempty"`
}
type GitlabPipelineJobNeed struct {
	Pipeline  string `yaml:"pipeline,omitempty"`
	Job       string `yaml:"job,omitempty"`
	Ref       string `yaml:"ref,omitempty"`
	Artifacts bool   `yaml:"artifacts,omitempty"`
	Optional  bool   `yaml:"optional,omitempty"`
}

type GitlabPipelineIdToken struct {
	Aud []string `yaml:"aud,omitempty"`
}

type GitlabPipelineJobHook struct {
	PreGetSourcesScript []string `yaml:"pre_get_sources_script,omitempty"`
}

type GitlabPipelineJobRelease struct {
	TagName     string                          `yaml:"tag_name,omitempty"`
	Name        string                          `yaml:"name,omitempty"`
	Description string                          `yaml:"description,omitempty"`
	TagMessage  string                          `yaml:"tag_message,omitempty"`
	ReleasedAt  string                          `yaml:"released_at,omitempty"`
	Milestones  []string                        `yaml:"milestones,omitempty"`
	Ref         string                          `yaml:"ref,omitempty"`
	Assets      *GitlabPipelineJobReleaseAssets `yaml:"assets,omitempty"`
}

type GitlabPipelineJobReleaseAssets struct {
	Links []*GitlabPipelineJobReleaseAssetsLinks `yaml:"links,omitempty"`
}

type GitlabPipelineJobReleaseAssetsLinks struct {
	Name     string `yaml:"name,omitempty"`
	Url      string `yaml:"url,omitempty"`
	Filepath string `yaml:"filepath,omitempty"`
	LinkType string `yaml:"link_type,omitempty"`
}

type GitlabPipelineJobInherit struct {
	Variable bool `yaml:"variable,omitempty"`
}

type GitlabPipelineJobImage struct {
	Name       string              `yaml:"name,omitempty"`
	Entrypoint []string            `yaml:"entrypoint,omitempty"`
	PullPolicy []IMAGE_PULL_POLICY `yaml:"pull_policy,omitempty"`
}

type GitlabPipelineJobEnvironment struct {
	Name       string `yaml:"name,omitempty"` // only required field
	Url        string `yaml:"url,omitempty"`
	Action     string `yaml:"action,omitempty"`
	AutoStopIn string `yaml:"auto_stop_in,omitempty"`
	Kubernets  struct {
		Namespace string `yaml:"namespace,omitempty"`
	} `yaml:"kubernetes,omitempty"`
	DeploymentTier string `yaml:"deployment_tier,omitempty"`
	Extends        string `yaml:"extends,omitempty"`
}

type GitlabPipelineJobCache struct {
	Key          interface{}      `yaml:"key,omitempty"`
	Untracked    bool             `yaml:"untracked,omitempty"`
	Unprotected  bool             `yaml:"unprotected,omitempty"`
	When         JOB_CACHE_WHEN   `yaml:"when,omitempty"`
	Policy       JOB_CACHE_POLICY `yaml:"policy,omitempty"`
	Paths        []string         `yaml:"paths,omitempty"`
	FallbackKeys []string         `yaml:"fallback_keys,omitempty"`
}

type GitlabPipelineJobArtifact struct {
	Paths    []string          `yaml:"paths,omitempty"`
	Exclude  []string          `yaml:"exclude,omitempty"`
	ExpireIn string            `yaml:"expire_in,omitempty"`
	ExposeAs string            `yaml:"expose_as,omitempty"`
	Name     string            `yaml:"name,omitempty"`
	Public   bool              `yaml:"public,omitempty"`
	Reports  map[string]string `yaml:"reports,omitempty"`
	When     JOB_ARTIFACT_WHEN `yaml:"when,omitempty"`
}

type IMAGE_PULL_POLICY = string

var IMAGE_PULL_POLICIES = []IMAGE_PULL_POLICY{
	"always",
	"if-not-present",
}

type JOB_RETRY_WHEN = string

var JOB_RETRY_WHENS = []JOB_RETRY_WHEN{
	"runner_system_failure",
	"stuck_or_timeout_failure",
	"always",
	"unknown_failure",
	"script_failure",
	"api_failure",
	"runner_system_failure",
	"runner_unsupported",
	"stale_schedule",
	"job_execution_timeout",
	"archived_failure",
	"unmet_prerequisites",
	"scheduler_failure",
	"dat_integrity_failure",
}

type SERVICE_PULL_POLICY = string

var SERVICE_PULL_POLICIES = []SERVICE_PULL_POLICY{
	"always",
	"if-not-present",
}

type JOB_RULES_WHEN = string

var JOB_RULES_WHENS = []string{
	"always",
	"on_success", // default
	"delayed",
}

type JOB_WHEN = string

var JOB_WHENS = []JOB_WHEN{
	"on_failure",
	"on_success",
	"never",
	"always",
	"manual",
	"delayed",
	"always",
}

type JOB_CACHE_POLICY = string

var JOB_CACHE_POLICIES = []JOB_CACHE_POLICY{
	"pull",
	"push",
	"pull-push",
}

type JOB_CACHE_WHEN = string

var JOB_CACHE_WHENS = []JOB_CACHE_WHEN{
	"on_failure",
	"on_success",
	"always",
}

type JOB_ARTIFACT_WHEN = string

var JOB_ARTIFACT_WHENS = []JOB_ARTIFACT_WHEN{
	"on_failure",
	"on_success",
	"always",
}

type WORKFLOW_WHEN = string

var WORKFLOW_WHENS = []WORKFLOW_WHEN{
	"always",
	"never",
}
