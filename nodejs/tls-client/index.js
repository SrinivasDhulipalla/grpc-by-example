const { readFile } = require("fs").promises;
const grpc = require("grpc");

async function main() {
    const tlsKey = "/path/to/tls.key";
    const tlsCert = "/path/to/tls.crt";
    const tlsCA = "/path/to/ca.crt";

    const [
        key,
        cert,
        ca,
    ] = await Promise.all([
        readFile(tlsKey),
        readFile(tlsCert),
        readFile(tlsCA),
    ]);

    const credentials = grpc.credentials.createSsl(ca, key, cert);

    // const client = new pb.MyServiceClient("localhost:8443", credentials)
}

main().catch((e) => console.error(e));
