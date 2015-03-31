var http = require("http"),
	fs = require("fs"),
	urlparse = require("url").parse,
	uuid = require("node-uuid")




http.createServer(function(req,res){
	var parsed = urlparse(req.url,true);
	var path = parsed.pathname;

	var components = path.split("/").filter(function(x){return x!=""});
	
	if (components.length < 2 ||
		components[0] != "grounding" ||
		components[1].length != 5
		) {
		res.writeHead("402");
		res.end("Bad request.");
		return;
	}

	var fileid = uuid.v4()

	var dest = fs.createWriteStream(__dirname+"/../data/"+fileid+".xml")
	req.pipe(dest);
	req.on("end",function(){
		res.writeHead(200);
		res.end("Message received");
	});

}).listen(9000);
console.log("Listening on port 9000");



