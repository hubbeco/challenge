const express = require('express')
const ControllerCrawler = require('../controllers/Crawler')
const router = express.Router()

router.post('/:url', async function (req, res) {
  const result = await new ControllerCrawler(req).get()
  res.send(result)
})

module.exports = router
