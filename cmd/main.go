package main

import (
	"fmt"
	"log"
	"net/url"
	"regexp"

	"github.com/pkg/browser"
)

var (
	// tag must specify major.minor.patch or major.minor.patch-rc#
	// (e.g. 1.12.1, 1.12.1-rc1)
	validTag = regexp.MustCompile(`^\d+\.\d+\.\d+(-rc\d+)?$`)
	// target must specify release branch
	validTarget = regexp.MustCompile(`^release-\d+$`)
	// true or false
	validPrerelease = regexp.MustCompile(`^{true|false}$`)
)

func main() {
	fmt.Println("Start creating new release interactively!\n")

	var tag string
	fmt.Println("1. Input version tag with the format {major}.{minor}.{patch} or {major}.{minor}.{patch}-{rc#})")
	fmt.Scanf("%s", &tag)

	if !validTag.Match([]byte(tag)) {
		fmt.Println("Version tag must follow the format {major}.{minor}.{patch} or {major}.{minor}.{patch}-{rc#})")
	}

	var target string
	fmt.Println("2. Input target with the format release-#")
	fmt.Scanf("%s", &target)

	if !validTarget.Match([]byte(target)) {
		fmt.Println("Target must follow the format release-#")
	}

	var preRelease string
	fmt.Println("3. Is prerelease? (Set true for release candidates) (true/false)")
	fmt.Scanf("%s", &preRelease)

	if !validPrerelease.Match([]byte(preRelease)) {
		fmt.Println("Prerelease must be true or false")
	}

	title := fmt.Sprintf("Release %s", tag)

	// https://docs.github.com/en/repositories/releasing-projects-on-github/automation-for-release-forms-with-query-parameters
	url, err := url.Parse(
		fmt.Sprintf("https://github.com/takanabe/github-releases-test/releases/new?tag=%s&target=%s&title=%s&prerelease=%s", tag, target, title, preRelease),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := browser.OpenURL(url.String()); err != nil {
		log.Fatal(err)
	}
}
