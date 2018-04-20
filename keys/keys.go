package keys

import (
	"github.com/kklash/gogpg/execution"
	"regexp"
	"strings"
)

const APP string = "gpg"

func GetUserIDsSecret() []string {
	process := execution.Command{
		App:  APP,
		Args: []string{"--with-colons", "--list-secret-keys"},
	}
	output, _ := process.Execute()
	lines := strings.Split(output, "\n")

	ids := make([]string, 0, 10)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if !strings.HasPrefix(line, "uid:") {
			continue
		}
		uid := strings.Split(line, ":")[9]
		if hasEmail, _ := regexp.MatchString("<.+>", uid); !hasEmail {
			ids = append(ids, uid)
			continue
		}
		re, _ := regexp.Compile("(.*) <(.+)>")
		matches := re.FindStringSubmatch(uid)
		ids = append(ids, matches[1], matches[2])
	}
	return ids
}

func SearchKey(uid string, list []string) bool {
	for _, key := range list {
		if strings.Contains(key, uid) {
			return true
		}
	}
	return false
}
