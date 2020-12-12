// @Desc: test internal import
// @CreateTime: 2020/12/12
package test

import (
    "gohomework.com/Week04/studyDemo/internal"
    "testing"
)

func TestIsZero(t *testing.T) {
    res := internal.IsZero(0)
    println(res)
}

