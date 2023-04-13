package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFormat(t *testing.T) {
	testCases := []struct {
		Name    string
		Message string
		want    Format
		wantErr error
	}{
		{
			Name:    "happy path",
			Message: "feat(test): samples",
			want: Format{
				Type:    "feat",
				Scope:   "test",
				Subject: "samples",
			},
			wantErr: nil,
		},
		{
			Name:    "happy path 2",
			Message: "feat(test):                         samples",
			want: Format{
				Type:    "feat",
				Scope:   "test",
				Subject: "samples",
			},
			wantErr: nil,
		},
		{
			Name:    "invalid format",
			Message: "feat(test):samples",
			wantErr: ErrFormat,
		},
		{
			Name:    "scope empty",
			Message: "feat: 调整",
			want: Format{
				Type:    "feat",
				Scope:   "",
				Subject: "调整",
			},
			wantErr: nil,
		},
		//{
		//	Name:    "scope empty 2",
		//	Message: "feat(): global",
		//	wantErr: ErrScope,
		//},
		//{
		//	Name:    "type empty",
		//	Message: "(test): test",
		//	wantErr: ErrFormat,
		//},
		//{
		//	Name:    "subject empty 1",
		//	Message: "feat(test):",
		//	wantErr: ErrFormat,
		//},
		//{
		//	Name:    "subject empty 2",
		//	Message: "feat(test):   ",
		//	wantErr: ErrFormat,
		//},
		//{
		//	Name:    "subject empty 3",
		//	Message: "feat(test):        		 ",
		//	wantErr: ErrFormat,
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			f, err := NewFormat(tc.Message)
			if tc.wantErr == nil {
				assert.NoError(t, err)
			}
			if err != nil {
				assert.Equal(t, tc.wantErr.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.want, f)
		})
	}
}

func TestVerify(t *testing.T) {
	testCases := []struct {
		Name    string
		Message string
		want    Format
		wantErr error
	}{
		{
			Name:    "happy path",
			Message: "feat(test): samples",
			want: Format{
				Type:    "feat",
				Scope:   "test",
				Subject: "samples",
			},
			wantErr: nil,
		},
		{
			Name:    "happy path no scope",
			Message: "feat: samples",
			want: Format{
				Type:    "feat",
				Scope:   "",
				Subject: "samples",
			},
			wantErr: nil,
		},
		{
			Name:    "happy path no scope",
			Message: "feat: 中文",
			want: Format{
				Type:    "feat",
				Scope:   "",
				Subject: "中文",
			},
			wantErr: nil,
		},
		{
			Name:    "invalid type",
			Message: "invalid(test): test",
			wantErr: ErrType,
		},
		{
			Name:    "invalid style",
			Message: "feat(Hest): test",
			wantErr: ErrStyle,
		},
		{
			Name:    "invalid subject",
			Message: "feat(test): Add hoge",
			wantErr: ErrSubject,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			f, err := NewFormat(tc.Message)
			if err != nil {
				assert.NoError(t, err)
			}
			c, _ := NewConfig("")

			err = f.Verify(c)
			if tc.wantErr == nil {
				assert.NoError(t, err)
			}
			if err != nil {
				assert.Equal(t, tc.wantErr.Error(), err.Error())
				return
			}

			assert.Equal(t, tc.want, f)
		})
	}
}

func TestPattern(t *testing.T) {
	str := "World"                  // 示例字符串，注意其中包含了一个中文空格字符
	pattern := "^[a-z\\p{Han}]\\w*" // 中括号里面分别是匹配小写字母或者中文字符

	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if matched {
		fmt.Println("匹配成功")
	} else {
		fmt.Println("匹配失败")
	}
}
