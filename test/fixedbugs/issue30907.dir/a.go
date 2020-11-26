// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package a

type UUID string

func New() UUID {
	return Must(NewRandom())
}

func NewRandom() (UUID, error) {
	return "", nil
}

func Must(uuid UUID, err error) UUID {
	return uuid
}
