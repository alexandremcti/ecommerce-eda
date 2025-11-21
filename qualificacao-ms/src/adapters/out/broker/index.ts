import { QualificacaoOutput } from "../../../ports/out/qualificacao.js";
import { constants } from "../../../main/config/constants.js";
import logger from "../../../main/shared/logs/index.js";
import { QualificacaoOutputImp } from "./qualificacao-output.js";
import { MessageBroker, RabbitMQMessageBroker } from "./rabbitmq-broker.js";

export class BrokerComponent {
    private _broker: MessageBroker
    private _qualificacaoOuput: QualificacaoOutput

    private constructor(params: BrokerComponentProperties) {
        this._broker = params.broker
        this._qualificacaoOuput = params.qualificacaoOuput
    }

    get broker(): MessageBroker {
        return this._broker
    }

    get quaqualificacaoOuput(): QualificacaoOutput {
        return this._qualificacaoOuput
    }

    static async create(): Promise<BrokerComponent> {
        logger.info("Connecting to message broker...")
        const broker = await RabbitMQMessageBroker.connect(constants.broker.URL)
        const qualificacaoOuput = new QualificacaoOutputImp(broker)
        return new BrokerComponent({broker, qualificacaoOuput})
    }
}

type BrokerComponentProperties = {
    broker: MessageBroker
    qualificacaoOuput: QualificacaoOutput
}