/*
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
*/
package org.apache.plc4x.java.opcua.protocol;

import org.apache.plc4x.java.api.messages.PlcSubscriptionEvent;
import org.apache.plc4x.java.api.messages.PlcSubscriptionResponse;
import org.apache.plc4x.java.api.model.PlcConsumerRegistration;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.opcua.field.OpcuaField;
import org.apache.plc4x.java.opcua.readwrite.*;
import org.apache.plc4x.java.opcua.readwrite.io.ExtensionObjectIO;
import org.apache.plc4x.java.opcua.readwrite.types.OpcuaNodeIdServices;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.apache.plc4x.java.spi.messages.DefaultPlcSubscriptionEvent;
import org.apache.plc4x.java.spi.messages.utils.ResponseItem;
import org.apache.plc4x.java.spi.model.DefaultPlcConsumerRegistration;
import org.apache.plc4x.java.spi.model.DefaultPlcSubscriptionHandle;
import org.apache.plc4x.java.spi.transaction.RequestTransactionManager;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Instant;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.atomic.AtomicLong;
import java.util.function.Consumer;

/**
 */
public class OpcuaSubscriptionHandle extends DefaultPlcSubscriptionHandle {

    private static final Logger LOGGER = LoggerFactory.getLogger(OpcuaSubscriptionHandle.class);

    private Set<Consumer<PlcSubscriptionEvent>> consumers = new HashSet<>();
    private OpcuaField field;
    private long clientHandle;

    private AtomicBoolean destroy = new AtomicBoolean(false);
    private OpcuaProtocolLogic plcSubscriber;
    private Long subscriptionId;
    private long cycleTime;

    /**
     * @param field    corresponding map key in the PLC4X request/reply map
     *
     */
    public OpcuaSubscriptionHandle(OpcuaProtocolLogic plcSubscriber, Long subscriptionId, OpcuaField field, long cycleTime) {
        super(plcSubscriber);
        this.field = field;
        this.subscriptionId = subscriptionId;
        this.plcSubscriber = plcSubscriber;
        this.cycleTime = cycleTime;
    }

    /**
     *
     * @return
     */
    public void startSubscriber() {
        CompletableFuture.supplyAsync(() -> {
            try {
                RequestTransactionManager tm = new RequestTransactionManager(1);
                LinkedList<Long> outstandingAcknowledgements = new LinkedList<>();
                LinkedList<Long> outstandingRequests = new LinkedList<>();
                AtomicInteger sequenceNumber = new AtomicInteger(1);
                while (!this.destroy.get()) {
                    LOGGER.info("SubscriberLoop");
                    try {
                        Thread.sleep(this.cycleTime);
                    } catch (InterruptedException e) {
                        LOGGER.info("Interrupted Exception");
                    }

                    int requestHandle = sequenceNumber.getAndIncrement();

                    RequestHeader requestHeader = new RequestHeader(this.plcSubscriber.getAuthenticationToken(),
                        OpcuaProtocolLogic.getCurrentDateTime(),
                        requestHandle,
                        0L,
                        OpcuaProtocolLogic.NULL_STRING,
                        OpcuaProtocolLogic.REQUEST_TIMEOUT_LONG,
                        OpcuaProtocolLogic.NULL_EXTENSION_OBJECT);

                    SubscriptionAcknowledgement[] acks = null;
                    int ackLength = -1;
                    LOGGER.info("-------------Oustanding Size1: - {}", outstandingAcknowledgements.size());
                    if (outstandingAcknowledgements.size() > 0) {
                        LOGGER.info("-------------Oustanding Size: - {}", outstandingAcknowledgements.size());
                        acks = new SubscriptionAcknowledgement[outstandingAcknowledgements.size()];
                        ackLength = outstandingAcknowledgements.size();
                        for (int i = 0; i < outstandingAcknowledgements.size(); i++) {
                            acks[i] = new SubscriptionAcknowledgement(this.subscriptionId, outstandingAcknowledgements.remove());
                        }
                    }

                    PublishRequest publishRequest = new PublishRequest(
                        requestHeader,
                        ackLength,
                        acks
                    );

                    ExpandedNodeId extExpandedNodeId = new ExpandedNodeId(false,           //Namespace Uri Specified
                        false,            //Server Index Specified
                        new NodeIdFourByte((short) 0, Integer.valueOf(publishRequest.getIdentifier())),
                        OpcuaProtocolLogic.NULL_STRING,
                        1L);

                    try {
                        WriteBuffer buffer = new WriteBuffer(publishRequest.getLengthInBytes(), true);
                        ExtensionObjectIO.staticSerialize(buffer, new ExtensionObject(
                            extExpandedNodeId,
                            null,
                            publishRequest));

                        int transactionId = this.plcSubscriber.getTransactionIdentifier();

                        OpcuaMessageRequest createMessageRequest = new OpcuaMessageRequest(OpcuaProtocolLogic.FINAL_CHUNK,
                            this.plcSubscriber.getChannelId(),
                            this.plcSubscriber.getTokenId(),
                            transactionId,
                            transactionId,
                            buffer.getData());

                        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
                        outstandingRequests.add((long) requestHandle);
                        transaction.submit(() -> this.plcSubscriber.getConversationContext().sendRequest(new OpcuaAPU(createMessageRequest))
                            .expectResponse(OpcuaAPU.class, OpcuaProtocolLogic.REQUEST_TIMEOUT)
                            .onError((p, e) -> LOGGER.error("Unable to maintain subscription"))
                            .check(p -> p.getMessage() instanceof OpcuaMessageResponse)
                            .unwrap(p -> (OpcuaMessageResponse) p.getMessage())
                            .handle(opcuaResponse -> {
                                PublishResponse responseMessage = null;
                                try {
                                    responseMessage = (PublishResponse) ExtensionObjectIO.staticParse(new ReadBuffer(opcuaResponse.getMessage(), true), false).getBody();
                                } catch (ParseException e) {
                                    e.printStackTrace();
                                }
                                outstandingRequests.remove(((ResponseHeader) responseMessage.getResponseHeader()).getRequestHandle());
                                outstandingAcknowledgements.add(((ResponseHeader) responseMessage.getResponseHeader()).getRequestHandle());

                                if (((NotificationMessage) responseMessage.getNotificationMessage()).getNoOfNotificationData() > 0) {

                                }
                                for (ExtensionObject notifications : ((NotificationMessage) responseMessage.getNotificationMessage()).getNotificationData()) {

                                }



                                // Pass the response back to the application.

                                // Finish the request-transaction.
                                transaction.endRequest();
                            }));
                    } catch (ParseException e) {
                        LOGGER.info("Unable to serialize subscription request");
                    }
                }


            } catch (Exception e) {
                LOGGER.error("Failed :(");
                e.printStackTrace();
            }
            return null;
        });
        return;
    }


    /**
     *
     * @return
     */
    public void stopSubscriber() {
        this.destroy.set(true);
    }


    public long getClientHandle() {
        return clientHandle;
    }

    /**
     * @param item
     * @param value
     */
    public void onSubscriptionValue(Object item, Object value) {
        consumers.forEach(plcSubscriptionEventConsumer -> {
            PlcResponseCode resultCode = PlcResponseCode.OK;
            PlcValue stringItem = null;
            /*if (value.getStatusCode() != StatusCode.GOOD) {
                resultCode = PlcResponseCode.NOT_FOUND;
            } else {
                stringItem = OpcuaTcpPlcConnection.encodePlcValue(value);

            }*/
            Map<String, ResponseItem<PlcValue>> fields = new HashMap<>();
            ResponseItem<PlcValue> newPair = new ResponseItem<>(resultCode, stringItem);
            fields.put(field.getIdentifier(), newPair);
            PlcSubscriptionEvent event = new DefaultPlcSubscriptionEvent(Instant.now(), fields);
            plcSubscriptionEventConsumer.accept(event);
        });
    }

    @Override
    public PlcConsumerRegistration register(Consumer<PlcSubscriptionEvent> consumer) {
        consumers.add(consumer);
        return null;
//        return () -> consumers.remove(consumer);
    }



}
