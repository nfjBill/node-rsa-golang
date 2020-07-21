const NodeRSA = require('node-rsa')

const key = new NodeRSA({b: 1024});
const pair = key.generateKeyPair();
const publicKey = pair.exportKey('pkcs1-public-pem')
const privateKey = pair.exportKey('pkcs1-private-pem')

console.log(`publicKey:${publicKey}`)
console.log(`privateKey:${privateKey}`)

const pubKey = new NodeRSA(publicKey);

const data = `adfadfasdfasdfdasfasdffadfasdf中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文`

const encData = pubKey.encrypt(data, 'base64');

console.log(`cipherData:${encData}`)


const priKey = new NodeRSA(privateKey);
const plainTxt = priKey.decrypt(encData, 'utf-8')

console.log(`plainTxt:${plainTxt}`)