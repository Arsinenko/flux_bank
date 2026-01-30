import 'package:moblile_app/domain/models/account.dart';

abstract class IAccountRepository {
  Future<List<Account>> getAccounts();
  Future<double> getTotalBalance();
}
