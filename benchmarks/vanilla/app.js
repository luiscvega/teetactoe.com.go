var express = require("express");
var app = express();

app.set("view engine", "swig");
app.engine("swig", require("swig").renderFile);
app.use(express.static(__dirname + "/public"));

app.get("/", function (req, res) {
    res.render("index", { message: "index" });
});

app.get("/signup", function (req, res) {
    res.render("signup", { message: "signup" });
});
 
app.listen(3000);
