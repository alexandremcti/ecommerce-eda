export const constants = {
    database:  {
        URL:String(process.env.DATABASE_URL),
        DB_NAME: String(process.env.DB_NAME)
    },
    broker: {
        URL: String(process.env.BROKER_URL)
    }
}