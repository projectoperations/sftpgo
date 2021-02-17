package version

import "strings"

const version = "2.0.2"

var (
	commit = ""
	date   = ""
	info   Info
)

// Info defines version details
type Info struct {
	Version    string   `json:"version"`
	BuildDate  string   `json:"build_date"`
	CommitHash string   `json:"commit_hash"`
	Features   []string `json:"features"`
}

// GetAsString returns the string representation of the version
func GetAsString() string {
	var sb strings.Builder
	sb.WriteString(info.Version)
	if len(info.CommitHash) > 0 {
		sb.WriteString("-")
		sb.WriteString(info.CommitHash)
	}
	if len(info.BuildDate) > 0 {
		sb.WriteString("-")
		sb.WriteString(info.BuildDate)
	}
	if len(info.Features) > 0 {
		sb.WriteString(" ")
		sb.WriteString(strings.Join(info.Features, " "))
	}
	return sb.String()
}

func init() {
	info = Info{
		Version:    version,
		CommitHash: commit,
		BuildDate:  date,
	}
}

// AddFeature adds a feature description
func AddFeature(feature string) {
	info.Features = append(info.Features, feature)
}

// Get returns the Info struct
func Get() Info {
	return info
}
