package history

import (
	"encoding/json"
	"time"

	"github.com/foomo/shop/utils"
	"github.com/sergi/go-diff/diffmatchpatch"
)

type Version struct {
	Number    int
	TimeStamp time.Time
}

func (v *Version) Increment() {
	v.Number = v.Number + 1
	v.TimeStamp = time.Now()
}

func (v *Version) GetVersion() int {
	return v.Number
}
func (v *Version) GetFormattedTime() string {
	return utils.GetFormattedTime(v.TimeStamp)
}

func DiffVersions(versionA interface{}, versionB interface{}) (string, error) {
	jsonA, err := json.MarshalIndent(versionA, "", "	")
	if err != nil {
		return "", err
	}
	jsonB, err := json.MarshalIndent(versionB, "", "	")
	if err != nil {
		return "", err
	}
	d := diffmatchpatch.New()
	diffs := d.DiffMain(string(jsonA), string(jsonB), false)

	html := d.DiffPrettyHtml(diffs)
	return html, nil
}
