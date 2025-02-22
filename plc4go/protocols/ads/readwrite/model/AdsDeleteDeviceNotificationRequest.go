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

// AdsDeleteDeviceNotificationRequest is the corresponding interface of AdsDeleteDeviceNotificationRequest
type AdsDeleteDeviceNotificationRequest interface {
	utils.LengthAware
	utils.Serializable
	AdsData
	// GetNotificationHandle returns NotificationHandle (property field)
	GetNotificationHandle() uint32
}

// AdsDeleteDeviceNotificationRequestExactly can be used when we want exactly this type and not a type which fulfills AdsDeleteDeviceNotificationRequest.
// This is useful for switch cases.
type AdsDeleteDeviceNotificationRequestExactly interface {
	AdsDeleteDeviceNotificationRequest
	isAdsDeleteDeviceNotificationRequest() bool
}

// _AdsDeleteDeviceNotificationRequest is the data-structure of this message
type _AdsDeleteDeviceNotificationRequest struct {
	*_AdsData
	NotificationHandle uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDeleteDeviceNotificationRequest) GetCommandId() CommandId {
	return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
}

func (m *_AdsDeleteDeviceNotificationRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDeleteDeviceNotificationRequest) InitializeParent(parent AdsData) {}

func (m *_AdsDeleteDeviceNotificationRequest) GetParent() AdsData {
	return m._AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDeleteDeviceNotificationRequest) GetNotificationHandle() uint32 {
	return m.NotificationHandle
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDeleteDeviceNotificationRequest factory function for _AdsDeleteDeviceNotificationRequest
func NewAdsDeleteDeviceNotificationRequest(notificationHandle uint32) *_AdsDeleteDeviceNotificationRequest {
	_result := &_AdsDeleteDeviceNotificationRequest{
		NotificationHandle: notificationHandle,
		_AdsData:           NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDeleteDeviceNotificationRequest(structType interface{}) AdsDeleteDeviceNotificationRequest {
	if casted, ok := structType.(AdsDeleteDeviceNotificationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDeleteDeviceNotificationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDeleteDeviceNotificationRequest) GetTypeName() string {
	return "AdsDeleteDeviceNotificationRequest"
}

func (m *_AdsDeleteDeviceNotificationRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsDeleteDeviceNotificationRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (notificationHandle)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsDeleteDeviceNotificationRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsDeleteDeviceNotificationRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsDeleteDeviceNotificationRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDeleteDeviceNotificationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDeleteDeviceNotificationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (notificationHandle)
	_notificationHandle, _notificationHandleErr := readBuffer.ReadUint32("notificationHandle", 32)
	if _notificationHandleErr != nil {
		return nil, errors.Wrap(_notificationHandleErr, "Error parsing 'notificationHandle' field of AdsDeleteDeviceNotificationRequest")
	}
	notificationHandle := _notificationHandle

	if closeErr := readBuffer.CloseContext("AdsDeleteDeviceNotificationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDeleteDeviceNotificationRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsDeleteDeviceNotificationRequest{
		_AdsData:           &_AdsData{},
		NotificationHandle: notificationHandle,
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsDeleteDeviceNotificationRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeleteDeviceNotificationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDeleteDeviceNotificationRequest")
		}

		// Simple Field (notificationHandle)
		notificationHandle := uint32(m.GetNotificationHandle())
		_notificationHandleErr := writeBuffer.WriteUint32("notificationHandle", 32, (notificationHandle))
		if _notificationHandleErr != nil {
			return errors.Wrap(_notificationHandleErr, "Error serializing 'notificationHandle' field")
		}

		if popErr := writeBuffer.PopContext("AdsDeleteDeviceNotificationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDeleteDeviceNotificationRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsDeleteDeviceNotificationRequest) isAdsDeleteDeviceNotificationRequest() bool {
	return true
}

func (m *_AdsDeleteDeviceNotificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
