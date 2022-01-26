const cors = require('cors')
const express = require('express')
const app = express()

app.use(cors())

app.get('/', (req, res) => {
    return res.json({ message: 'Teste Front: ok'})
})

app.listen('7000')