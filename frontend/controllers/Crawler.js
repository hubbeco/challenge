const got = require('got')

class Crawler {
  constructor (req) {
    this.req = req
  }

  async findHtml () {
    try {
      const response = await got(`http://${this.req.params.address}`)
      return response.body
    } catch (error) {
      return error
    }
  }
}

module.exports = Crawler
