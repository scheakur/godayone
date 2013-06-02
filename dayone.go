package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"os"
	"strings"
	"text/template"
	"time"
)

const version = "0.0.1"

func main() {
	argv := os.Args
	argc := len(argv)

	if argc < 2 {
		usage()
	}

	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	now := time.Now()
	text := argv[1]

	entry := &Entry{text, id, now}

	var dir string
	if argc == 3 {
		dir = argv[2]
	} else {
		dir = os.ExpandEnv("$HOME/Dropbox/Apps/Day One/Journal.dayone/entries")
	}

	entry.Write(dir)
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: godayone text [dir]
Version %s
`, version)
	os.Exit(-1)
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

func (self *Entry) Write(dir string) {
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
