// Copyright 2012-2016, 2019 Charles Banning. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

package mxj

// ---------------- expose Map methods to MapSeq type ---------------------------

// Pretty print a Map.
func (mv MapSeq) StringIndent(offset ...int) string {
	return writeMap(map[string]interface{}(mv), true, true, offset...)
}

// Pretty print a Map without the value type information - just key:value entries.
func (mv MapSeq) StringIndentNoTypeInfo(offset ...int) string {
	return writeMap(map[string]interface{}(mv), false, true, offset...)
}
