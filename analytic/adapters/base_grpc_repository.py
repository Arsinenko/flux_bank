import grpc


class BaseGrpcRepository:
    def __init__(self, channel: grpc.aio.Channel):
        self.channel = channel

    async def close(self):
        await self.channel.close()

    @staticmethod
    async def _execute(coro):
        try:
            return await coro
        except grpc.aio.AioRpcError as ex:
            print(f"grpc error: {ex.code()} - {ex.details()}")
            if ex.code() == grpc.StatusCode.NOT_FOUND:
                return None
            raise