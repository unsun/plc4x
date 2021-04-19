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

package org.apache.plc4x.java.spi.connection;

import static org.apache.plc4x.java.spi.configuration.ConfigurationFactory.*;

import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandler;
import io.netty.channel.ChannelPipeline;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.listener.EventListener;
import org.apache.plc4x.java.spi.EventListenerMessageCodec;
import org.apache.plc4x.java.spi.Plc4xNettyWrapper;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.configuration.Configuration;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.context.ToolingContext;
import org.apache.plc4x.java.spi.generation.Message;
import org.apache.plc4x.java.spi.generation.MessageIO;

import java.lang.reflect.InvocationTargetException;
import java.util.List;
import java.util.function.Consumer;
import java.util.function.Supplier;
import java.util.function.ToIntFunction;

/**
 * Builds a Protocol Stack.
 */
public class SingleProtocolStackConfigurer<BASE_PACKET_CLASS extends Message> implements ProtocolStackConfigurer<BASE_PACKET_CLASS> {

    private final Class<BASE_PACKET_CLASS> basePacketClass;
    private final boolean bigEndian;
    private final Supplier<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocolSupplier;
    private final Supplier<? extends DriverContext> driverContextSupplier;
    private final MessageIO<BASE_PACKET_CLASS, BASE_PACKET_CLASS> protocolIO;
    private final Supplier<? extends ToIntFunction<ByteBuf>> packetSizeEstimatorSupplier;
    private final Supplier<? extends Consumer<ByteBuf>> corruptPacketRemoverSupplier;
    private final Object[] parserArgs;

    public static <BPC extends Message> SingleProtocolStackBuilder<BPC> builder(Class<BPC> basePacketClass, Class<? extends MessageIO<BPC, BPC>> messageIoClass, Configuration configuration) {
        return new SingleProtocolStackBuilder<>(basePacketClass, messageIoClass, configuration);
    }

    /** Only accessible via Builder */
    SingleProtocolStackConfigurer(Class<BASE_PACKET_CLASS> basePacketClass,
                                  boolean bigEndian,
                                  Object[] parserArgs,
                                  Supplier<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocolSupplier,
                                  Supplier<? extends DriverContext> driverContextSupplier,
                                  MessageIO<BASE_PACKET_CLASS, BASE_PACKET_CLASS> protocolIO,
                                  Supplier<? extends ToIntFunction<ByteBuf>> packetSizeEstimatorSupplier,
                                  Supplier<? extends Consumer<ByteBuf>> corruptPacketRemoverSupplier) {
        this.basePacketClass = basePacketClass;
        this.bigEndian = bigEndian;
        this.parserArgs = parserArgs;
        this.protocolSupplier = protocolSupplier;
        this.driverContextSupplier = driverContextSupplier;
        this.protocolIO = protocolIO;
        this.packetSizeEstimatorSupplier = packetSizeEstimatorSupplier;
        this.corruptPacketRemoverSupplier = corruptPacketRemoverSupplier;
    }

    private ChannelHandler getMessageCodec() {
        return new GeneratedProtocolMessageCodec<>(basePacketClass, protocolIO, bigEndian, parserArgs,
            packetSizeEstimatorSupplier.get(), corruptPacketRemoverSupplier.get());
    }

    /** Applies the given Stack to the Pipeline */
    @Override
    public Plc4xProtocolBase<BASE_PACKET_CLASS> configurePipeline(ChannelPipeline pipeline, boolean passive, List<EventListener> listeners) {
        pipeline.addLast(getMessageCodec());
        pipeline.addLast(new EventListenerMessageCodec(listeners));
        Plc4xProtocolBase<BASE_PACKET_CLASS> protocol = protocolSupplier.get();
        DriverContext driverContext = driverContextSupplier.get();
        if (driverContext != null) {
            protocol.setDriverContext(driverContext);
        }
        Plc4xNettyWrapper<BASE_PACKET_CLASS> context = new Plc4xNettyWrapper<>(pipeline, passive, protocol, basePacketClass);
        pipeline.addLast(context);
        return protocol;
    }

    /**
     * Used to Build Instances of {@link SingleProtocolStackConfigurer}.
     *
     * @param <BASE_PACKET_CLASS> Type of Created Message that is Exchanged.
     */
    public static final class SingleProtocolStackBuilder<BASE_PACKET_CLASS extends Message> {

        private final Class<BASE_PACKET_CLASS> basePacketClass;
        private final Class<? extends MessageIO<BASE_PACKET_CLASS, BASE_PACKET_CLASS>> messageIoClass;
        private final Configuration configuration;
        private Supplier<? extends DriverContext> driverContextSupplier = () -> null;
        private boolean bigEndian = true;
        private Object[] parserArgs;
        private Supplier<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocolSupplier;
        private Supplier<? extends ToIntFunction<ByteBuf>> packetSizeEstimator = () -> null;
        private Supplier<? extends Consumer<ByteBuf>> corruptPacketRemover = () -> null;

        public SingleProtocolStackBuilder(Class<BASE_PACKET_CLASS> basePacketClass, Class<? extends MessageIO<BASE_PACKET_CLASS, BASE_PACKET_CLASS>> messageIoClass, Configuration configuration) {
            this.basePacketClass = basePacketClass;
            this.messageIoClass = messageIoClass;
            this.configuration = configuration;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withDriverContext(Class<? extends DriverContext> driverContextClass) {
            return withDriverContext(configuredType(driverContextClass));
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withDriverContext(Supplier<? extends DriverContext> driverContextClass) {
            this.driverContextSupplier = driverContextClass;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> littleEndian() {
            this.bigEndian = false;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withParserArgs(Object... parserArgs) {
            this.parserArgs = parserArgs;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withProtocol(Class<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocol) {
            return withProtocol(configuredType(protocol));
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withProtocol(Supplier<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocol) {
            this.protocolSupplier = protocol;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withPacketSizeEstimator(Class<? extends ToIntFunction<ByteBuf>> packetSizeEstimator) {
            return withPacketSizeEstimator(configuredType(packetSizeEstimator));
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withPacketSizeEstimator(Supplier<? extends ToIntFunction<ByteBuf>> packetSizeEstimator) {
            this.packetSizeEstimator = packetSizeEstimator;
            return this;
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withCorruptPacketRemover(Class<? extends Consumer<ByteBuf>> corruptPacketRemover) {
            return withCorruptPacketRemover(configuredType(corruptPacketRemover));
        }

        public SingleProtocolStackBuilder<BASE_PACKET_CLASS> withCorruptPacketRemover(Supplier<? extends Consumer<ByteBuf>> corruptPacketRemover) {
            this.corruptPacketRemover = corruptPacketRemover;
            return this;
        }

        public SingleProtocolStackConfigurer<BASE_PACKET_CLASS> build() {
            assert this.protocolSupplier != null;
            return new SingleProtocolStackConfigurer<>(
                basePacketClass, bigEndian, parserArgs, protocolSupplier,
                driverContextSupplier,
                new SimpleTypeSupplier<>(messageIoClass).get(),
                packetSizeEstimator,
                corruptPacketRemover
            );
        }

        protected final <T> Supplier<T> configuredType(Class<T> protocol) {
            return new ConfiguringSupplier<>(() -> configuration, new SimpleTypeSupplier<>(protocol));
        }

    }

    protected static class SimpleTypeSupplier<T> implements Supplier<T> {

        private final Class<T> type;

        public SimpleTypeSupplier(Class<T> type) {
            this.type = type;
        }

        @Override
        public T get() {
            try {
                return type.getDeclaredConstructor().newInstance();
            } catch (InstantiationException | IllegalAccessException | NoSuchMethodException e) {
                throw new PlcRuntimeException("Could not construct instance of " + type.getName(), e);
            } catch (InvocationTargetException e) {
                throw new PlcRuntimeException("Initialization of " + type.getName() + " instance raised an error", e);
            }
        }
    }

    protected static class ConfiguringSupplier<T> implements Supplier<T> {

        private final Supplier<Configuration> configuration;
        private final Supplier<T> delegate;

        public ConfiguringSupplier(Supplier<Configuration> configuration, Supplier<T> delegate) {
            this.configuration = configuration;
            this.delegate = delegate;
        }

        @Override
        public T get() {
            if (configuration == null) {
                return delegate.get();
            }

            return configure(configuration.get(), delegate.get());
        }
    }
}
