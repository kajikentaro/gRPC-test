// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import { grpc } from "@improbable-eng/grpc-web";
import type { NextApiRequest, NextApiResponse } from "next";
import { HelloReply, HelloRequest } from "../../grpc_out/grpc_pb";
import { Greeter } from "../../grpc_out/grpc_pb_service";

type Data = {
  name: string;
};

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  const queryBooksRequest = new HelloRequest();
  queryBooksRequest.setName("Geor");
  grpc.invoke(Greeter.SayHello, {
    request: queryBooksRequest,
    host: "https://example.com",
    onMessage: (message: HelloReply) => {
      console.log("got book: ", message.toObject());
    },
    onEnd: (
      code: grpc.Code,
      msg: string | undefined,
      trailers: grpc.Metadata
    ) => {
      if (code == grpc.Code.OK) {
        console.log("all ok");
      } else {
        console.log("hit an error", code, msg, trailers);
      }
    },
  });

  res.status(200).json({ name: "John Doe" });
}
