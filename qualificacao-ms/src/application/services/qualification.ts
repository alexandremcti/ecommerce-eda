import STATE from "../../main/shared/data/state";
import { QualificacaoServicePort, QualificacaoServicePortInputDTO } from "@ports/in/qualificacao";
import { QualificacaoOutput } from "../../ports/out/qualificacao";


export class QualificacaoService implements QualificacaoServicePort {
    constructor(private qualificacaoOutput: QualificacaoOutput) {}

    async create(input: QualificacaoServicePortInputDTO): Promise<void> {
        
        const isOrderQualified = STATE.find(({acronym}) => input.customer.deliveryAddress.state)
        if(!isOrderQualified) {
            const order = {...input, status: 'recusado'}
            await this.qualificacaoOutput.OrderRecused(order)
            return
        }
        const order = {...input, status: 'qualificado'}
        await this.qualificacaoOutput.OrderQualified(order)
    }
}