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
package org.apache.plc4x.java.canopen.field;

import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class CANOpenNMTFieldTest {

    @Test
    public void testNodeSyntax() {
        final CANOpenNMTField canField = CANOpenNMTField.of("NMT:20");

        assertEquals(20, canField.getNodeId());
        assertFalse(canField.isWildcard());
    }

    @Test
    public void testWildcardSyntax() {
        CANOpenNMTField canField = CANOpenNMTField.of("NMT:0");

        assertEquals(0, canField.getNodeId());
        assertTrue(canField.isWildcard());

        // an simplified syntax
        canField = CANOpenNMTField.of("NMT");

        assertEquals(0, canField.getNodeId());
        assertTrue(canField.isWildcard());
    }

    @Test
    public void testInvalidSyntax() {
        assertThrows(PlcInvalidFieldException.class, () -> CANOpenNMTField.of("NMT:"));
    }

}