import 'package:moblile_app/domain/models/transaction.dart';

abstract class ITransactionRepository {
  Future<List<Transaction>> getTransactions({int? accountId});
}
