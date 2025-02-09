import "@typespec/http";
import "@typespec/versioning";
import "@typespec/openapi3";
import "@typespec/rest";


import "./components/blocks/index.tsp";
import "./components/transactions/index.tsp";

using TypeSpec.Http;
using TypeSpec.Versioning;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;
@service({
    title: "Eth indexer svc",
    description: "Service for indexing eth blocks and transactions",
})
// @info({
//     contact: {
//         name: "API Support",
//         email: "contact@contoso.com",
//     },
// })
@server("http:localhost:8089", "Single server for eth indexer")
@versioned(Versions)

namespace EthIndexer;

enum Versions {
    v1: "0.1.0",
}