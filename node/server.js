var http = require("http"),
	fs = require("fs"),
	cluster = require("cluster"),
	urlparse = require("url").parse,
	uuid = require("node-uuid")

var numCPUs = require("os").cpus().length;


if(cluster.isMaster) {
	for(var i=0;i<numCPUs -1;i++) {
		cluster.fork();
	}

	cluster.on("exit",function(worker,code,signal){
		console.log("worker:",worker.process.pid,"is off");
	});

	cluster.on("fork",function(worker,code,signal){
		console.log("worker:", worker.process.pid, "is online");
	});

} else {
	http.createServer(function(req,res){
		var parsed = urlparse(req.url,true);
	 	var path = parsed.pathname;
		var components = path.split("/").filter(function(x){return x!=""});
		//console.log(components);
		if (components[0] != "grounding" ||
			components[1].length != 5 || 
			components.length > 2) {
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
}


