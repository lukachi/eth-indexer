import "reflect-metadata"

import express from "express";
import ViteExpress from "vite-express";
import BlocksController from "./block";
import TransactionsController from "./transaction";
import {generateDocument} from "openapi-metadata";
import {generateScalarUI} from "openapi-metadata/ui";

const app = express();

const bootstrap = async () => {
  const document = await generateDocument({
    controllers: [BlocksController, TransactionsController],
    document: {
      info: {
        title: "My API",
        version: "1.0.0",
      },
    },
  });

  app.get("/hello", (_, res) => {
    res.send("Hello Vite + TypeScript!");
  });

  app.get("/api", async (_, res) => {
    res.send(JSON.stringify(document));
  })

  app.get("/api/docs", (_, res) => {
    const ui = generateScalarUI("/api")

    res.send(ui)
  })

  ViteExpress.listen(app, 3000, () =>
      console.log("Server is listening on port 3000..."),
  );
}

bootstrap()
