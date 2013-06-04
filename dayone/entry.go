package dayone

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"os"
	"strings"
	"text/template"
	"time"
)

const version = "0.0.1"

func NewEntry(text string) (entry *Entry) {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	now := time.Now()

	return &Entry{text, id, now}
}

type Entry struct {
	Text string
	id   *uuid.UUID
	time time.Time
}

func (self *Entry) Time() string {
	return self.time.UTC().Format(timeFormat)
}

func (self *Entry) Id() string {
	return strings.ToUpper(strings.Replace(self.id.String(), "-", "", -1))
}

func (self *Entry) WriteIn(dir string) {
	os.MkdirAll(dir, 0755)

	id := self.Id()
	path := fmt.Sprintf("%s/%s.doentry", dir, id)

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	tmpl, err := template.New("entry").Parse(entryTemplate)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(file, self)
	if err != nil {
		panic(err)
	}
}

// This time format contains no time zone indicator.
// So if you parse with this format you will get a time in UTC.
const timeFormat = "2006-01-02T15:04:05Z"

const entryTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Creation Date</key>
	<date>{{.Time}}</date>
	<key>Entry Text</key>
	<string>{{html .Text}}</string>
	<key>Starred</key>
	<false/>
	<key>Time Zone</key>
	<string>Asia/Tokyo</string>
	<key>UUID</key>
	<string>{{.Id}}</string>
</dict>
</plist>
`
