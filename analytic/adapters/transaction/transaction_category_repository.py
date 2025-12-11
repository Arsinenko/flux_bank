from typing import List

import grpc

from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.transaction_category_pb2 import *
from api.generated.transaction_category_pb2_grpc import TransactionCategoryServiceStub
from domain.transaction.transaction_category import TransactionCategory
from domain.transaction.transaction_category_repo import TransactionCategoryRepositoryAbc


class TransactionCategoryRepository(TransactionCategoryRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = TransactionCategoryServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

    @staticmethod
    def to_domain(model: TransactionCategoryModel) -> TransactionCategory:
        return TransactionCategory(
            category_id=model.category_id,
            name=model.name
        )

    @staticmethod
    def response_to_list(response: GetAllTransactionCategoriesResponse) -> List[TransactionCategory]:
        return [TransactionCategoryRepository.to_domain(model) for model in response.transaction_categories]

    async def get_all(self, page_n: int, page_size: int) -> List[TransactionCategory]:
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, category_id: int) -> TransactionCategory | None:
        try:
            result = await self.stub.GetById(GetTransactionCategoryByIdRequest(category_id=category_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None
