package models

type Commit struct {
	Type            string
	Scope           string
	Summary         string
	Details         string
	BreakingChange  bool
	BreakingMessage string
	Footer          string
}
type ShellCommands struct {
	Name      string
	Arguments []string
}

const (
	LINE = "------------------------------------------------"
)

var AllCommands = map[string]ShellCommands{
	"gitBranch": {
		Name:      "git",
		Arguments: []string{"branch", "--show-current"},
	},
	"gitFetchBranch": {
		Name:      "git",
		Arguments: []string{"branch", "--list"},
	},
	"gitCommit": {
		Name:      "git",
		Arguments: []string{"commit", "-F", "-"},
	},
	"gitAdd": {
		Name:      "git",
		Arguments: []string{"add"},
	},
	"gitPush": {
		Name:      "git",
		Arguments: []string{"push", "origin", "HEAD"},
	},
	"gitReset": {
		Name:      "git",
		Arguments: []string{"reset", "--mixed"},
	},
	"gitPull": {
		Name:      "git",
		Arguments: []string{"pull", "origin"},
	},
	"gitStatus": {
		Name:      "git",
		Arguments: []string{"status", "--porcelain"},
	},
	"gitCheckout": {
		Name:      "git",
		Arguments: []string{"checkout"},
	},
	"gitMerge": {
		Name:      "git",
		Arguments: []string{"merge"},
	},
	"createBranch": {
		Name:      "git",
		Arguments: []string{"branch"},
	},
}
