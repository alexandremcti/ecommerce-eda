import {createLogger, transports, format} from 'winston'
const {combine, timestamp, logstash} = format

const logstashFormat = combine(
    timestamp(),
    logstash()
)

const logger = createLogger({
    transports:[
        new transports.Console()
    ],
    format: logstashFormat
})


export default logger
