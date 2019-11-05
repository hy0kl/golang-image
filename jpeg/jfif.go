// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package jpeg

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Jfif struct {
	mVersion uint8
	sVersion uint8
	density  uint8
	xDensity uint16
	yDensity uint16
}

func NewJfif() *Jfif {
	return &Jfif{mVersion: 1, sVersion: 1,
		density: 1, xDensity: 1, yDensity: 1}
}

func (j *Jfif) Version() string {
	return fmt.Sprintf("%d.%d", j.mVersion, j.sVersion)
}

func (j *Jfif) SetDensity(d uint8) {
	j.density = d
}

func (j *Jfif) SetXDensity(d uint16) {
	j.xDensity = d
}

func (j *Jfif) SetYDensity(d uint16) {
	j.yDensity = d
}
func (j *Jfif) ToBytes() ([]byte, error) {
	res := []byte{0xff, 0xe0}
	buf := bytes.NewBufferString("JFIF")
	buf.WriteByte(0)
	buf.WriteByte(byte(j.mVersion))
	buf.WriteByte(byte(j.sVersion))
	buf.WriteByte(byte(j.density))
	var err error
	d := make([]byte, 2)
	binary.BigEndian.PutUint16(d, j.xDensity)
	_, err = buf.Write(d)
	if err != nil {
		return nil, err
	}
	binary.BigEndian.PutUint16(d, j.xDensity)
	_, err = buf.Write(d)
	if err != nil {
		return nil, err
	}
	var xThumnail uint8 = 0
	var yThumnail uint8 = 0
	// write xThumnail
	err = buf.WriteByte(byte(xThumnail))
	if err != nil {
		return nil, err
	}
	// write yThumnail
	err = buf.WriteByte(byte(yThumnail))
	if err != nil {
		return nil, err
	}
	//add jfif metadata length
	binary.BigEndian.PutUint16(d, uint16(buf.Len()+2))
	res = append(res, d...)
	res = append(res, buf.Bytes()...)
	return res, err
}
