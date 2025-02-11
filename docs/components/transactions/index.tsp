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

model TransactionsMeta {
    totalCount?: integer;
}

model GetTransactionsResponse {
    data: Transaction[];
    links?: Links;
    meta?: TransactionsMeta;
}

@error
model TransactionNotFoundError {
    code: "TRANSACTION_NOT_FOUND";
    message: string;
}

@route("/transactions")
namespace Transactions {
    @get
    op getTransactions(
        @query `filter[from]`?: string,
        @query `filter[to]`?: string,
        @query `filter[block_number]`?: integer,
        @query `filter[timestamp]`?: integer,
        @query `page[number]`?: integer,
        @query `page[size]`?: integer
    ): {
        @statusCode statusCode: 200;
        @body response: GetTransactionsResponse;
    } | {
        @statusCode statusCode: 500;
        error: InternalServerError;
    };

    @get
    op getTransaction(@path transactionHash: string): {
        @statusCode statusCode: 200;
        @body transaction: Transaction;
    } | {
        @statusCode statusCode: 404;
        error: TransactionNotFoundError;
    };
}
