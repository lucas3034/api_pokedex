const cors = require('cors')
const express = require('express')
const app = express()
const axios = require('axios')

app.use(cors())

app.get('/', (req, res) => {
    return res.json([
        { mensagem: 'Teste: on'}
])
})

app.listen('7000')