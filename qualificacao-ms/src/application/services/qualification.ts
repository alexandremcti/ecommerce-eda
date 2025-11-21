import STATE from "../../main/shared/data/state";
import { QualificacaoServicePort, QualificacaoServicePortInputDTO } from "@ports/in/qualificacao";
import { QualificacaoOutput } from "../../ports/out/qualificacao";


export class QualificacaoService implements QualificacaoServicePort {
    constructor(private qualificacaoOutput: QualificacaoOutput) {}

    async create(input: QualificacaoServicePortInputDTO): Promise<void> {
        
        const isOrderQualified = STATE.find(({acronym}) => input.customer.deliveryAddress.state)
        if(!isOrderQualified) {
            await this.qualificacaoOutput.OrderRecused(input)
            return
        }
        await this.qualificacaoOutput.OrderQualified(input)
    }
}