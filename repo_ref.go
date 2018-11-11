// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// GetRefs returns all references of the repository.
func (repo *Repository) GetRefs() ([]*Reference, error) {
	r, err := git.PlainOpen(repo.Path)
	if err != nil {
		return nil, err
	}

	refsIter, err := r.References()
	if err != nil {
		return nil, err
	}
	refs := make([]*Reference, 0)
	err = refsIter.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name() != plumbing.HEAD && !ref.Name().IsRemote() {
			r := &Reference{
				Name:   ref.Name().String(),
				Object: SHA1(ref.Hash()),
				Type:   "commit",
			}
			if ref.Name().IsTag() {
				r.Type = "tag"
			}
			refs = append(refs, r)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return refs, nil
}
