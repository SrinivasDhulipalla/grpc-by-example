const grpc = require("grpc");

async function main() {
    const credentials = grpc.credentials.createInsecure();

    // const client = new pb.MyServiceClient("localhost:8080", credentials);
}

main().catch((e) => console.error(e));
