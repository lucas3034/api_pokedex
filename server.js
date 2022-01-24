const express = require('express')
const app = express()

app.get('/', (req, res) => {
    return res.json({ message: 'Teste Front: ok'})
})

app.listen('7000')