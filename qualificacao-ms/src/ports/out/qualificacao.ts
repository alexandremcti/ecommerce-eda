import { QualificacaoServicePortInputDTO } from "../in/qualificacao"

export type OrderDTO = QualificacaoServicePortInputDTO

export interface QualificacaoOutput {
    OrderQualified(data: OrderDTO ): Promise<void>
    OrderRecused(data: OrderDTO): Promise<void>
}