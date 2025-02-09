import {ApiOperation, ApiProperty, ApiQuery, ApiResponse} from "openapi-metadata/decorators";

class BlockAttributes {
    @ApiProperty({type: "number", example: 123456})
    declare number: number
    @ApiProperty({type: "string", example: "0x....."})
    declare hash: string
    @ApiProperty({type: "string", example: "0x....."})
    declare parentHash: string
    @ApiProperty({type: "number", example: 1739051540})
    declare timestamp: number
}

class Block {
    @ApiProperty({type: "string", example: "1234567"})
    declare id: string
    @ApiProperty({type: "string", example: "block"})
    declare type: 'block'
    @ApiProperty({type: BlockAttributes})
    declare attributes: BlockAttributes
}

export default class BlocksController {
    @ApiOperation({
        methods: ["get"],
        path: "/block",
        summary: ""
    })
    @ApiQuery({
        name: 'number'
    })
    @ApiResponse({ type: Block })
    async getBlock() {
    }
}