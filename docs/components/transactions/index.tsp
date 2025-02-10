import "@typespec/http";
using TypeSpec.Http;

namespace EthIndexer;

model TransactionAttributes {
    blockNumber: integer;
    hash: string;
    from: string;
    to: string;
    value: string;
    timestamp: integer;
}

model Transaction {
    id: string;
    type: string;
    attributes: TransactionAttributes;
}

@route("/transactions")
namespace Transactions {
    @get
    op getTransaction(@path transactionId: string): {
        @statusCode statusCode: 200;
        @body transaction: Transaction;
    } | {
        @statusCode statusCode: 404;
        error: TransactionNotFoundError;
    };
}

@error
model TransactionNotFoundError {
    code: "TRANSACTION_NOT_FOUND";
    message: string;
}
