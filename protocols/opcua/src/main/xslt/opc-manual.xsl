<?xml version="1.0" encoding="UTF-8"?>
<!--
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
-->
<xsl:stylesheet version="2.0"
                xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
                xmlns:xs="http://www.w3.org/2001/XMLSchema"
                xmlns:opc="http://opcfoundation.org/BinarySchema/"
                xmlns:plc4x="https://plc4x.apache.org/"
                xmlns:map="http://www.w3.org/2005/xpath-functions/map"
                xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                xmlns:ua="http://opcfoundation.org/UA/"
                xmlns:tns="http://opcfoundation.org/UA/"
                xmlns:node="http://opcfoundation.org/UA/2011/03/UANodeSet.xsd">

    <xsl:output
        method="text"
        indent="no"
        encoding="utf-8"
    />

    <xsl:import href="opc-common.xsl"/>

    <xsl:variable name="originaldoc" select="/"/>

    <xsl:param name="services"></xsl:param>
    <xsl:param name="file" select="document($services)"/>

    <xsl:template match="/">
// Remark: The different fields are encoded in Little-endian.

[type 'OpcuaAPU' [bit 'response']
    [simple         MessagePDU   'message' ['response']]
]

[discriminatedType 'MessagePDU' [bit 'response']
    [discriminator string '24'          'messageType']
    [typeSwitch 'messageType','response'
        ['HEL','false'     OpcuaHelloRequest
            [simple          string '8'         'chunk']
            [implicit        int 32             'messageSize' 'lengthInBytes']
            [simple          int 32             'version']
            [simple          int 32             'receiveBufferSize']
            [simple          int 32             'sendBufferSize']
            [simple          int 32             'maxMessageSize']
            [simple          int 32             'maxChunkCount']
            [simple          PascalString       'endpoint']
        ]
        ['ACK','true'     OpcuaAcknowledgeResponse
            [simple          string '8'         'chunk']
            [implicit        int 32             'messageSize' 'lengthInBytes']
            [simple          int 32             'version']
            [simple          int 32             'receiveBufferSize']
            [simple          int 32             'sendBufferSize']
            [simple          int 32             'maxMessageSize']
            [simple          int 32             'maxChunkCount']
        ]
        ['OPN','false'     OpcuaOpenRequest
            [simple          string '8'         'chunk']
            [implicit        int 32             'messageSize' 'lengthInBytes']
            [simple          int 32             'secureChannelId']
            [simple          PascalString       'endpoint']
            [simple          PascalByteString   'senderCertificate']
            [simple          PascalByteString   'receiverCertificateThumbprint']
            [simple          int 32             'sequenceNumber']
            [simple          int 32             'requestId']
            [array           int 8              'message' count 'messageSize - (endpoint.stringLength == -1 ? 0 : endpoint.stringLength ) - (senderCertificate.stringLength == -1 ? 0 : senderCertificate.stringLength) - (receiverCertificateThumbprint.stringLength == -1 ? 0 : receiverCertificateThumbprint.stringLength) - 32']
       ]
       ['OPN','true'     OpcuaOpenResponse
           [simple          string '8'         'chunk']
           [implicit        int 32             'messageSize' 'lengthInBytes']
           [simple          int 32             'secureChannelId']
           [simple          PascalString       'securityPolicyUri']
           [simple          PascalByteString   'senderCertificate']
           [simple          PascalByteString   'receiverCertificateThumbprint']
           [simple          int 32             'sequenceNumber']
           [simple          int 32             'requestId']
           [array           int 8              'message' count 'messageSize - (securityPolicyUri.stringLength == -1 ? 0 : securityPolicyUri.stringLength) - (senderCertificate.stringLength == -1 ? 0 : senderCertificate.stringLength) - (receiverCertificateThumbprint.stringLength == -1 ? 0 : receiverCertificateThumbprint.stringLength) - 32']
       ]
       ['CLO','false'     OpcuaCloseRequest
           [simple          string '8'         'chunk']
           [implicit        int 32             'messageSize' 'lengthInBytes']
           [simple          int 32             'secureChannelId']
           [simple          int 32             'secureTokenId']
           [simple          int 32             'sequenceNumber']
           [simple          int 32             'requestId']
           [simple          ExtensionObject       'message']
       ]
       ['MSG','false'     OpcuaMessageRequest
           [simple          string '8'         'chunk']
           [implicit        int 32             'messageSize' 'lengthInBytes']
           [simple          int 32             'secureChannelId']
           [simple          int 32             'secureTokenId']
           [simple          int 32             'sequenceNumber']
           [simple          int 32             'requestId']
           [array           int 8              'message' count 'messageSize - 24']
       ]
       ['MSG','true'     OpcuaMessageResponse
           [simple          string '8'         'chunk']
           [implicit        int 32             'messageSize' 'lengthInBytes']
           [simple          int 32             'secureChannelId']
           [simple          int 32             'secureTokenId']
           [simple          int 32             'sequenceNumber']
           [simple          int 32             'requestId']
           [array           int 8              'message' count 'messageSize - 24']
       ]
    ]
]

[type 'ByteStringArray'
    [simple int 32 'arrayLength']
    [array uint 8 'value' count 'arrayLength']
]

[type 'GuidValue'
    [simple uint 32 'data1']
    [simple uint 16 'data2']
    [simple uint 16 'data3']
    [array int 8 'data4' count '2']
    [array int 8 'data5' count '6']
]

[type 'ExpandedNodeId'
    [simple bit 'namespaceURISpecified']
    [simple bit 'serverIndexSpecified']
    [simple NodeIdTypeDefinition 'nodeId']
    [virtual string '-1' 'utf-8' 'identifier' 'nodeId.identifier']
    [optional PascalString 'namespaceURI' 'namespaceURISpecified']
    [optional uint 32 'serverIndex' 'serverIndexSpecified']
]

[type 'ExtensionHeader'
    [reserved int 5 '0x00']
    [simple bit 'xmlbody']
    [simple bit 'binaryBody]
]

[type 'ExtensionObject'
    //A serialized object prefixed with its data type identifier.
    [simple ExpandedNodeId 'typeId']
    [virtual string '-1' 'identifier' 'typeId.identifier']
    [simple ExtensionObjectDefinition 'body' ['identifier']]
]

[discriminatedType 'ExtensionObjectDefinition' [string '-1' 'identifier']
    [typeSwitch 'identifier'
        ['0' NullExtension
            [reserved int 5 '0x00']
            [simple bit 'xmlbody']
            [simple bit 'binaryBody]
            [simple bit 'typeIdSpecified']
        ]

        <xsl:for-each select="/opc:TypeDictionary/opc:StructuredType[(@BaseType = 'ua:ExtensionObject') and not(@Name = 'UserIdentityToken') and not(@Name = 'PublishedDataSetDataType') and not(@Name = 'DataSetReaderDataType')]">
            <xsl:message><xsl:value-of select="@Name"/></xsl:message>
            <xsl:variable name="extensionName" select="@Name"/>
            <xsl:apply-templates select="$file/node:UANodeSet/node:UADataType[@BrowseName=$extensionName]"/>
        </xsl:for-each>

        ['316' UserIdentityToken
            [simple UserIdentityTokenDefinition 'userIdentityTokenDefinition']
        ]
    ]
]

[discriminatedType 'UserIdentityTokenDefinition'
    [implicit int 32 'policyIdLength' 'policyId.length']
    [discriminator string '-1' 'policyId']
    [typeSwitch 'policyId'
        ['none' AnonymousIdentityToken
        ]
        ['username' UserNameIdentityToken
            [simple PascalString 'userName']
            [simple PascalByteString 'password']
            [simple PascalString 'encryptionAlgorithm']
        ]
        ['certificate' X509IdentityToken
            [simple PascalByteString 'certificateData']
        ]
        ['identity' IssuedIdentityToken
            [simple PascalByteString 'tokenData']
            [simple PascalString 'encryptionAlgorithm']
        ]
    ]
]


[discriminatedType 'Variant'
    [simple bit 'arrayLengthSpecified']
    [simple bit 'arrayDimensionsSpecified']
    [discriminator uint 6 'VariantType']
    [typeSwitch 'VariantType','arrayLengthSpecified'
        ['1' VariantBoolean [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array int 8 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['2' VariantSByte [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array int 8 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['3' VariantByte [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array uint 8 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['4' VariantInt16 [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array int 16 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['5' VariantUInt16 [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array uint 16 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['6' VariantInt32 [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array int 32 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['7' VariantUInt32 [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array uint 32 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['8' VariantInt64 [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array int 64 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['9' VariantUInt64 [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array uint 64 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['10' VariantFloat [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array float 8.23 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['11' VariantDouble [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array float 11.52 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['12' VariantString [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array PascalString 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['13' VariantDateTime [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array int 64 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['14' VariantGuid [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array GuidValue 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['15' VariantByteString [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array ByteStringArray 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['16' VariantXmlElement [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array PascalString 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['17' VariantNodeId [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array NodeId 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['18' VariantExpandedNodeId [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array ExpandedNodeId 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['19' VariantStatusCode [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array StatusCode 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['20' VariantQualifiedName [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array QualifiedName 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['21' VariantLocalizedText [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array LocalizedText 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['22' VariantExtensionObject [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array ExtensionObject 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['23' VariantDataValue [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array DataValue 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['24' VariantVariant [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array Variant 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
        ['25' VariantDiagnosticInfo [bit 'arrayLengthSpecified']
            [optional int 32 'arrayLength' 'arrayLengthSpecified']
            [array DiagnosticInfo 'value' count 'arrayLength == null ? 1 : arrayLength']
        ]
    ]
    [optional int 32 'noOfArrayDimensions' 'arrayDimensionsSpecified']
    [array bit 'arrayDimensions' count 'noOfArrayDimensions == null ? 0 : noOfArrayDimensions']
]

[discriminatedType 'NodeIdTypeDefinition'
    [abstract string '-1' 'identifier']
    [discriminator NodeIdType 'nodeType']
    [typeSwitch 'nodeType'
        ['nodeIdTypeTwoByte' NodeIdTwoByte
            [simple uint 8 'id']
            [virtual string '-1' 'identifier' 'id']
        ]
        ['nodeIdTypeFourByte' NodeIdFourByte
            [simple uint 8 'namespaceIndex']
            [simple uint 16 'id']
            [virtual string '-1' 'identifier' 'id']
        ]
        ['nodeIdTypeNumeric' NodeIdNumeric
            [simple uint 16 'namespaceIndex']
            [simple uint 32 'id']
            [virtual string '-1' 'identifier' 'id']
        ]
        ['nodeIdTypeString' NodeIdString
            [simple uint 16 'namespaceIndex']
            [simple PascalString 'id']
            [virtual string '-1' 'identifier' 'id.stringValue']
        ]
        ['nodeIdTypeGuid' NodeIdGuid
            [simple uint 16 'namespaceIndex']
            [simple string '-1' 'id']
            [virtual string '-1' 'identifier' 'id']
        ]
        ['nodeIdTypeByteString' NodeIdByteString
            [simple uint 16 'namespaceIndex']
            [simple PascalByteString 'id']
            [virtual string '-1' 'identifier' 'id.stringValue']
        ]
    ]
]

[type 'NodeId'
    [reserved int 2 '0x00']
    [simple NodeIdTypeDefinition 'nodeId']
    [virtual string '-1' 'id' 'nodeId.identifier']
]

[type 'PascalString'
    [implicit int 32 'sLength'          'stringValue.length == 0 ? -1 : stringValue.length']
    [simple string 'sLength == -1 ? 0 : sLength * 8' 'UTF-8' 'stringValue']
    [virtual  int 32 'stringLength'     'stringValue.length == -1 ? 0 : stringValue.length']
]

[type 'PascalByteString'
    [simple int 32 'stringLength']
    [array int 8 'stringValue' count 'stringLength == -1 ? 0 : stringLength' ]
]

[type 'Structure'

]

[type 'DataTypeDefinition'

]

<xsl:apply-templates select="/opc:TypeDictionary/opc:StructuredType[(@Name != 'ExtensionObject') and (@Name != 'Variant') and (@Name != 'NodeId') and (@Name != 'ExpandedNodeId') and not(@BaseType)]"/>
<xsl:apply-templates select="/opc:TypeDictionary/opc:EnumeratedType"/>
<xsl:apply-templates select="/opc:TypeDictionary/opc:OpaqueType"/>

[enum string '-1' 'OpcuaDataType' [uint 8 'variantType']
    ['IEC61131_NULL' NULL ['0']]
    ['IEC61131_BOOL' BOOL ['1']]
    ['IEC61131_BYTE' BYTE ['3']]
    ['IEC61131_SINT' SINT ['2']]
    ['IEC61131_INT' INT ['4']]
    ['IEC61131_DINT' DINT ['6']]
    ['IEC61131_LINT' LINT ['8']]
    ['IEC61131_USINT' USINT ['3']]
    ['IEC61131_UINT' UINT ['5']]
    ['IEC61131_UDINT' UDINT ['7']]
    ['IEC61131_ULINT' ULINT ['9']]
    ['IEC61131_REAL' REAL ['10']]
    ['IEC61131_LREAL' LREAL ['11']]
    ['IEC61131_TIME' TIME ['1']]
    ['IEC61131_LTIME' LTIME ['1']]
    ['IEC61131_DATE' DATE ['1']]
    ['IEC61131_LDATE' LDATE ['1']]
    ['IEC61131_TIME_OF_DAY' TIME_OF_DAY ['1']]
    ['IEC61131_LTIME_OF_DAY' LTIME_OF_DAY ['1']]
    ['IEC61131_DATE_AND_TIME' DATE_AND_TIME ['13']]
    ['IEC61131_LDATE_AND_TIME' LDATE_AND_TIME ['1']]
    ['IEC61131_CHAR' CHAR ['1']]
    ['IEC61131_WCHAR' WCHAR ['1']]
    ['IEC61131_STRING' STRING ['12']]
]

[enum string '-1' 'OpcuaIdentifierType'
    ['s' STRING_IDENTIFIER]
    ['i' NUMBER_IDENTIFIER]
    ['g' GUID_IDENTIFIER]
    ['b' BINARY_IDENTIFIER]
]

    </xsl:template>
</xsl:stylesheet>
