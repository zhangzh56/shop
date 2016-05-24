package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"time"

	"github.com/skratchdot/open-golang/open"
)

// Returns string in format YYYYMMDD
func GetDateYYYYMMDD() string {
	return time.Now().Format("20060102")
}

func GetFormattedTime(t time.Time) string {
	return t.Format("Mon Jan 2 2006 15:04:05")
}
func GetFormattedTimeShort(t time.Time) string {
	return t.Format("Mon Jan 2 2006 150405")
}

func TimeNow() time.Time {
	// Trunacte makes sure that the exact same value that goes in the db,
	// will come out of it again (otherwise mongo will cut precison and persistence tests will fail)
	return time.Now().Truncate(time.Second)
}

func PrintJSON(v interface{}) {
	fmt.Println(ToJSON(v))
}

// ToJSON creates a JSON from v and returns it as string
func ToJSON(v interface{}) string {
	json, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return "Could not print JSON"
	}
	return string(json)
}

func GetFunctionName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// OpenInBrowser stores the html under "<name><date>.html" in the current
// directory and opens it in the browser
func OpenInBrowser(name string, html string) error {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("PWD", pwd)

	//path := "file://" + pwd + "/" + "diff.html"
	path := pwd + "/" + "diff-" + name + "_" + GetFormattedTimeShort(time.Now()) + ".html"
	fmt.Println("PATH", path)
	tmpFile, err := os.Create(path)
	if err != nil {
		return err
	}
	tmpFile.WriteString(html)
	tmpFile.Close()
	err = open.Start(path)
	if err != nil {
		log.Println(err)
	}
	return err

}
