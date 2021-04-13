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
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduDataGroupValueResponse struct {
	DataFirstByte int8
	Data          []int8
	Parent        *ApduData
}

// The corresponding interface
type IApduDataGroupValueResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataGroupValueResponse) ApciType() uint8 {
	return 0x1
}

func (m *ApduDataGroupValueResponse) InitializeParent(parent *ApduData) {
}

func NewApduDataGroupValueResponse(dataFirstByte int8, data []int8) *ApduData {
	child := &ApduDataGroupValueResponse{
		DataFirstByte: dataFirstByte,
		Data:          data,
		Parent:        NewApduData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataGroupValueResponse(structType interface{}) *ApduDataGroupValueResponse {
	castFunc := func(typ interface{}) *ApduDataGroupValueResponse {
		if casted, ok := typ.(ApduDataGroupValueResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataGroupValueResponse); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataGroupValueResponse(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataGroupValueResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataGroupValueResponse) GetTypeName() string {
	return "ApduDataGroupValueResponse"
}

func (m *ApduDataGroupValueResponse) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (dataFirstByte)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataGroupValueResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataGroupValueResponseParse(io *utils.ReadBuffer, dataLength uint8) (*ApduData, error) {

	// Simple Field (dataFirstByte)
	dataFirstByte, _dataFirstByteErr := io.ReadInt8(6)
	if _dataFirstByteErr != nil {
		return nil, errors.Wrap(_dataFirstByteErr, "Error parsing 'dataFirstByte' field")
	}

	// Array field (data)
	// Count array
	data := make([]int8, utils.InlineIf(bool(bool((dataLength) < (1))), func() uint16 { return uint16(uint16(0)) }, func() uint16 { return uint16(uint16(dataLength) - uint16(uint16(1))) }))
	for curItem := uint16(0); curItem < uint16(utils.InlineIf(bool(bool((dataLength) < (1))), func() uint16 { return uint16(uint16(0)) }, func() uint16 { return uint16(uint16(dataLength) - uint16(uint16(1))) })); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &ApduDataGroupValueResponse{
		DataFirstByte: dataFirstByte,
		Data:          data,
		Parent:        &ApduData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataGroupValueResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (dataFirstByte)
		dataFirstByte := int8(m.DataFirstByte)
		_dataFirstByteErr := io.WriteInt8(6, (dataFirstByte))
		if _dataFirstByteErr != nil {
			return errors.Wrap(_dataFirstByteErr, "Error serializing 'dataFirstByte' field")
		}

		// Array Field (data)
		if m.Data != nil {
			for _, _element := range m.Data {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataGroupValueResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "dataFirstByte":
				var data int8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DataFirstByte = data
			case "data":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *ApduDataGroupValueResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.DataFirstByte, xml.StartElement{Name: xml.Name{Local: "dataFirstByte"}}); err != nil {
		return err
	}
	_encodedData := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Data))
	_encodedData = strings.ToUpper(_encodedData)
	if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m ApduDataGroupValueResponse) String() string {
	return string(m.Box("ApduDataGroupValueResponse", utils.DefaultWidth*2))
}

func (m ApduDataGroupValueResponse) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ApduDataGroupValueResponse"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("DataFirstByte", m.DataFirstByte, width-2))
	boxes = append(boxes, utils.BoxAnything("Data", m.Data, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
