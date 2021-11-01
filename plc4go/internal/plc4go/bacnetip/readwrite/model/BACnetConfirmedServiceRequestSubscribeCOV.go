/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const BACnetConfirmedServiceRequestSubscribeCOV_SUBSCRIBERPROCESSIDENTIFIERHEADER uint8 = 0x09
const BACnetConfirmedServiceRequestSubscribeCOV_MONITOREDOBJECTIDENTIFIERHEADER uint8 = 0x1C
const BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSHEADER uint8 = 0x29
const BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSSKIPBITS uint8 = 0x00
const BACnetConfirmedServiceRequestSubscribeCOV_LIFETIMEHEADER uint8 = 0x07

// The data-structure of this message
type BACnetConfirmedServiceRequestSubscribeCOV struct {
	SubscriberProcessIdentifier   uint8
	MonitoredObjectType           uint16
	MonitoredObjectInstanceNumber uint32
	IssueConfirmedNotifications   bool
	LifetimeLength                uint8
	LifetimeSeconds               []int8
	Parent                        *BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestSubscribeCOV interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestSubscribeCOV) ServiceChoice() uint8 {
	return 0x05
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func NewBACnetConfirmedServiceRequestSubscribeCOV(subscriberProcessIdentifier uint8, monitoredObjectType uint16, monitoredObjectInstanceNumber uint32, issueConfirmedNotifications bool, lifetimeLength uint8, lifetimeSeconds []int8) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestSubscribeCOV{
		SubscriberProcessIdentifier:   subscriberProcessIdentifier,
		MonitoredObjectType:           monitoredObjectType,
		MonitoredObjectInstanceNumber: monitoredObjectInstanceNumber,
		IssueConfirmedNotifications:   issueConfirmedNotifications,
		LifetimeLength:                lifetimeLength,
		LifetimeSeconds:               lifetimeSeconds,
		Parent:                        NewBACnetConfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetConfirmedServiceRequestSubscribeCOV(structType interface{}) *BACnetConfirmedServiceRequestSubscribeCOV {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestSubscribeCOV {
		if casted, ok := typ.(BACnetConfirmedServiceRequestSubscribeCOV); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestSubscribeCOV); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestSubscribeCOV(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestSubscribeCOV(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOV"
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Const Field (subscriberProcessIdentifierHeader)
	lengthInBits += 8

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += 8

	// Const Field (monitoredObjectIdentifierHeader)
	lengthInBits += 8

	// Simple field (monitoredObjectType)
	lengthInBits += 10

	// Simple field (monitoredObjectInstanceNumber)
	lengthInBits += 22

	// Const Field (issueConfirmedNotificationsHeader)
	lengthInBits += 8

	// Const Field (issueConfirmedNotificationsSkipBits)
	lengthInBits += 7

	// Simple field (issueConfirmedNotifications)
	lengthInBits += 1

	// Const Field (lifetimeHeader)
	lengthInBits += 5

	// Simple field (lifetimeLength)
	lengthInBits += 3

	// Array field
	if len(m.LifetimeSeconds) > 0 {
		lengthInBits += 8 * uint16(len(m.LifetimeSeconds))
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestSubscribeCOVParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOV"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (subscriberProcessIdentifierHeader)
	subscriberProcessIdentifierHeader, _subscriberProcessIdentifierHeaderErr := readBuffer.ReadUint8("subscriberProcessIdentifierHeader", 8)
	if _subscriberProcessIdentifierHeaderErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierHeaderErr, "Error parsing 'subscriberProcessIdentifierHeader' field")
	}
	if subscriberProcessIdentifierHeader != BACnetConfirmedServiceRequestSubscribeCOV_SUBSCRIBERPROCESSIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestSubscribeCOV_SUBSCRIBERPROCESSIDENTIFIERHEADER) + " but got " + fmt.Sprintf("%d", subscriberProcessIdentifierHeader))
	}

	// Simple Field (subscriberProcessIdentifier)
	subscriberProcessIdentifier, _subscriberProcessIdentifierErr := readBuffer.ReadUint8("subscriberProcessIdentifier", 8)
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field")
	}

	// Const Field (monitoredObjectIdentifierHeader)
	monitoredObjectIdentifierHeader, _monitoredObjectIdentifierHeaderErr := readBuffer.ReadUint8("monitoredObjectIdentifierHeader", 8)
	if _monitoredObjectIdentifierHeaderErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierHeaderErr, "Error parsing 'monitoredObjectIdentifierHeader' field")
	}
	if monitoredObjectIdentifierHeader != BACnetConfirmedServiceRequestSubscribeCOV_MONITOREDOBJECTIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestSubscribeCOV_MONITOREDOBJECTIDENTIFIERHEADER) + " but got " + fmt.Sprintf("%d", monitoredObjectIdentifierHeader))
	}

	// Simple Field (monitoredObjectType)
	monitoredObjectType, _monitoredObjectTypeErr := readBuffer.ReadUint16("monitoredObjectType", 10)
	if _monitoredObjectTypeErr != nil {
		return nil, errors.Wrap(_monitoredObjectTypeErr, "Error parsing 'monitoredObjectType' field")
	}

	// Simple Field (monitoredObjectInstanceNumber)
	monitoredObjectInstanceNumber, _monitoredObjectInstanceNumberErr := readBuffer.ReadUint32("monitoredObjectInstanceNumber", 22)
	if _monitoredObjectInstanceNumberErr != nil {
		return nil, errors.Wrap(_monitoredObjectInstanceNumberErr, "Error parsing 'monitoredObjectInstanceNumber' field")
	}

	// Const Field (issueConfirmedNotificationsHeader)
	issueConfirmedNotificationsHeader, _issueConfirmedNotificationsHeaderErr := readBuffer.ReadUint8("issueConfirmedNotificationsHeader", 8)
	if _issueConfirmedNotificationsHeaderErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsHeaderErr, "Error parsing 'issueConfirmedNotificationsHeader' field")
	}
	if issueConfirmedNotificationsHeader != BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSHEADER) + " but got " + fmt.Sprintf("%d", issueConfirmedNotificationsHeader))
	}

	// Const Field (issueConfirmedNotificationsSkipBits)
	issueConfirmedNotificationsSkipBits, _issueConfirmedNotificationsSkipBitsErr := readBuffer.ReadUint8("issueConfirmedNotificationsSkipBits", 7)
	if _issueConfirmedNotificationsSkipBitsErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsSkipBitsErr, "Error parsing 'issueConfirmedNotificationsSkipBits' field")
	}
	if issueConfirmedNotificationsSkipBits != BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSSKIPBITS {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSSKIPBITS) + " but got " + fmt.Sprintf("%d", issueConfirmedNotificationsSkipBits))
	}

	// Simple Field (issueConfirmedNotifications)
	issueConfirmedNotifications, _issueConfirmedNotificationsErr := readBuffer.ReadBit("issueConfirmedNotifications")
	if _issueConfirmedNotificationsErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsErr, "Error parsing 'issueConfirmedNotifications' field")
	}

	// Const Field (lifetimeHeader)
	lifetimeHeader, _lifetimeHeaderErr := readBuffer.ReadUint8("lifetimeHeader", 5)
	if _lifetimeHeaderErr != nil {
		return nil, errors.Wrap(_lifetimeHeaderErr, "Error parsing 'lifetimeHeader' field")
	}
	if lifetimeHeader != BACnetConfirmedServiceRequestSubscribeCOV_LIFETIMEHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestSubscribeCOV_LIFETIMEHEADER) + " but got " + fmt.Sprintf("%d", lifetimeHeader))
	}

	// Simple Field (lifetimeLength)
	lifetimeLength, _lifetimeLengthErr := readBuffer.ReadUint8("lifetimeLength", 3)
	if _lifetimeLengthErr != nil {
		return nil, errors.Wrap(_lifetimeLengthErr, "Error parsing 'lifetimeLength' field")
	}

	// Array field (lifetimeSeconds)
	if pullErr := readBuffer.PullContext("lifetimeSeconds", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	lifetimeSeconds := make([]int8, lifetimeLength)
	for curItem := uint16(0); curItem < uint16(lifetimeLength); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'lifetimeSeconds' field")
		}
		lifetimeSeconds[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("lifetimeSeconds", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOV"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestSubscribeCOV{
		SubscriberProcessIdentifier:   subscriberProcessIdentifier,
		MonitoredObjectType:           monitoredObjectType,
		MonitoredObjectInstanceNumber: monitoredObjectInstanceNumber,
		IssueConfirmedNotifications:   issueConfirmedNotifications,
		LifetimeLength:                lifetimeLength,
		LifetimeSeconds:               lifetimeSeconds,
		Parent:                        &BACnetConfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOV"); pushErr != nil {
			return pushErr
		}

		// Const Field (subscriberProcessIdentifierHeader)
		_subscriberProcessIdentifierHeaderErr := writeBuffer.WriteUint8("subscriberProcessIdentifierHeader", 8, 0x09)
		if _subscriberProcessIdentifierHeaderErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierHeaderErr, "Error serializing 'subscriberProcessIdentifierHeader' field")
		}

		// Simple Field (subscriberProcessIdentifier)
		subscriberProcessIdentifier := uint8(m.SubscriberProcessIdentifier)
		_subscriberProcessIdentifierErr := writeBuffer.WriteUint8("subscriberProcessIdentifier", 8, (subscriberProcessIdentifier))
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Const Field (monitoredObjectIdentifierHeader)
		_monitoredObjectIdentifierHeaderErr := writeBuffer.WriteUint8("monitoredObjectIdentifierHeader", 8, 0x1C)
		if _monitoredObjectIdentifierHeaderErr != nil {
			return errors.Wrap(_monitoredObjectIdentifierHeaderErr, "Error serializing 'monitoredObjectIdentifierHeader' field")
		}

		// Simple Field (monitoredObjectType)
		monitoredObjectType := uint16(m.MonitoredObjectType)
		_monitoredObjectTypeErr := writeBuffer.WriteUint16("monitoredObjectType", 10, (monitoredObjectType))
		if _monitoredObjectTypeErr != nil {
			return errors.Wrap(_monitoredObjectTypeErr, "Error serializing 'monitoredObjectType' field")
		}

		// Simple Field (monitoredObjectInstanceNumber)
		monitoredObjectInstanceNumber := uint32(m.MonitoredObjectInstanceNumber)
		_monitoredObjectInstanceNumberErr := writeBuffer.WriteUint32("monitoredObjectInstanceNumber", 22, (monitoredObjectInstanceNumber))
		if _monitoredObjectInstanceNumberErr != nil {
			return errors.Wrap(_monitoredObjectInstanceNumberErr, "Error serializing 'monitoredObjectInstanceNumber' field")
		}

		// Const Field (issueConfirmedNotificationsHeader)
		_issueConfirmedNotificationsHeaderErr := writeBuffer.WriteUint8("issueConfirmedNotificationsHeader", 8, 0x29)
		if _issueConfirmedNotificationsHeaderErr != nil {
			return errors.Wrap(_issueConfirmedNotificationsHeaderErr, "Error serializing 'issueConfirmedNotificationsHeader' field")
		}

		// Const Field (issueConfirmedNotificationsSkipBits)
		_issueConfirmedNotificationsSkipBitsErr := writeBuffer.WriteUint8("issueConfirmedNotificationsSkipBits", 7, 0x00)
		if _issueConfirmedNotificationsSkipBitsErr != nil {
			return errors.Wrap(_issueConfirmedNotificationsSkipBitsErr, "Error serializing 'issueConfirmedNotificationsSkipBits' field")
		}

		// Simple Field (issueConfirmedNotifications)
		issueConfirmedNotifications := bool(m.IssueConfirmedNotifications)
		_issueConfirmedNotificationsErr := writeBuffer.WriteBit("issueConfirmedNotifications", (issueConfirmedNotifications))
		if _issueConfirmedNotificationsErr != nil {
			return errors.Wrap(_issueConfirmedNotificationsErr, "Error serializing 'issueConfirmedNotifications' field")
		}

		// Const Field (lifetimeHeader)
		_lifetimeHeaderErr := writeBuffer.WriteUint8("lifetimeHeader", 5, 0x07)
		if _lifetimeHeaderErr != nil {
			return errors.Wrap(_lifetimeHeaderErr, "Error serializing 'lifetimeHeader' field")
		}

		// Simple Field (lifetimeLength)
		lifetimeLength := uint8(m.LifetimeLength)
		_lifetimeLengthErr := writeBuffer.WriteUint8("lifetimeLength", 3, (lifetimeLength))
		if _lifetimeLengthErr != nil {
			return errors.Wrap(_lifetimeLengthErr, "Error serializing 'lifetimeLength' field")
		}

		// Array Field (lifetimeSeconds)
		if m.LifetimeSeconds != nil {
			if pushErr := writeBuffer.PushContext("lifetimeSeconds", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.LifetimeSeconds {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'lifetimeSeconds' field")
				}
			}
			if popErr := writeBuffer.PopContext("lifetimeSeconds", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOV"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
