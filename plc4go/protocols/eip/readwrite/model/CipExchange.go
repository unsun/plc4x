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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CipExchange_ITEMCOUNT uint16 = 0x02
const CipExchange_NULLPTR uint32 = 0x0
const CipExchange_UNCONNECTEDDATA uint16 = 0x00B2

// CipExchange is the corresponding interface of CipExchange
type CipExchange interface {
	utils.LengthAware
	utils.Serializable
	// GetService returns Service (property field)
	GetService() CipService
}

// CipExchangeExactly can be used when we want exactly this type and not a type which fulfills CipExchange.
// This is useful for switch cases.
type CipExchangeExactly interface {
	CipExchange
	isCipExchange() bool
}

// _CipExchange is the data-structure of this message
type _CipExchange struct {
	Service CipService

	// Arguments.
	ExchangeLen uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipExchange) GetService() CipService {
	return m.Service
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CipExchange) GetItemCount() uint16 {
	return CipExchange_ITEMCOUNT
}

func (m *_CipExchange) GetNullPtr() uint32 {
	return CipExchange_NULLPTR
}

func (m *_CipExchange) GetUnconnectedData() uint16 {
	return CipExchange_UNCONNECTEDDATA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipExchange factory function for _CipExchange
func NewCipExchange(service CipService, exchangeLen uint16) *_CipExchange {
	return &_CipExchange{Service: service, ExchangeLen: exchangeLen}
}

// Deprecated: use the interface for direct cast
func CastCipExchange(structType interface{}) CipExchange {
	if casted, ok := structType.(CipExchange); ok {
		return casted
	}
	if casted, ok := structType.(*CipExchange); ok {
		return *casted
	}
	return nil
}

func (m *_CipExchange) GetTypeName() string {
	return "CipExchange"
}

func (m *_CipExchange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CipExchange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (itemCount)
	lengthInBits += 16

	// Const Field (nullPtr)
	lengthInBits += 32

	// Const Field (unconnectedData)
	lengthInBits += 16

	// Implicit Field (size)
	lengthInBits += 16

	// Simple field (service)
	lengthInBits += m.Service.GetLengthInBits()

	return lengthInBits
}

func (m *_CipExchange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CipExchangeParse(readBuffer utils.ReadBuffer, exchangeLen uint16) (CipExchange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipExchange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipExchange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (itemCount)
	itemCount, _itemCountErr := readBuffer.ReadUint16("itemCount", 16)
	if _itemCountErr != nil {
		return nil, errors.Wrap(_itemCountErr, "Error parsing 'itemCount' field of CipExchange")
	}
	if itemCount != CipExchange_ITEMCOUNT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CipExchange_ITEMCOUNT) + " but got " + fmt.Sprintf("%d", itemCount))
	}

	// Const Field (nullPtr)
	nullPtr, _nullPtrErr := readBuffer.ReadUint32("nullPtr", 32)
	if _nullPtrErr != nil {
		return nil, errors.Wrap(_nullPtrErr, "Error parsing 'nullPtr' field of CipExchange")
	}
	if nullPtr != CipExchange_NULLPTR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CipExchange_NULLPTR) + " but got " + fmt.Sprintf("%d", nullPtr))
	}

	// Const Field (unconnectedData)
	unconnectedData, _unconnectedDataErr := readBuffer.ReadUint16("unconnectedData", 16)
	if _unconnectedDataErr != nil {
		return nil, errors.Wrap(_unconnectedDataErr, "Error parsing 'unconnectedData' field of CipExchange")
	}
	if unconnectedData != CipExchange_UNCONNECTEDDATA {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CipExchange_UNCONNECTEDDATA) + " but got " + fmt.Sprintf("%d", unconnectedData))
	}

	// Implicit Field (size) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	size, _sizeErr := readBuffer.ReadUint16("size", 16)
	_ = size
	if _sizeErr != nil {
		return nil, errors.Wrap(_sizeErr, "Error parsing 'size' field of CipExchange")
	}

	// Simple Field (service)
	if pullErr := readBuffer.PullContext("service"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for service")
	}
	_service, _serviceErr := CipServiceParse(readBuffer, uint16(uint16(exchangeLen)-uint16(uint16(10))))
	if _serviceErr != nil {
		return nil, errors.Wrap(_serviceErr, "Error parsing 'service' field of CipExchange")
	}
	service := _service.(CipService)
	if closeErr := readBuffer.CloseContext("service"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for service")
	}

	if closeErr := readBuffer.CloseContext("CipExchange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipExchange")
	}

	// Create the instance
	return &_CipExchange{
		ExchangeLen: exchangeLen,
		Service:     service,
	}, nil
}

func (m *_CipExchange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CipExchange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CipExchange")
	}

	// Const Field (itemCount)
	_itemCountErr := writeBuffer.WriteUint16("itemCount", 16, 0x02)
	if _itemCountErr != nil {
		return errors.Wrap(_itemCountErr, "Error serializing 'itemCount' field")
	}

	// Const Field (nullPtr)
	_nullPtrErr := writeBuffer.WriteUint32("nullPtr", 32, 0x0)
	if _nullPtrErr != nil {
		return errors.Wrap(_nullPtrErr, "Error serializing 'nullPtr' field")
	}

	// Const Field (unconnectedData)
	_unconnectedDataErr := writeBuffer.WriteUint16("unconnectedData", 16, 0x00B2)
	if _unconnectedDataErr != nil {
		return errors.Wrap(_unconnectedDataErr, "Error serializing 'unconnectedData' field")
	}

	// Implicit Field (size) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	size := uint16(uint16(uint16(uint16(m.GetLengthInBytes()))-uint16(uint16(8))) - uint16(uint16(2)))
	_sizeErr := writeBuffer.WriteUint16("size", 16, (size))
	if _sizeErr != nil {
		return errors.Wrap(_sizeErr, "Error serializing 'size' field")
	}

	// Simple Field (service)
	if pushErr := writeBuffer.PushContext("service"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for service")
	}
	_serviceErr := writeBuffer.WriteSerializable(m.GetService())
	if popErr := writeBuffer.PopContext("service"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for service")
	}
	if _serviceErr != nil {
		return errors.Wrap(_serviceErr, "Error serializing 'service' field")
	}

	if popErr := writeBuffer.PopContext("CipExchange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CipExchange")
	}
	return nil
}

////
// Arguments Getter

func (m *_CipExchange) GetExchangeLen() uint16 {
	return m.ExchangeLen
}

//
////

func (m *_CipExchange) isCipExchange() bool {
	return true
}

func (m *_CipExchange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
