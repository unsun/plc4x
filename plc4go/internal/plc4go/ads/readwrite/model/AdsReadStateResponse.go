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
type AdsReadStateResponse struct {
	Result      ReturnCode
	AdsState    uint16
	DeviceState uint16
	Parent      *AdsData
}

// The corresponding interface
type IAdsReadStateResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsReadStateResponse) CommandId() CommandId {
	return CommandId_ADS_READ_STATE
}

func (m *AdsReadStateResponse) Response() bool {
	return true
}

func (m *AdsReadStateResponse) InitializeParent(parent *AdsData) {
}

func NewAdsReadStateResponse(result ReturnCode, adsState uint16, deviceState uint16) *AdsData {
	child := &AdsReadStateResponse{
		Result:      result,
		AdsState:    adsState,
		DeviceState: deviceState,
		Parent:      NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsReadStateResponse(structType interface{}) *AdsReadStateResponse {
	castFunc := func(typ interface{}) *AdsReadStateResponse {
		if casted, ok := typ.(AdsReadStateResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsReadStateResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsReadStateResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsReadStateResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsReadStateResponse) GetTypeName() string {
	return "AdsReadStateResponse"
}

func (m *AdsReadStateResponse) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (result)
	lengthInBits += 32

	// Simple field (adsState)
	lengthInBits += 16

	// Simple field (deviceState)
	lengthInBits += 16

	return lengthInBits
}

func (m *AdsReadStateResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsReadStateResponseParse(io *utils.ReadBuffer) (*AdsData, error) {

	// Simple Field (result)
	result, _resultErr := ReturnCodeParse(io)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}

	// Simple Field (adsState)
	adsState, _adsStateErr := io.ReadUint16(16)
	if _adsStateErr != nil {
		return nil, errors.Wrap(_adsStateErr, "Error parsing 'adsState' field")
	}

	// Simple Field (deviceState)
	deviceState, _deviceStateErr := io.ReadUint16(16)
	if _deviceStateErr != nil {
		return nil, errors.Wrap(_deviceStateErr, "Error parsing 'deviceState' field")
	}

	// Create a partially initialized instance
	_child := &AdsReadStateResponse{
		Result:      result,
		AdsState:    adsState,
		DeviceState: deviceState,
		Parent:      &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsReadStateResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (result)
		_resultErr := m.Result.Serialize(io)
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (adsState)
		adsState := uint16(m.AdsState)
		_adsStateErr := io.WriteUint16(16, (adsState))
		if _adsStateErr != nil {
			return errors.Wrap(_adsStateErr, "Error serializing 'adsState' field")
		}

		// Simple Field (deviceState)
		deviceState := uint16(m.DeviceState)
		_deviceStateErr := io.WriteUint16(16, (deviceState))
		if _deviceStateErr != nil {
			return errors.Wrap(_deviceStateErr, "Error serializing 'deviceState' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsReadStateResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "result":
				var data ReturnCode
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Result = data
			case "adsState":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.AdsState = data
			case "deviceState":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DeviceState = data
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

func (m *AdsReadStateResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Result, xml.StartElement{Name: xml.Name{Local: "result"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.AdsState, xml.StartElement{Name: xml.Name{Local: "adsState"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DeviceState, xml.StartElement{Name: xml.Name{Local: "deviceState"}}); err != nil {
		return err
	}
	return nil
}

func (m AdsReadStateResponse) String() string {
	return string(m.Box("AdsReadStateResponse", utils.DefaultWidth*2))
}

func (m AdsReadStateResponse) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "AdsReadStateResponse"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Result", m.Result, width-2))
	boxes = append(boxes, utils.BoxAnything("AdsState", m.AdsState, width-2))
	boxes = append(boxes, utils.BoxAnything("DeviceState", m.DeviceState, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
