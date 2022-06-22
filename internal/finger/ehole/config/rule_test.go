/*
 *
 * rule_test.go
 * finger
 *
 * Created by lintao on 2022/6/14 4:03 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package config

import (
	"testing"
)

func TestWebFingerRule(t *testing.T) {
	tests := []struct {
		name    string
		want    *Rule
		wantErr bool
	}{
		{
			name:    "init",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := WebFingerRule()
			if (err != nil) != tt.wantErr {
				t.Errorf("WebFingerRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
