package appdirs

import (
	"fmt"
	"os/user"
	"runtime"
	"testing"
)

func getExpectations() map[string]string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	expectations := map[string]map[string]string{
		"darwin": {
			"UserData":   fmt.Sprintf("/Users/%s/Library/Application Support/Testing/1.0.0", u.Username),
			"UserCache":  fmt.Sprintf("/Users/%s/Library/Caches/Testing/1.0.0", u.Username),
			"UserConfig": fmt.Sprintf("/Users/%s/Library/Application Support/Testing/1.0.0", u.Username),
			"UserLog":    fmt.Sprintf("/Users/%s/Library/Logs/Testing/1.0.0", u.Username),
			"SiteData":   "/Library/Application Support/Testing/1.0.0",
			"SiteConfig": "/Library/Application Support/Testing/1.0.0",
		},
		"linux": {
			"UserData":   fmt.Sprintf("/home/%s/.local/share/Testing/1.0.0", u.Username),
			"UserCache":  fmt.Sprintf("/home/%s/.cache/Testing/1.0.0", u.Username),
			"UserConfig": fmt.Sprintf("/home/%s/.config/Testing/1.0.0", u.Username),
			"UserLog":    fmt.Sprintf("/home/%s/.cache/Testing/1.0.0/log", u.Username),
			"SiteData":   "/usr/local/share/Testing/1.0.0",
			"SiteConfig": "/etc/xdg/Testing/1.0.0",
		},
		"windows": {
			"UserData":   fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Tester\\Testing\\1.0.0", u.Username),
			"UserCache":  fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Tester\\Testing\\Cache\\1.0.0", u.Username),
			"UserConfig": fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Tester\\Testing\\1.0.0", u.Username),
			"UserLog":    fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Tester\\Testing\\1.0.0\\Logs", u.Username),
			"SiteData":   "C:\\ProgramData\\Tester\\Testing\\1.0.0",
			"SiteConfig": "C:\\ProgramData\\Tester\\Testing\\1.0.0",
		},
	}

	return expectations[runtime.GOOS]
}

func testApp() *App {
	app := New("Testing")
	app.Author = "Tester"
	app.Version = "1.0.0"
	return app
}

func TestUserData(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserData(); r != expected["UserData"] {
		t.Fatalf("Expected %s for UserData got: %s", expected["UserData"], r)
	}
}

func TestSiteData(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.SiteData(); r != expected["SiteData"] {
		t.Fatalf("Expected %s for SiteData got: %s", expected["SiteData"], r)
	}
}

func TestSiteConfig(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.SiteConfig(); r != expected["SiteConfig"] {
		t.Fatalf("Expected %s for SiteConfig got: %s", expected["SiteConfig"], r)
	}
}

func TestUserCache(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserCache(); r != expected["UserCache"] {
		t.Fatalf("Expected %s for UserCache got: %s", expected["UserCache"], r)
	}
}

func TestUserConfig(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserConfig(); r != expected["UserConfig"] {
		t.Fatalf("Expected %s for UserConfig got: %s", expected["UserConfig"], r)
	}
}

func TestUserLog(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserLog(); r != expected["UserLog"] {
		t.Fatalf("Expected %s for UserLog got: %s", expected["UserLog"], r)
	}
}
