import { Collection, Db, MongoClient } from "mongodb"
import logger from "../../../main/shared/logs/index.js"

export interface Database {
    disconnect(): Promise<void>
    getRepository<T>(collectionName: string): T
}

export type MongoDBDatabaseInput = {
    client: MongoClient
    database: Db
}

export class MongoDBDatabase implements Database {

    private client: MongoClient
    private database: Db
    
    private constructor(private input: MongoDBDatabaseInput) {
        this.client = this.input.client
        this.database = this.input.database
    }

    static async connect(url: string, options: {dbName: string}): Promise<Database> {
        // Implementation for connecting to MongoDB
        const client = await new MongoClient(url).connect()
        const database = client.db(options.dbName)
        return new MongoDBDatabase({client, database})
    }

    async disconnect(): Promise<void> {
        // Implementation for disconnecting from MongoDB
        logger.info("Disconnecting from database...")
        await this.client.close()
    }

    getRepository<T = Collection<Document>>(collectionName: string): T {
        if (!this.database) {
            throw new Error("Database is not connected.")
        }
        return this.database.collection(collectionName) as T
    }
}

