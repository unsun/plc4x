/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataEventAlgorithmInhibit is the corresponding interface of BACnetConstructedDataEventAlgorithmInhibit
type BACnetConstructedDataEventAlgorithmInhibit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEventAlgorithmInhibit returns EventAlgorithmInhibit (property field)
	GetEventAlgorithmInhibit() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataEventAlgorithmInhibitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventAlgorithmInhibit.
// This is useful for switch cases.
type BACnetConstructedDataEventAlgorithmInhibitExactly interface {
	BACnetConstructedDataEventAlgorithmInhibit
	isBACnetConstructedDataEventAlgorithmInhibit() bool
}

// _BACnetConstructedDataEventAlgorithmInhibit is the data-structure of this message
type _BACnetConstructedDataEventAlgorithmInhibit struct {
	*_BACnetConstructedData
	EventAlgorithmInhibit BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventAlgorithmInhibit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_ALGORITHM_INHIBIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEventAlgorithmInhibit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibit) GetEventAlgorithmInhibit() BACnetApplicationTagBoolean {
	return m.EventAlgorithmInhibit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibit) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetEventAlgorithmInhibit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventAlgorithmInhibit factory function for _BACnetConstructedDataEventAlgorithmInhibit
func NewBACnetConstructedDataEventAlgorithmInhibit(eventAlgorithmInhibit BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventAlgorithmInhibit {
	_result := &_BACnetConstructedDataEventAlgorithmInhibit{
		EventAlgorithmInhibit:  eventAlgorithmInhibit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventAlgorithmInhibit(structType interface{}) BACnetConstructedDataEventAlgorithmInhibit {
	if casted, ok := structType.(BACnetConstructedDataEventAlgorithmInhibit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventAlgorithmInhibit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventAlgorithmInhibit) GetTypeName() string {
	return "BACnetConstructedDataEventAlgorithmInhibit"
}

func (m *_BACnetConstructedDataEventAlgorithmInhibit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEventAlgorithmInhibit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventAlgorithmInhibit)
	lengthInBits += m.EventAlgorithmInhibit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventAlgorithmInhibit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventAlgorithmInhibitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventAlgorithmInhibit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventAlgorithmInhibit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventAlgorithmInhibit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventAlgorithmInhibit)
	if pullErr := readBuffer.PullContext("eventAlgorithmInhibit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventAlgorithmInhibit")
	}
	_eventAlgorithmInhibit, _eventAlgorithmInhibitErr := BACnetApplicationTagParse(readBuffer)
	if _eventAlgorithmInhibitErr != nil {
		return nil, errors.Wrap(_eventAlgorithmInhibitErr, "Error parsing 'eventAlgorithmInhibit' field of BACnetConstructedDataEventAlgorithmInhibit")
	}
	eventAlgorithmInhibit := _eventAlgorithmInhibit.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("eventAlgorithmInhibit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventAlgorithmInhibit")
	}

	// Virtual field
	_actualValue := eventAlgorithmInhibit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventAlgorithmInhibit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventAlgorithmInhibit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEventAlgorithmInhibit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		EventAlgorithmInhibit: eventAlgorithmInhibit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEventAlgorithmInhibit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventAlgorithmInhibit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventAlgorithmInhibit")
		}

		// Simple Field (eventAlgorithmInhibit)
		if pushErr := writeBuffer.PushContext("eventAlgorithmInhibit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventAlgorithmInhibit")
		}
		_eventAlgorithmInhibitErr := writeBuffer.WriteSerializable(m.GetEventAlgorithmInhibit())
		if popErr := writeBuffer.PopContext("eventAlgorithmInhibit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventAlgorithmInhibit")
		}
		if _eventAlgorithmInhibitErr != nil {
			return errors.Wrap(_eventAlgorithmInhibitErr, "Error serializing 'eventAlgorithmInhibit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventAlgorithmInhibit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventAlgorithmInhibit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventAlgorithmInhibit) isBACnetConstructedDataEventAlgorithmInhibit() bool {
	return true
}

func (m *_BACnetConstructedDataEventAlgorithmInhibit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
