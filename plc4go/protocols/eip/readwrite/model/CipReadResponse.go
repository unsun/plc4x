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
)

// Code generated by code-generation. DO NOT EDIT.

// CipReadResponse is the corresponding interface of CipReadResponse
type CipReadResponse interface {
	utils.LengthAware
	utils.Serializable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetExtStatus returns ExtStatus (property field)
	GetExtStatus() uint8
	// GetDataType returns DataType (property field)
	GetDataType() CIPDataTypeCode
	// GetData returns Data (property field)
	GetData() []byte
}

// CipReadResponseExactly can be used when we want exactly this type and not a type which fulfills CipReadResponse.
// This is useful for switch cases.
type CipReadResponseExactly interface {
	CipReadResponse
	isCipReadResponse() bool
}

// _CipReadResponse is the data-structure of this message
type _CipReadResponse struct {
	*_CipService
	Status    uint8
	ExtStatus uint8
	DataType  CIPDataTypeCode
	Data      []byte
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipReadResponse) GetService() uint8 {
	return 0xCC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipReadResponse) InitializeParent(parent CipService) {}

func (m *_CipReadResponse) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipReadResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_CipReadResponse) GetExtStatus() uint8 {
	return m.ExtStatus
}

func (m *_CipReadResponse) GetDataType() CIPDataTypeCode {
	return m.DataType
}

func (m *_CipReadResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipReadResponse factory function for _CipReadResponse
func NewCipReadResponse(status uint8, extStatus uint8, dataType CIPDataTypeCode, data []byte, serviceLen uint16) *_CipReadResponse {
	_result := &_CipReadResponse{
		Status:      status,
		ExtStatus:   extStatus,
		DataType:    dataType,
		Data:        data,
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipReadResponse(structType interface{}) CipReadResponse {
	if casted, ok := structType.(CipReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipReadResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipReadResponse) GetTypeName() string {
	return "CipReadResponse"
}

func (m *_CipReadResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CipReadResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (extStatus)
	lengthInBits += 8

	// Simple field (dataType)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_CipReadResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CipReadResponseParse(readBuffer utils.ReadBuffer, serviceLen uint16) (CipReadResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipReadResponse")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of CipReadResponse")
	}
	status := _status

	// Simple Field (extStatus)
	_extStatus, _extStatusErr := readBuffer.ReadUint8("extStatus", 8)
	if _extStatusErr != nil {
		return nil, errors.Wrap(_extStatusErr, "Error parsing 'extStatus' field of CipReadResponse")
	}
	extStatus := _extStatus

	// Simple Field (dataType)
	if pullErr := readBuffer.PullContext("dataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataType")
	}
	_dataType, _dataTypeErr := CIPDataTypeCodeParse(readBuffer)
	if _dataTypeErr != nil {
		return nil, errors.Wrap(_dataTypeErr, "Error parsing 'dataType' field of CipReadResponse")
	}
	dataType := _dataType
	if closeErr := readBuffer.CloseContext("dataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataType")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(serviceLen) - uint16(uint16(6)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of CipReadResponse")
	}

	if closeErr := readBuffer.CloseContext("CipReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipReadResponse")
	}

	// Create a partially initialized instance
	_child := &_CipReadResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		Status:         status,
		ExtStatus:      extStatus,
		DataType:       dataType,
		Data:           data,
		reservedField0: reservedField0,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipReadResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipReadResponse")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (status)
		status := uint8(m.GetStatus())
		_statusErr := writeBuffer.WriteUint8("status", 8, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (extStatus)
		extStatus := uint8(m.GetExtStatus())
		_extStatusErr := writeBuffer.WriteUint8("extStatus", 8, (extStatus))
		if _extStatusErr != nil {
			return errors.Wrap(_extStatusErr, "Error serializing 'extStatus' field")
		}

		// Simple Field (dataType)
		if pushErr := writeBuffer.PushContext("dataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dataType")
		}
		_dataTypeErr := writeBuffer.WriteSerializable(m.GetDataType())
		if popErr := writeBuffer.PopContext("dataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dataType")
		}
		if _dataTypeErr != nil {
			return errors.Wrap(_dataTypeErr, "Error serializing 'dataType' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("CipReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipReadResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CipReadResponse) isCipReadResponse() bool {
	return true
}

func (m *_CipReadResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
