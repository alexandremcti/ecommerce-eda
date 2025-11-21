import {Channel, ChannelModel, connect} from 'amqplib'
import { Handler } from '../../in/handler/index.js'
import logger from '../../../main/shared/logs/index.js'

export interface MessageBroker {
    disconnect(): Promise<void>
    createPublisher(): Promise<void>
    createSubscriber(handler: Handler): Promise<void>
    notify(topic: string, message: string): Promise<boolean>
}

export type RabbitMQMessageBrokerInput = {
    channelModel: ChannelModel
    channel: Channel
}

export class RabbitMQMessageBroker implements MessageBroker {
    
    channelModel: ChannelModel
    channel: Channel
    private static exchange = 'qualificacao-criada-out-0'
    private static queue = 'pedido-criado-out-0'

    private constructor(input: RabbitMQMessageBrokerInput) {
        this.channelModel = input.channelModel
        this.channel = input.channel 
    }

    static async connect(url: string): Promise<MessageBroker> {
        // Logic to connect to RabbitMQ
        const channelModel = await connect(url);
        const channel = await channelModel.createChannel();
        return new RabbitMQMessageBroker({channel, channelModel})
    }
    
    async disconnect(): Promise<void> {
        // Logic to disconnect from RabbitMQ
        logger.info("Disconnecting from message broker...")
        await this.channelModel.close()
    }

    async createPublisher(): Promise<void> {
        await this.channel.assertExchange(RabbitMQMessageBroker.exchange, 'fanout', {
            durable: false
        })
    }

    async createSubscriber(handler: Handler): Promise<void> {
        const q = await this.channel.assertQueue(RabbitMQMessageBroker.queue, {
            exclusive: true
        })
        await this.channel.bindQueue(q.queue, RabbitMQMessageBroker.queue, '')
        await this.channel.consume(
            q.queue, (msg) => {
                if(msg?.content) handler.handle(msg.content.toString())
            }, 
            {noAck: true}
        )
    }
    
    async notify(topic: string, message: string): Promise<boolean> {
        return this.channel.publish(topic, '', Buffer.from(message))
    }
}