import STATE from "../../main/shared/data/state";
import { QualificacaoServicePort, QualificacaoServicePortInputDTO } from "@ports/in/qualificacao";
import { QualificacaoOutput } from "../../ports/out/qualificacao";
import logger from "@shared/logs";


export class QualificacaoService implements QualificacaoServicePort {
    constructor(private qualificacaoOutput: QualificacaoOutput) {}

    async create(input: QualificacaoServicePortInputDTO): Promise<void> {
        
        const isOrderQualified = STATE.find(({acronym}) => input.customer.deliveryAddress.state === acronym)
        if(!isOrderQualified) {
            const order = {...input, status: 'recusado'}
            await this.qualificacaoOutput.OrderRecused(order)
            logger.info(`Order ${input.id} refused. State ${input.customer.deliveryAddress.state} is not qualified.`)
            return
        }
        const order = {...input, status: 'qualificado'}
        await this.qualificacaoOutput.OrderQualified(order)
        logger.info(`Order ${input.id} qualified successfully.`)
    }
}