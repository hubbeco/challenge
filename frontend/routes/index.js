const express = require('express')
const addressRoute = require('./addressRoute')

const router = express.Router()

router.use('/address', addressRoute)

module.exports = router
