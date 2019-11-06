// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package jpeg

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"image"
)

const thumbnail_max_px = 21845

type Jfif struct {
	mVersion   uint8
	sVersion   uint8
	density    uint8
	xDensity   uint16
	yDensity   uint16
	xthumbnail uint8
	ythumbnail uint8
	thumbnail  []byte
}

func NewJfif() *Jfif {
	return &Jfif{mVersion: 1, sVersion: 1,
		density: 1, xDensity: 1, yDensity: 1, thumbnail: nil}
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
	// write xThumnail
	err = buf.WriteByte(byte(j.xthumbnail))
	if err != nil {
		return nil, err
	}
	// write yThumnail
	err = buf.WriteByte(byte(j.ythumbnail))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(j.thumbnail)
	if err != nil {
		return nil, err
	}
	//add jfif metadata length
	binary.BigEndian.PutUint16(d, uint16(buf.Len()+2))
	res = append(res, d...)
	res = append(res, buf.Bytes()...)
	return res, err
}
func (jfif *Jfif) SetThumbnail(img image.Image) error {
	b := img.Bounds()
	if b.Dx()*b.Dy() > thumbnail_max_px {
		return errors.New("thumbnail should not lager than 21840px")
	}
	jfif.xthumbnail = uint8(b.Dx())
	jfif.ythumbnail = uint8(b.Dy())
	fmt.Println(jfif.xthumbnail)
	fmt.Println(jfif.ythumbnail)
	jfif.thumbnail = make([]byte, 3*b.Dx()*b.Dy())
	index := 0
	for j := 0; j < b.Dy(); j++ {
		for i := 0; i < b.Dx(); i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			jfif.thumbnail[index] = byte(uint8(r))
			index++
			jfif.thumbnail[index] = byte(uint8(g))
			index++
			jfif.thumbnail[index] = byte(uint8(b))
			index++
		}
	}
	return nil
}
