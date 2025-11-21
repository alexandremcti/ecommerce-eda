import { Collection } from "mongodb"
import { Database } from "./mongodb-database.js"
import { Qualificacao } from "../../../domain/qualificacao.js"
import { QualificacaoRepository } from "../../../ports/in/qualificacao.js"

export class QualificacaoRepositoryImp implements QualificacaoRepository {
    private collection: Collection<Document>

    constructor(database: Database) {
        this.collection = database.getRepository('Qualificacao')
    }

    async save(data: Qualificacao): Promise<void> {
        // Implementation for saving Qualificacao data to the database
        this.collection.insertOne({} as Document) // Placeholder implementation
    }
}