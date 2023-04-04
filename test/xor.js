var fs = require("fs");

var key = '123'
  
  /**
   * @description 字符串异或加密，并做base64转码
   * 异或加密即对当前字符串每位与当前约定的key的每位进行异或操作，求出结果
   * charCodeAt: 返回当前字符的Unicode 编码
   * 异或：两个值相同返回1，两个值不同，返回0
   */
  const XORencryption = (val, key) => {
    if (typeof val !== 'string') return val;
    let message = '';
    for (var i = 0; i < val.length; i++) {
      message += String.fromCharCode(val.charCodeAt(i) ^ key.charCodeAt(i % key.length));
    }
    return message;
  };
  
  /**
   * @description 解密异或加密的密文
   */
  const decodeXOR = (val, key) => {
    if (typeof val !== 'string') return val;
    let message = '';
    for (var i = 0; i < val.length; i++) {
      message += String.fromCharCode(val.charCodeAt(i) ^ key.charCodeAt(i % key.length));
    }
    return message;
  };

var data = fs.readFileSync('./test/test.yml');
console.log("同步读取: " + data.toString());

const enc = XORencryption(data.toString(), key)
console.log('enc', enc)

fs.writeFile('input-enc.yml', enc,  function(err) {
    if (err) {
        return console.error(err);
    }
    console.log("数据写入成功！");
 });
 fs.writeFile('input-dec.yml', decodeXOR(enc, key),  function(err) {
    if (err) {
        return console.error(err);
    }
    console.log("数据写入成功！");
 });
