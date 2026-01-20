from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.transaction_category_pb2 import *
from api.generated.transaction_category_pb2_grpc import TransactionCategoryServiceStub
from domain.transaction.transaction_category import TransactionCategory
from domain.transaction.transaction_category_repo import TransactionCategoryRepositoryAbc


class TransactionCategoryRepository(TransactionCategoryRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = TransactionCategoryServiceStub(channel=self.chanel)

    @staticmethod
    def to_domain(model: TransactionCategoryModel) -> TransactionCategory:
        return TransactionCategory(
            category_id=model.category_id,
            name=model.name
        )

    @staticmethod
    def to_model(domain: TransactionCategory) -> TransactionCategoryModel:
        return TransactionCategoryModel(
            category_id=domain.category_id,
            name=domain.name
        )

    @staticmethod
    def response_to_list(response: GetAllTransactionCategoriesResponse) -> List[TransactionCategory]:
        return [TransactionCategoryRepository.to_domain(model) for model in response.transaction_categories]

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[TransactionCategory]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, category_id: int) -> TransactionCategory | None:
        request = GetTransactionCategoryByIdRequest(category_id=category_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None
