var http=require("http"),
    fs = require("fs");

var postData = fs.readFileSync("postData.xml")


var req = http.request({
    hostname:"172.7.1.39",
    port: 80,
    method:"POST",
    path: "/kjdec/upload.asmx",
    headers: {
        "Content-Type" : "application/soap+xml; charset=utf-8",
        "Content-Length": postData.length
    },
},function(res){
    console.log('STATUS: ' + res.statusCode);
    console.log('HEADERS: ' + JSON.stringify(res.headers));
    res.setEncoding('utf-8');
    res.on('data',function(chunk){
        console.log("BODY: " + chunk);
    });
});

req.on("error", function(err){
    console.log("Problem with request: " + err.message);
});

req.write(postData);
req.end();
console.log("message sent");
