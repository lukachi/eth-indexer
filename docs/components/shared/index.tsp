import "@typespec/http";
using TypeSpec.Http;

namespace EthIndexer;

model Links {
    first?: string;
    last?: string;
    prev?: string;
    next?: string;
}

model QueryPage {
    number?: integer;
    size?: integer;
}
