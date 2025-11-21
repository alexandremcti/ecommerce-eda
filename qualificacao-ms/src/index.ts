import express from 'express'
import { Server } from './main/server'

const app = express()
app.use(express.json())
const port = 3000

Server
    .start()
    .then(() => {
        
        app.get('/', (req, res) => {
          res.send('Hello, World!')
        })
        
        app.listen(port, () => {
            console.log(`Server is running at http://localhost:${port}`)
        })
    })
    .catch((err) => {
        console.log(`falha ao iniciar servidor. Error: ${err}`)
    })



