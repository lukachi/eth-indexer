import {ApiOperation, ApiParam, ApiProperty, ApiResponse} from "openapi-metadata/decorators";

class TransactionAttributes {
    @ApiProperty({ type: "string" })
    declare hash: string
    @ApiProperty({ type: "string" })
    declare from: string
    @ApiProperty({ type: "string" })
    declare to: string
    @ApiProperty({ type: "string" })
    declare value: string
    @ApiProperty({ type: "number" })
    declare blockNumber: string
    @ApiProperty({ type: "number" })
    declare timestamp: number
}

class Transaction {
    @ApiProperty({ type: "string", example: "1234567" })
    declare id: string
    @ApiProperty({ type: "string" })
    declare type: 'transaction'
    @ApiProperty({ type: TransactionAttributes })
    declare attributes: TransactionAttributes
}

export default class TransactionsController {
    @ApiOperation({
        methods: ["get"],
        path: "/transaction",
        summary: ""
    })
    @ApiParam({
        name: "hash",
        in: "query"
    })
    @ApiResponse({ type: [Transaction] })
    async getTransaction() {}
}