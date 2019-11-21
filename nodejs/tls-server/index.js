const {Server, ServerCredentials} = require("grpc");
const fs = require("fs");
const asyncFs = fs.promises;

async function main() {
    const tlsKey = "/path/to/tls.key";
    const tlsCert = "/path/to/tls.crt";
    const tlsCa = "/path/to/ca.crt";

    const server = new Server();

    const [ key, cert, ca ] = await Promise.all([
        asyncFs.readFile(tlsKey),
        asyncFs.readFile(tlsCert),
        asyncFs.readFile(tlsCa),
    ]);

    const credentials = ServerCredentials.createSsl(ca, [ {
        private_key: key,
        cert_chain: cert,
    } ], true);

    // server.addService(MyService.service, implementation);

    server.bind("0.0.0.0:8443", credentials);
    server.start();
}

main().catch((e) => console.error(e));
