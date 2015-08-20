/* global CryptoJS */

var hashcash = function (callback) {

  var base = '060408-23:userName@soulog:'
  var rand = function () {
    return Math.random().toString(36).slice(-16)
  }
  var rand2 = function () {
    var chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var randomstring = []
    for (var i = 0; i < 16; i++) {
      randomstring.push(chars.substr(Math.floor(Math.random() * chars.length), 1))
    }

    return randomstring.join('')
  }

  var endCall = function (hash) {
    callback(hash)
  }
  var repeat = function () {
    var num = 3000
    setTimeout(function () {
      var findIt = false
      for (var i = 0; i < num; i++) {
        var s = CryptoJS.SHA1(base + rand2()).toString()
        if (s.charAt(0) === '0' &&
            s.charAt(1) === '0' &&
            s.charAt(2) === '0' &&
            s.charAt(3) === '0' &&
            s.charAt(4) === '0') {
          endCall(s)
          findIt = true
          break
        }
      }
      if (!findIt) repeat()
    }, 0)
  }
  repeat()

}
