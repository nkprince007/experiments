const enableDevTools = window.__GRPCWEB_DEVTOOLS__ || (() => {});
const { Request, Response } = require("./proto/echo_pb")
const { EchoClient } = require("./proto/echo_grpc_web_pb")
var client = new EchoClient('http://localhost:8080')
enableDevTools([client,])

var request = new Request()
request.setMsg("Hello World!")

client.say(request, {}, (err, response) => {
  console.log(response.getMsg())
})
