import { QualificacaoOutput } from "../../../ports/out/qualificacao.js";
import { MessageBroker } from "./rabbitmq-broker.js";

export class QualificacaoOutputImp implements QualificacaoOutput {
    constructor(private broker: MessageBroker) {}

    async QualificacaoCreated(data: any): Promise<void> {
        await this.broker.notify('qualificacao-created-out-0', data);
    }
}