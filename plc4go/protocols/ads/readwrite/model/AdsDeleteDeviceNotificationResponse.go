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

// AdsDeleteDeviceNotificationResponse is the corresponding interface of AdsDeleteDeviceNotificationResponse
type AdsDeleteDeviceNotificationResponse interface {
	utils.LengthAware
	utils.Serializable
	AdsData
	// GetResult returns Result (property field)
	GetResult() ReturnCode
}

// AdsDeleteDeviceNotificationResponseExactly can be used when we want exactly this type and not a type which fulfills AdsDeleteDeviceNotificationResponse.
// This is useful for switch cases.
type AdsDeleteDeviceNotificationResponseExactly interface {
	AdsDeleteDeviceNotificationResponse
	isAdsDeleteDeviceNotificationResponse() bool
}

// _AdsDeleteDeviceNotificationResponse is the data-structure of this message
type _AdsDeleteDeviceNotificationResponse struct {
	*_AdsData
	Result ReturnCode
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) GetCommandId() CommandId {
	return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
}

func (m *_AdsDeleteDeviceNotificationResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) InitializeParent(parent AdsData) {}

func (m *_AdsDeleteDeviceNotificationResponse) GetParent() AdsData {
	return m._AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) GetResult() ReturnCode {
	return m.Result
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDeleteDeviceNotificationResponse factory function for _AdsDeleteDeviceNotificationResponse
func NewAdsDeleteDeviceNotificationResponse(result ReturnCode) *_AdsDeleteDeviceNotificationResponse {
	_result := &_AdsDeleteDeviceNotificationResponse{
		Result:   result,
		_AdsData: NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDeleteDeviceNotificationResponse(structType interface{}) AdsDeleteDeviceNotificationResponse {
	if casted, ok := structType.(AdsDeleteDeviceNotificationResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDeleteDeviceNotificationResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDeleteDeviceNotificationResponse) GetTypeName() string {
	return "AdsDeleteDeviceNotificationResponse"
}

func (m *_AdsDeleteDeviceNotificationResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsDeleteDeviceNotificationResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsDeleteDeviceNotificationResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsDeleteDeviceNotificationResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsDeleteDeviceNotificationResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDeleteDeviceNotificationResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDeleteDeviceNotificationResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for result")
	}
	_result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field of AdsDeleteDeviceNotificationResponse")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for result")
	}

	if closeErr := readBuffer.CloseContext("AdsDeleteDeviceNotificationResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDeleteDeviceNotificationResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsDeleteDeviceNotificationResponse{
		_AdsData: &_AdsData{},
		Result:   result,
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsDeleteDeviceNotificationResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeleteDeviceNotificationResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDeleteDeviceNotificationResponse")
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for result")
		}
		_resultErr := writeBuffer.WriteSerializable(m.GetResult())
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for result")
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		if popErr := writeBuffer.PopContext("AdsDeleteDeviceNotificationResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDeleteDeviceNotificationResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsDeleteDeviceNotificationResponse) isAdsDeleteDeviceNotificationResponse() bool {
	return true
}

func (m *_AdsDeleteDeviceNotificationResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
