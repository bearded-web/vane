package checks

import (
	"os"
)

var (
	readmefixtures  = os.Getenv("FIXTURESPATH") + "/readme/"
	robotsfixtures  = os.Getenv("FIXTURESPATH") + "/robotstxt/"
	rssfixtures     = os.Getenv("FIXTURESPATH") + "/rss_url/"
	has_logfixtures = os.Getenv("FIXTURESPATH") + "/has_log/"

	fakeHTTPaddress = "http://127.0.0.1:9999"
)
