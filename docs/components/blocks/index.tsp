import "../shared/index.tsp";

import "@typespec/http";
using TypeSpec.Http;

namespace EthIndexer;

model BlockAttributes {
    number: integer;
    hash: string;
    parentHash: string;
    timestamp: integer;
}

model Block {
    id: string;
    type: string;
    attributes: BlockAttributes;
}

model BlocksMeta {
    totalCount?: integer;
}

model GetBlocksResponse {
    data: Block[];
    links?: Links;
    meta?: BlocksMeta;
}

@error
model BlockNotFoundError {
    code: "BLOCK_NOT_FOUND";
    message: string;
}

@error
model InternalServerError {
    code: "INTERNAL_SERVER_ERROR";
    message: string;
}

@route("/blocks")
namespace Blocks {
    @get
    op getBlocks(
        @query `filter[number]`?: integer,
        @query `filter[hash]`?: string,
        @query `filter[timestamp]`?: integer,
        @query `page[number]`?: integer,
        @query `page[size]`?: integer
        // @query filter?: {
            // number?: integer;
            // hash?: string;
            // timestamp?: integer;
        // },
        // @query page?: QueryPage;
    ): {
        @statusCode statusCode: 200;
        @body response: GetBlocksResponse;
    } | {
        @statusCode statusCode: 500;
        error: InternalServerError;
    };

    @get
    op getBlock(@path blockId: string): {
        @statusCode statusCode: 200;
        @body block: Block;
    } | {
        @statusCode statusCode: 404;
        error: BlockNotFoundError;
    };
}