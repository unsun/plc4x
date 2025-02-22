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

// BACnetServiceAckVTData is the corresponding interface of BACnetServiceAckVTData
type BACnetServiceAckVTData interface {
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetVtSessionIdentifier returns VtSessionIdentifier (property field)
	GetVtSessionIdentifier() BACnetApplicationTagUnsignedInteger
	// GetVtNewData returns VtNewData (property field)
	GetVtNewData() BACnetApplicationTagOctetString
	// GetVtDataFlag returns VtDataFlag (property field)
	GetVtDataFlag() BACnetApplicationTagUnsignedInteger
}

// BACnetServiceAckVTDataExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckVTData.
// This is useful for switch cases.
type BACnetServiceAckVTDataExactly interface {
	BACnetServiceAckVTData
	isBACnetServiceAckVTData() bool
}

// _BACnetServiceAckVTData is the data-structure of this message
type _BACnetServiceAckVTData struct {
	*_BACnetServiceAck
	VtSessionIdentifier BACnetApplicationTagUnsignedInteger
	VtNewData           BACnetApplicationTagOctetString
	VtDataFlag          BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckVTData) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_DATA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckVTData) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckVTData) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckVTData) GetVtSessionIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.VtSessionIdentifier
}

func (m *_BACnetServiceAckVTData) GetVtNewData() BACnetApplicationTagOctetString {
	return m.VtNewData
}

func (m *_BACnetServiceAckVTData) GetVtDataFlag() BACnetApplicationTagUnsignedInteger {
	return m.VtDataFlag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckVTData factory function for _BACnetServiceAckVTData
func NewBACnetServiceAckVTData(vtSessionIdentifier BACnetApplicationTagUnsignedInteger, vtNewData BACnetApplicationTagOctetString, vtDataFlag BACnetApplicationTagUnsignedInteger, serviceAckLength uint16) *_BACnetServiceAckVTData {
	_result := &_BACnetServiceAckVTData{
		VtSessionIdentifier: vtSessionIdentifier,
		VtNewData:           vtNewData,
		VtDataFlag:          vtDataFlag,
		_BACnetServiceAck:   NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckVTData(structType interface{}) BACnetServiceAckVTData {
	if casted, ok := structType.(BACnetServiceAckVTData); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckVTData); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckVTData) GetTypeName() string {
	return "BACnetServiceAckVTData"
}

func (m *_BACnetServiceAckVTData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetServiceAckVTData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (vtSessionIdentifier)
	lengthInBits += m.VtSessionIdentifier.GetLengthInBits()

	// Simple field (vtNewData)
	lengthInBits += m.VtNewData.GetLengthInBits()

	// Simple field (vtDataFlag)
	lengthInBits += m.VtDataFlag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetServiceAckVTData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckVTDataParse(readBuffer utils.ReadBuffer, serviceAckLength uint16) (BACnetServiceAckVTData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckVTData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckVTData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vtSessionIdentifier)
	if pullErr := readBuffer.PullContext("vtSessionIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtSessionIdentifier")
	}
	_vtSessionIdentifier, _vtSessionIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _vtSessionIdentifierErr != nil {
		return nil, errors.Wrap(_vtSessionIdentifierErr, "Error parsing 'vtSessionIdentifier' field of BACnetServiceAckVTData")
	}
	vtSessionIdentifier := _vtSessionIdentifier.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("vtSessionIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtSessionIdentifier")
	}

	// Simple Field (vtNewData)
	if pullErr := readBuffer.PullContext("vtNewData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtNewData")
	}
	_vtNewData, _vtNewDataErr := BACnetApplicationTagParse(readBuffer)
	if _vtNewDataErr != nil {
		return nil, errors.Wrap(_vtNewDataErr, "Error parsing 'vtNewData' field of BACnetServiceAckVTData")
	}
	vtNewData := _vtNewData.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("vtNewData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtNewData")
	}

	// Simple Field (vtDataFlag)
	if pullErr := readBuffer.PullContext("vtDataFlag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtDataFlag")
	}
	_vtDataFlag, _vtDataFlagErr := BACnetApplicationTagParse(readBuffer)
	if _vtDataFlagErr != nil {
		return nil, errors.Wrap(_vtDataFlagErr, "Error parsing 'vtDataFlag' field of BACnetServiceAckVTData")
	}
	vtDataFlag := _vtDataFlag.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("vtDataFlag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtDataFlag")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckVTData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckVTData")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckVTData{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		VtSessionIdentifier: vtSessionIdentifier,
		VtNewData:           vtNewData,
		VtDataFlag:          vtDataFlag,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckVTData) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckVTData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckVTData")
		}

		// Simple Field (vtSessionIdentifier)
		if pushErr := writeBuffer.PushContext("vtSessionIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vtSessionIdentifier")
		}
		_vtSessionIdentifierErr := writeBuffer.WriteSerializable(m.GetVtSessionIdentifier())
		if popErr := writeBuffer.PopContext("vtSessionIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vtSessionIdentifier")
		}
		if _vtSessionIdentifierErr != nil {
			return errors.Wrap(_vtSessionIdentifierErr, "Error serializing 'vtSessionIdentifier' field")
		}

		// Simple Field (vtNewData)
		if pushErr := writeBuffer.PushContext("vtNewData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vtNewData")
		}
		_vtNewDataErr := writeBuffer.WriteSerializable(m.GetVtNewData())
		if popErr := writeBuffer.PopContext("vtNewData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vtNewData")
		}
		if _vtNewDataErr != nil {
			return errors.Wrap(_vtNewDataErr, "Error serializing 'vtNewData' field")
		}

		// Simple Field (vtDataFlag)
		if pushErr := writeBuffer.PushContext("vtDataFlag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vtDataFlag")
		}
		_vtDataFlagErr := writeBuffer.WriteSerializable(m.GetVtDataFlag())
		if popErr := writeBuffer.PopContext("vtDataFlag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vtDataFlag")
		}
		if _vtDataFlagErr != nil {
			return errors.Wrap(_vtDataFlagErr, "Error serializing 'vtDataFlag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckVTData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckVTData")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetServiceAckVTData) isBACnetServiceAckVTData() bool {
	return true
}

func (m *_BACnetServiceAckVTData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
