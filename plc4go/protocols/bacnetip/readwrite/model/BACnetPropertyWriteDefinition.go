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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyWriteDefinition is the corresponding interface of BACnetPropertyWriteDefinition
type BACnetPropertyWriteDefinition interface {
	utils.LengthAware
	utils.Serializable
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
}

// BACnetPropertyWriteDefinitionExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyWriteDefinition.
// This is useful for switch cases.
type BACnetPropertyWriteDefinitionExactly interface {
	BACnetPropertyWriteDefinition
	isBACnetPropertyWriteDefinition() bool
}

// _BACnetPropertyWriteDefinition is the data-structure of this message
type _BACnetPropertyWriteDefinition struct {
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	PropertyValue      BACnetConstructedData
	Priority           BACnetContextTagUnsignedInteger

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyWriteDefinition) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetPropertyWriteDefinition) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetPropertyWriteDefinition) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

func (m *_BACnetPropertyWriteDefinition) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyWriteDefinition factory function for _BACnetPropertyWriteDefinition
func NewBACnetPropertyWriteDefinition(propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, propertyValue BACnetConstructedData, priority BACnetContextTagUnsignedInteger, objectTypeArgument BACnetObjectType) *_BACnetPropertyWriteDefinition {
	return &_BACnetPropertyWriteDefinition{PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, PropertyValue: propertyValue, Priority: priority, ObjectTypeArgument: objectTypeArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyWriteDefinition(structType interface{}) BACnetPropertyWriteDefinition {
	if casted, ok := structType.(BACnetPropertyWriteDefinition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyWriteDefinition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyWriteDefinition) GetTypeName() string {
	return "BACnetPropertyWriteDefinition"
}

func (m *_BACnetPropertyWriteDefinition) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyWriteDefinition) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits()
	}

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += m.PropertyValue.GetLengthInBits()
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetPropertyWriteDefinition) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyWriteDefinitionParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPropertyWriteDefinition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyWriteDefinition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyWriteDefinition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetPropertyWriteDefinition")
	}
	propertyIdentifier := _propertyIdentifier.(BACnetPropertyIdentifierTagged)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for arrayIndex")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field of BACnetPropertyWriteDefinition")
		default:
			arrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if a given expression evaluates to false)
	var propertyValue BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for propertyValue")
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(2), objectTypeArgument, propertyIdentifier.GetValue(), (CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() interface{} { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() interface{} { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyValue' field of BACnetPropertyWriteDefinition")
		default:
			propertyValue = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for propertyValue")
			}
		}
	}

	// Optional Field (priority) (Can be skipped, if a given expression evaluates to false)
	var priority BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for priority")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priority' field of BACnetPropertyWriteDefinition")
		default:
			priority = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for priority")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyWriteDefinition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyWriteDefinition")
	}

	// Create the instance
	return &_BACnetPropertyWriteDefinition{
		ObjectTypeArgument: objectTypeArgument,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		PropertyValue:      propertyValue,
		Priority:           priority,
	}, nil
}

func (m *_BACnetPropertyWriteDefinition) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPropertyWriteDefinition"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyWriteDefinition")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
	}
	_propertyIdentifierErr := writeBuffer.WriteSerializable(m.GetPropertyIdentifier())
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	if m.GetArrayIndex() != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for arrayIndex")
		}
		arrayIndex = m.GetArrayIndex()
		_arrayIndexErr := writeBuffer.WriteSerializable(arrayIndex)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for arrayIndex")
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if the value is null)
	var propertyValue BACnetConstructedData = nil
	if m.GetPropertyValue() != nil {
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyValue")
		}
		propertyValue = m.GetPropertyValue()
		_propertyValueErr := writeBuffer.WriteSerializable(propertyValue)
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyValue")
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}
	}

	// Optional Field (priority) (Can be skipped, if the value is null)
	var priority BACnetContextTagUnsignedInteger = nil
	if m.GetPriority() != nil {
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priority")
		}
		priority = m.GetPriority()
		_priorityErr := writeBuffer.WriteSerializable(priority)
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priority")
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyWriteDefinition"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyWriteDefinition")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPropertyWriteDefinition) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetPropertyWriteDefinition) isBACnetPropertyWriteDefinition() bool {
	return true
}

func (m *_BACnetPropertyWriteDefinition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
