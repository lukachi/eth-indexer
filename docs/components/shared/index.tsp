import "@typespec/http";
using TypeSpec.Http;

namespace EthIndexer;

model Links {
    self?: string;
    first?: string;
    last?: string;
    prev?: string;
    next?: string;
}

model QueryPage {
    number?: integer;
    size?: integer;
}
