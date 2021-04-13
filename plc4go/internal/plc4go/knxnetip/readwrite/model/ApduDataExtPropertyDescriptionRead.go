//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduDataExtPropertyDescriptionRead struct {
	ObjectIndex uint8
	PropertyId  uint8
	Index       uint8
	Parent      *ApduDataExt
}

// The corresponding interface
type IApduDataExtPropertyDescriptionRead interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtPropertyDescriptionRead) ExtApciType() uint8 {
	return 0x18
}

func (m *ApduDataExtPropertyDescriptionRead) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtPropertyDescriptionRead(objectIndex uint8, propertyId uint8, index uint8) *ApduDataExt {
	child := &ApduDataExtPropertyDescriptionRead{
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Index:       index,
		Parent:      NewApduDataExt(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataExtPropertyDescriptionRead(structType interface{}) *ApduDataExtPropertyDescriptionRead {
	castFunc := func(typ interface{}) *ApduDataExtPropertyDescriptionRead {
		if casted, ok := typ.(ApduDataExtPropertyDescriptionRead); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtPropertyDescriptionRead); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtPropertyDescriptionRead(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtPropertyDescriptionRead(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtPropertyDescriptionRead) GetTypeName() string {
	return "ApduDataExtPropertyDescriptionRead"
}

func (m *ApduDataExtPropertyDescriptionRead) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (index)
	lengthInBits += 8

	return lengthInBits
}

func (m *ApduDataExtPropertyDescriptionRead) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtPropertyDescriptionReadParse(io *utils.ReadBuffer) (*ApduDataExt, error) {

	// Simple Field (objectIndex)
	objectIndex, _objectIndexErr := io.ReadUint8(8)
	if _objectIndexErr != nil {
		return nil, errors.Wrap(_objectIndexErr, "Error parsing 'objectIndex' field")
	}

	// Simple Field (propertyId)
	propertyId, _propertyIdErr := io.ReadUint8(8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field")
	}

	// Simple Field (index)
	index, _indexErr := io.ReadUint8(8)
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtPropertyDescriptionRead{
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Index:       index,
		Parent:      &ApduDataExt{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataExtPropertyDescriptionRead) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (objectIndex)
		objectIndex := uint8(m.ObjectIndex)
		_objectIndexErr := io.WriteUint8(8, (objectIndex))
		if _objectIndexErr != nil {
			return errors.Wrap(_objectIndexErr, "Error serializing 'objectIndex' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.PropertyId)
		_propertyIdErr := io.WriteUint8(8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (index)
		index := uint8(m.Index)
		_indexErr := io.WriteUint8(8, (index))
		if _indexErr != nil {
			return errors.Wrap(_indexErr, "Error serializing 'index' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtPropertyDescriptionRead) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "objectIndex":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ObjectIndex = data
			case "propertyId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.PropertyId = data
			case "index":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Index = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *ApduDataExtPropertyDescriptionRead) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.ObjectIndex, xml.StartElement{Name: xml.Name{Local: "objectIndex"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.PropertyId, xml.StartElement{Name: xml.Name{Local: "propertyId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Index, xml.StartElement{Name: xml.Name{Local: "index"}}); err != nil {
		return err
	}
	return nil
}

func (m ApduDataExtPropertyDescriptionRead) String() string {
	return string(m.Box("ApduDataExtPropertyDescriptionRead", utils.DefaultWidth*2))
}

func (m ApduDataExtPropertyDescriptionRead) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ApduDataExtPropertyDescriptionRead"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("ObjectIndex", m.ObjectIndex, width-2))
	boxes = append(boxes, utils.BoxAnything("PropertyId", m.PropertyId, width-2))
	boxes = append(boxes, utils.BoxAnything("Index", m.Index, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
