// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_GetRefs(t *testing.T) {
	bareRepo1Path := filepath.Join(testReposDir, "repo1_bare")
	bareRepo1, err := OpenRepository(bareRepo1Path)
	assert.NoError(t, err)

	refs, err := bareRepo1.GetRefs()

	assert.NoError(t, err)
	assert.Len(t, refs, 4)

	expectedRefs := []string{
		BranchPrefix + "branch1",
		BranchPrefix + "branch2",
		BranchPrefix + "master",
		TagPrefix + "test",
	}

	for _, ref := range refs {
		assert.Contains(t, expectedRefs, ref.Name)
	}
}
