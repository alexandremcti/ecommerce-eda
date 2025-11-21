import { BrokerComponent } from "../adapter/out/broker"
import { DatabaseComponent } from "../adapter/out/database"
import logger from "./shared/logs"

export class Server {

    private databaseComponent: DatabaseComponent
    private brokerComponent: BrokerComponent

    private constructor(input: ServerProperties){
        this.databaseComponent = input.databaseComponent
        this.brokerComponent = input.brokerComponent
    }

    static async start(): Promise<Server> {
        logger.info("Starting server...")
        const databaseComponent = await DatabaseComponent.create()
        logger.info("Database connected.")
        const brokerComponent = await BrokerComponent.create()
        logger.info("Message broker connected.")


        const server =  new Server({brokerComponent: brokerComponent, databaseComponent: databaseComponent})

        //Close resources on process termination
        process.on('SIGINT', async () => {
            logger.info("Shutting down server...")
            await Promise.all([
                server.databaseComponent.database.disconnect(),
                server.brokerComponent.broker.disconnect()
            ])
            process.exit(0)
        })

        return server
    }
}

export type ServerProperties = {
    databaseComponent: DatabaseComponent
    brokerComponent: BrokerComponent
}