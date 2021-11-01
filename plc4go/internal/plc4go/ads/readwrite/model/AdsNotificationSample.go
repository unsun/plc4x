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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type AdsNotificationSample struct {
	NotificationHandle uint32
	SampleSize         uint32
	Data               []byte
}

// The corresponding interface
type IAdsNotificationSample interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewAdsNotificationSample(notificationHandle uint32, sampleSize uint32, data []byte) *AdsNotificationSample {
	return &AdsNotificationSample{NotificationHandle: notificationHandle, SampleSize: sampleSize, Data: data}
}

func CastAdsNotificationSample(structType interface{}) *AdsNotificationSample {
	castFunc := func(typ interface{}) *AdsNotificationSample {
		if casted, ok := typ.(AdsNotificationSample); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsNotificationSample); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsNotificationSample) GetTypeName() string {
	return "AdsNotificationSample"
}

func (m *AdsNotificationSample) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsNotificationSample) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (notificationHandle)
	lengthInBits += 32

	// Simple field (sampleSize)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AdsNotificationSample) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsNotificationSampleParse(readBuffer utils.ReadBuffer) (*AdsNotificationSample, error) {
	if pullErr := readBuffer.PullContext("AdsNotificationSample"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (notificationHandle)
	notificationHandle, _notificationHandleErr := readBuffer.ReadUint32("notificationHandle", 32)
	if _notificationHandleErr != nil {
		return nil, errors.Wrap(_notificationHandleErr, "Error parsing 'notificationHandle' field")
	}

	// Simple Field (sampleSize)
	sampleSize, _sampleSizeErr := readBuffer.ReadUint32("sampleSize", 32)
	if _sampleSizeErr != nil {
		return nil, errors.Wrap(_sampleSizeErr, "Error parsing 'sampleSize' field")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(sampleSize)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("AdsNotificationSample"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAdsNotificationSample(notificationHandle, sampleSize, data), nil
}

func (m *AdsNotificationSample) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AdsNotificationSample"); pushErr != nil {
		return pushErr
	}

	// Simple Field (notificationHandle)
	notificationHandle := uint32(m.NotificationHandle)
	_notificationHandleErr := writeBuffer.WriteUint32("notificationHandle", 32, (notificationHandle))
	if _notificationHandleErr != nil {
		return errors.Wrap(_notificationHandleErr, "Error serializing 'notificationHandle' field")
	}

	// Simple Field (sampleSize)
	sampleSize := uint32(m.SampleSize)
	_sampleSizeErr := writeBuffer.WriteUint32("sampleSize", 32, (sampleSize))
	if _sampleSizeErr != nil {
		return errors.Wrap(_sampleSizeErr, "Error serializing 'sampleSize' field")
	}

	// Array Field (data)
	if m.Data != nil {
		// Byte Array field (data)
		_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
		}
	}

	if popErr := writeBuffer.PopContext("AdsNotificationSample"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AdsNotificationSample) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
