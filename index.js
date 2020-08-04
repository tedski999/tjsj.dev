const fs = require("fs");
const express = require("express");
const https = require("https");
const sslCertDirectory = "sslcert";
const port = 443;

// Load SSL Credentials
const credentials = {
	key: fs.readFileSync(sslCertDirectory + "/privkey.pem", "utf8"),
	cert: fs.readFileSync(sslCertDirectory + "/fullchain.pem", "utf8")
};

// Server setup
let app = express();
app.set("view engine", "ejs");
app.get("/", (req, res) => res.render("pages/index"));
app.use((req, res, next) => res.status(404).render("pages/404"));
https.createServer(credentials, app).listen(port, () => console.log(`HTTPS server listening on port ${port}`));
