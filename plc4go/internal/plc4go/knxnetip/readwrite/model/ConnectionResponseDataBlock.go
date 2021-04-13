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
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ConnectionResponseDataBlock struct {
	Child IConnectionResponseDataBlockChild
}

// The corresponding interface
type IConnectionResponseDataBlock interface {
	ConnectionType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IConnectionResponseDataBlockParent interface {
	SerializeParent(io utils.WriteBuffer, child IConnectionResponseDataBlock, serializeChildFunction func() error) error
	GetTypeName() string
}

type IConnectionResponseDataBlockChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *ConnectionResponseDataBlock)
	GetTypeName() string
	IConnectionResponseDataBlock
}

func NewConnectionResponseDataBlock() *ConnectionResponseDataBlock {
	return &ConnectionResponseDataBlock{}
}

func CastConnectionResponseDataBlock(structType interface{}) *ConnectionResponseDataBlock {
	castFunc := func(typ interface{}) *ConnectionResponseDataBlock {
		if casted, ok := typ.(ConnectionResponseDataBlock); ok {
			return &casted
		}
		if casted, ok := typ.(*ConnectionResponseDataBlock); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ConnectionResponseDataBlock) GetTypeName() string {
	return "ConnectionResponseDataBlock"
}

func (m *ConnectionResponseDataBlock) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8
	// Discriminator Field (connectionType)
	lengthInBits += 8

	// Length of sub-type elements will be added by sub-type...
	lengthInBits += m.Child.LengthInBits()

	return lengthInBits
}

func (m *ConnectionResponseDataBlock) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ConnectionResponseDataBlockParse(io *utils.ReadBuffer) (*ConnectionResponseDataBlock, error) {

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := io.ReadUint8(8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Discriminator Field (connectionType) (Used as input to a switch field)
	connectionType, _connectionTypeErr := io.ReadUint8(8)
	if _connectionTypeErr != nil {
		return nil, errors.Wrap(_connectionTypeErr, "Error parsing 'connectionType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ConnectionResponseDataBlock
	var typeSwitchError error
	switch {
	case connectionType == 0x03: // ConnectionResponseDataBlockDeviceManagement
		_parent, typeSwitchError = ConnectionResponseDataBlockDeviceManagementParse(io)
	case connectionType == 0x04: // ConnectionResponseDataBlockTunnelConnection
		_parent, typeSwitchError = ConnectionResponseDataBlockTunnelConnectionParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ConnectionResponseDataBlock) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *ConnectionResponseDataBlock) SerializeParent(io utils.WriteBuffer, child IConnectionResponseDataBlock, serializeChildFunction func() error) error {

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.LengthInBytes()))
	_structureLengthErr := io.WriteUint8(8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Discriminator Field (connectionType) (Used as input to a switch field)
	connectionType := uint8(child.ConnectionType())
	_connectionTypeErr := io.WriteUint8(8, (connectionType))

	if _connectionTypeErr != nil {
		return errors.Wrap(_connectionTypeErr, "Error serializing 'connectionType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	return nil
}

func (m *ConnectionResponseDataBlock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of ConnectionResponseDataBlock")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionResponseDataBlockDeviceManagement":
					var dt *ConnectionResponseDataBlockDeviceManagement
					if m.Child != nil {
						dt = m.Child.(*ConnectionResponseDataBlockDeviceManagement)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ConnectionResponseDataBlockTunnelConnection":
					var dt *ConnectionResponseDataBlockTunnelConnection
					if m.Child != nil {
						dt = m.Child.(*ConnectionResponseDataBlockTunnelConnection)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				}
			}
		}
	}
}

func (m *ConnectionResponseDataBlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.knxnetip.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m ConnectionResponseDataBlock) String() string {
	return string(m.Box("ConnectionResponseDataBlock", utils.DefaultWidth*2))
}

func (m ConnectionResponseDataBlock) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ConnectionResponseDataBlock"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("", m.Child, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
