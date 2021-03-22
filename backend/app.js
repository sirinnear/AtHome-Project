const mongoose = require('mongoose');
mongoose.connect('mongodb+srv://Anonymous643:a643295@cluster0.q2aa4.mongodb.net/Athome?retryWrites=true&w=majority')
    .then(() => {
        console.log('Successfully connected to MongoDB Atlas!');
    })
    .catch((error) => {
        console.log('Unable to connect to MongoDB Atlas!');
        console.error(error);
    });
const express = require('express');
const cors = require('cors');
var bodyParser = require('body-parser');
const app = express();
module.exports = app;
app.use(cors());
app.use(bodyParser.json());
const userRoutes = require('./routes/user');
app.use('/api/auth', userRoutes);
//const stuffRoutes = require('./routes/stuff');
//app.use('/api/stuff', stuffRoutes);