const express = require('express')
const addressRoute = require('./addressRoute')

const router = express.Router()

router.use('/url', addressRoute)

module.exports = router
