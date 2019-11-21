import {Server, ServerCredentials} from "grpc";

async function main() {
    const server = new Server();

    // server.addService(MyService.service, implementation);

    server.bind("0.0.0.0:8080", ServerCredentials.createInsecure());
    server.start();
}

main().catch((e) => console.error(e));
