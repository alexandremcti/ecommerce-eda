import { QualificacaoRepository } from "../../../ports/out/qualificacao.js";
import { constants } from "../../../main/config/constants.js";
import logger from "../../../main/shared/logs/index.js";
import { Database, MongoDBDatabase } from "./mongodb-database.js";
import { QualificacaoRepositoryImp } from "./qualificacao-repository.js";

export class DatabaseComponent {
    private _database: Database;
    private _qualificacaoRepository: QualificacaoRepository

    private constructor(params: DatabaseComponentProperties) {
        this._database = params.database
        this._qualificacaoRepository = params.qualificacaoRepository
    }

    get database(): Database {
        return this._database
    }

    get qualificacaoRepository(): QualificacaoRepository {
        return this._qualificacaoRepository
    }

    static async create(): Promise<DatabaseComponent> {
        logger.info("Connecting to database...")
        const database = await MongoDBDatabase.connect(constants.database.URL, {dbName: constants.database.DB_NAME})
        const qualificacaoRepository = new QualificacaoRepositoryImp(database)
        return new DatabaseComponent({database, qualificacaoRepository})
    }
}

type DatabaseComponentProperties = {
    database: Database;
    qualificacaoRepository: QualificacaoRepository
}