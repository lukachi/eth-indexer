import "@typespec/http";
using TypeSpec.Http;

namespace EthIndexer;

model Block {
    id: string;
    type: string;
    attributes: {
        number: integer;
        hash: string;
        parentHash: string;
        timestamp: integer;
    }
}

@error
model BlockNotFoundError {
    code: "BLOCK_NOT_FOUND";
    message: string;
}

@route("/blocks")
namespace Blocks {
    @get
    op getBlock(@path blockId: string): {
        @statusCode statusCode: 200;
        @body block: Block;
    } | {
        @statusCode statusCode: 404;
        error: BlockNotFoundError;
    };
}