/*
 *
 * load.go
 * testcases
 *
 * Created by lintao on 2022/6/27 3:32 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package db

import (
	"github.com/NSObjects/scout/internal/waf/gotestwaf/config"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const testCaseExt = ".yml"

func LoadTestCases(cfg *config.Config) ([]Case, error) {
	var files []string
	var testCases []Case

	if cfg.TestCasesPath == "" {
		return nil, errors.New("empty test cases path")
	}

	if err := filepath.Walk(cfg.TestCasesPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	}); err != nil {
		return nil, err
	}

	for _, testCaseFile := range files {
		if filepath.Ext(testCaseFile) != testCaseExt {
			continue
		}

		// Ignore subdirectories, process as .../<testSetName>/<testCaseName>/<case>.yml
		parts := strings.Split(testCaseFile, string(os.PathSeparator))
		parts = parts[len(parts)-3:]

		testSetName := parts[1]
		testCaseName := strings.TrimSuffix(parts[2], testCaseExt)

		if cfg.TestSet != "" && testSetName != cfg.TestSet {
			continue
		}

		if cfg.TestCase != "" && testCaseName != cfg.TestCase {
			continue
		}

		yamlFile, err := ioutil.ReadFile(testCaseFile)
		if err != nil {
			return nil, err
		}

		var t Case
		err = yaml.Unmarshal(yamlFile, &t)
		if err != nil {
			return nil, err
		}

		t.Name = testCaseName
		t.Set = testSetName

		if strings.Contains(testSetName, "false") {
			t.IsTruePositive = false // test case is false positive
		} else {
			t.IsTruePositive = true // test case is true positive
		}

		testCases = append(testCases, t)
	}

	if testCases == nil {
		return nil, errors.New("no tests were selected")
	}

	return testCases, nil
}
