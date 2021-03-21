const Cert = require('../models/certificates');
exports.createCert = (req, res, next) =>{
    const cert = new Cert({
        name: req.body.name,
        cert: req.body.cert,
        verify: req.body.verify
    });
    cert.save().then(
        () => {
            res.status(201).json({
                message: 'Post saved successfully!'
            });
        }
    ).catch(
        (error) => {
            res.status(400).json({
                error: error
            });
        }
    );
};
exports.verifyCert = (req, res, next) =>{ //change status for querying
    Cert.findOne({
        verify: "waiting"
    }).then(
        Cert.certificates.update(
            { },
            { $set: { verify:verified} }
        )
    ).catch(
        (error) => {
            res.status(404).json({
                error: error
            });
        }
    );
}
exports.deleteCert = (req, res, next) =>{ //delete
    Cert.deleteOne({name: req.params.name}).then(
        () => {
            res.status(200).json({
                message: 'Deleted!'
            });
        }
    ).catch(
        (error) => {
            res.status(400).json({
                error: error
            });
        }
    );
}
exports.getCertByName = (req, res, next) =>{ //get cert by name
    Cert.find({
        name: req.body.name
    }).then(
        (cert) => {
            res.status(200).json(cert);
        }
    ).catch(
        (error) => {
            res.status(400).json({
                error: error
            });
        }
    );
}
exports.getCertByStatus = (req, res, next) =>{ //list all that is waiting
    Cert.find({
        verify: "waiting"
    }).then(
        (cert) => {
            res.status(200).json(cert);
        }
    ).catch(
        (error) => {
            res.status(400).json({
                error: error
            });
        }
    );
}