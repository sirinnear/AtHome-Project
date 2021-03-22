const mongoose = require('mongoose');

const certificatesSchema = mongoose.Schema({
    name: {type:String, required:true},
    cert: {type:String, required:true},
    verify: {type:String, required:true}
});

module.exports = mongoose.model('Cert', certificatesSchema);