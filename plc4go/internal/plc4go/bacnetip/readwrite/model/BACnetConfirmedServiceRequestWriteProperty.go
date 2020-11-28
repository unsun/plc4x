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
    "encoding/base64"
    "encoding/xml"
    "errors"
    "io"
    "github.com/apache/plc4x/plc4go/internal/plc4go/utils"
    "strconv"
)

// Constant values.
const BACnetConfirmedServiceRequestWriteProperty_OBJECTIDENTIFIERHEADER uint8 = 0x0C
const BACnetConfirmedServiceRequestWriteProperty_PROPERTYIDENTIFIERHEADER uint8 = 0x03
const BACnetConfirmedServiceRequestWriteProperty_OPENINGTAG uint8 = 0x3E
const BACnetConfirmedServiceRequestWriteProperty_CLOSINGTAG uint8 = 0x3F

// The data-structure of this message
type BACnetConfirmedServiceRequestWriteProperty struct {
    ObjectType uint16
    ObjectInstanceNumber uint32
    PropertyIdentifierLength uint8
    PropertyIdentifier []int8
    Value *BACnetTag
    Priority *BACnetTag
    Parent *BACnetConfirmedServiceRequest
    IBACnetConfirmedServiceRequestWriteProperty
}

// The corresponding interface
type IBACnetConfirmedServiceRequestWriteProperty interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestWriteProperty) ServiceChoice() uint8 {
    return 0x0F
}


func (m *BACnetConfirmedServiceRequestWriteProperty) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func NewBACnetConfirmedServiceRequestWriteProperty(objectType uint16, objectInstanceNumber uint32, propertyIdentifierLength uint8, propertyIdentifier []int8, value *BACnetTag, priority *BACnetTag, ) *BACnetConfirmedServiceRequest {
    child := &BACnetConfirmedServiceRequestWriteProperty{
        ObjectType: objectType,
        ObjectInstanceNumber: objectInstanceNumber,
        PropertyIdentifierLength: propertyIdentifierLength,
        PropertyIdentifier: propertyIdentifier,
        Value: value,
        Priority: priority,
        Parent: NewBACnetConfirmedServiceRequest(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastBACnetConfirmedServiceRequestWriteProperty(structType interface{}) *BACnetConfirmedServiceRequestWriteProperty {
    castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestWriteProperty {
        if casted, ok := typ.(BACnetConfirmedServiceRequestWriteProperty); ok {
            return &casted
        }
        if casted, ok := typ.(*BACnetConfirmedServiceRequestWriteProperty); ok {
            return casted
        }
        if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
            return CastBACnetConfirmedServiceRequestWriteProperty(casted.Child)
        }
        if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
            return CastBACnetConfirmedServiceRequestWriteProperty(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestWriteProperty) GetTypeName() string {
    return "BACnetConfirmedServiceRequestWriteProperty"
}

func (m *BACnetConfirmedServiceRequestWriteProperty) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Const Field (objectIdentifierHeader)
    lengthInBits += 8

    // Simple field (objectType)
    lengthInBits += 10

    // Simple field (objectInstanceNumber)
    lengthInBits += 22

    // Const Field (propertyIdentifierHeader)
    lengthInBits += 5

    // Simple field (propertyIdentifierLength)
    lengthInBits += 3

    // Array field
    if len(m.PropertyIdentifier) > 0 {
        lengthInBits += 8 * uint16(len(m.PropertyIdentifier))
    }

    // Const Field (openingTag)
    lengthInBits += 8

    // Simple field (value)
    lengthInBits += m.Value.LengthInBits()

    // Const Field (closingTag)
    lengthInBits += 8

    // Optional Field (priority)
    if m.Priority != nil {
        lengthInBits += (*m.Priority).LengthInBits()
    }

    return lengthInBits
}

func (m *BACnetConfirmedServiceRequestWriteProperty) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestWritePropertyParse(io *utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
    var startPos = io.GetPos()
    var curPos uint16

    // Const Field (objectIdentifierHeader)
    objectIdentifierHeader, _objectIdentifierHeaderErr := io.ReadUint8(8)
    if _objectIdentifierHeaderErr != nil {
        return nil, errors.New("Error parsing 'objectIdentifierHeader' field " + _objectIdentifierHeaderErr.Error())
    }
    if objectIdentifierHeader != BACnetConfirmedServiceRequestWriteProperty_OBJECTIDENTIFIERHEADER {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestWriteProperty_OBJECTIDENTIFIERHEADER)) + " but got " + strconv.Itoa(int(objectIdentifierHeader)))
    }

    // Simple Field (objectType)
    objectType, _objectTypeErr := io.ReadUint16(10)
    if _objectTypeErr != nil {
        return nil, errors.New("Error parsing 'objectType' field " + _objectTypeErr.Error())
    }

    // Simple Field (objectInstanceNumber)
    objectInstanceNumber, _objectInstanceNumberErr := io.ReadUint32(22)
    if _objectInstanceNumberErr != nil {
        return nil, errors.New("Error parsing 'objectInstanceNumber' field " + _objectInstanceNumberErr.Error())
    }

    // Const Field (propertyIdentifierHeader)
    propertyIdentifierHeader, _propertyIdentifierHeaderErr := io.ReadUint8(5)
    if _propertyIdentifierHeaderErr != nil {
        return nil, errors.New("Error parsing 'propertyIdentifierHeader' field " + _propertyIdentifierHeaderErr.Error())
    }
    if propertyIdentifierHeader != BACnetConfirmedServiceRequestWriteProperty_PROPERTYIDENTIFIERHEADER {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestWriteProperty_PROPERTYIDENTIFIERHEADER)) + " but got " + strconv.Itoa(int(propertyIdentifierHeader)))
    }

    // Simple Field (propertyIdentifierLength)
    propertyIdentifierLength, _propertyIdentifierLengthErr := io.ReadUint8(3)
    if _propertyIdentifierLengthErr != nil {
        return nil, errors.New("Error parsing 'propertyIdentifierLength' field " + _propertyIdentifierLengthErr.Error())
    }

    // Array field (propertyIdentifier)
    // Count array
    propertyIdentifier := make([]int8, propertyIdentifierLength)
    for curItem := uint16(0); curItem < uint16(propertyIdentifierLength); curItem++ {
        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'propertyIdentifier' field " + _err.Error())
        }
        propertyIdentifier[curItem] = _item
    }

    // Const Field (openingTag)
    openingTag, _openingTagErr := io.ReadUint8(8)
    if _openingTagErr != nil {
        return nil, errors.New("Error parsing 'openingTag' field " + _openingTagErr.Error())
    }
    if openingTag != BACnetConfirmedServiceRequestWriteProperty_OPENINGTAG {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestWriteProperty_OPENINGTAG)) + " but got " + strconv.Itoa(int(openingTag)))
    }

    // Simple Field (value)
    value, _valueErr := BACnetTagParse(io)
    if _valueErr != nil {
        return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
    }

    // Const Field (closingTag)
    closingTag, _closingTagErr := io.ReadUint8(8)
    if _closingTagErr != nil {
        return nil, errors.New("Error parsing 'closingTag' field " + _closingTagErr.Error())
    }
    if closingTag != BACnetConfirmedServiceRequestWriteProperty_CLOSINGTAG {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestWriteProperty_CLOSINGTAG)) + " but got " + strconv.Itoa(int(closingTag)))
    }

    // Optional Field (priority) (Can be skipped, if a given expression evaluates to false)
    curPos = io.GetPos() - startPos
    var priority *BACnetTag = nil
    if bool((curPos) < (((len) - ((1))))) {
        _message, _err := BACnetTagParse(io)
        if _err != nil {
            return nil, errors.New("Error parsing 'priority' field " + _err.Error())
        }
        priority = _message
    }

    // Create a partially initialized instance
    _child := &BACnetConfirmedServiceRequestWriteProperty{
        ObjectType: objectType,
        ObjectInstanceNumber: objectInstanceNumber,
        PropertyIdentifierLength: propertyIdentifierLength,
        PropertyIdentifier: propertyIdentifier,
        Value: value,
        Priority: priority,
        Parent: &BACnetConfirmedServiceRequest{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *BACnetConfirmedServiceRequestWriteProperty) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Const Field (objectIdentifierHeader)
    _objectIdentifierHeaderErr := io.WriteUint8(8, 0x0C)
    if _objectIdentifierHeaderErr != nil {
        return errors.New("Error serializing 'objectIdentifierHeader' field " + _objectIdentifierHeaderErr.Error())
    }

    // Simple Field (objectType)
    objectType := uint16(m.ObjectType)
    _objectTypeErr := io.WriteUint16(10, (objectType))
    if _objectTypeErr != nil {
        return errors.New("Error serializing 'objectType' field " + _objectTypeErr.Error())
    }

    // Simple Field (objectInstanceNumber)
    objectInstanceNumber := uint32(m.ObjectInstanceNumber)
    _objectInstanceNumberErr := io.WriteUint32(22, (objectInstanceNumber))
    if _objectInstanceNumberErr != nil {
        return errors.New("Error serializing 'objectInstanceNumber' field " + _objectInstanceNumberErr.Error())
    }

    // Const Field (propertyIdentifierHeader)
    _propertyIdentifierHeaderErr := io.WriteUint8(5, 0x03)
    if _propertyIdentifierHeaderErr != nil {
        return errors.New("Error serializing 'propertyIdentifierHeader' field " + _propertyIdentifierHeaderErr.Error())
    }

    // Simple Field (propertyIdentifierLength)
    propertyIdentifierLength := uint8(m.PropertyIdentifierLength)
    _propertyIdentifierLengthErr := io.WriteUint8(3, (propertyIdentifierLength))
    if _propertyIdentifierLengthErr != nil {
        return errors.New("Error serializing 'propertyIdentifierLength' field " + _propertyIdentifierLengthErr.Error())
    }

    // Array Field (propertyIdentifier)
    if m.PropertyIdentifier != nil {
        for _, _element := range m.PropertyIdentifier {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'propertyIdentifier' field " + _elementErr.Error())
            }
        }
    }

    // Const Field (openingTag)
    _openingTagErr := io.WriteUint8(8, 0x3E)
    if _openingTagErr != nil {
        return errors.New("Error serializing 'openingTag' field " + _openingTagErr.Error())
    }

    // Simple Field (value)
    _valueErr := m.Value.Serialize(io)
    if _valueErr != nil {
        return errors.New("Error serializing 'value' field " + _valueErr.Error())
    }

    // Const Field (closingTag)
    _closingTagErr := io.WriteUint8(8, 0x3F)
    if _closingTagErr != nil {
        return errors.New("Error serializing 'closingTag' field " + _closingTagErr.Error())
    }

    // Optional Field (priority) (Can be skipped, if the value is null)
    var priority *BACnetTag = nil
    if m.Priority != nil {
        priority = m.Priority
        _priorityErr := priority.Serialize(io)
        if _priorityErr != nil {
            return errors.New("Error serializing 'priority' field " + _priorityErr.Error())
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetConfirmedServiceRequestWriteProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "objectType":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ObjectType = data
            case "objectInstanceNumber":
                var data uint32
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ObjectInstanceNumber = data
            case "propertyIdentifierLength":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.PropertyIdentifierLength = data
            case "propertyIdentifier":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.PropertyIdentifier = utils.ByteArrayToInt8Array(_decoded[0:_len])
            case "value":
                var dt *BACnetTag
                if err := d.DecodeElement(&dt, &tok); err != nil {
                    return err
                }
                m.Value = dt
            case "priority":
                var dt *BACnetTag
                if err := d.DecodeElement(&dt, &tok); err != nil {
                    return err
                }
                m.Priority = dt
            }
        }
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
}

func (m *BACnetConfirmedServiceRequestWriteProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.ObjectType, xml.StartElement{Name: xml.Name{Local: "objectType"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ObjectInstanceNumber, xml.StartElement{Name: xml.Name{Local: "objectInstanceNumber"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.PropertyIdentifierLength, xml.StartElement{Name: xml.Name{Local: "propertyIdentifierLength"}}); err != nil {
        return err
    }
    _encodedPropertyIdentifier := make([]byte, base64.StdEncoding.EncodedLen(len(m.PropertyIdentifier)))
    base64.StdEncoding.Encode(_encodedPropertyIdentifier, utils.Int8ArrayToByteArray(m.PropertyIdentifier))
    if err := e.EncodeElement(_encodedPropertyIdentifier, xml.StartElement{Name: xml.Name{Local: "propertyIdentifier"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Value, xml.StartElement{Name: xml.Name{Local: "value"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Priority, xml.StartElement{Name: xml.Name{Local: "priority"}}); err != nil {
        return err
    }
    return nil
}

