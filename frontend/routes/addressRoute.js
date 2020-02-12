const express = require('express')
const ControllerCrawler = require('../controllers/Crawler')
const router = express.Router()

router.post('/:address', async function (req, res) {
  const result = await new ControllerCrawler(req).findHtml()
  res.send(result)
})

module.exports = router
