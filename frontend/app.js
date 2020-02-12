const express = require('express')
const indexRoute = require('./routes')
const bodyParser = require('body-parser')
const app = express()

app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: false }))
app.use('/', indexRoute)

app.listen(3000, function () {
  console.log('listening port 3000')
})

module.exports = app
