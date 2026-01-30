// ignore: unused_import
import 'package:intl/intl.dart' as intl;
import 'app_localizations.dart';

// ignore_for_file: type=lint

/// The translations for English (`en`).
class AppLocalizationsEn extends AppLocalizations {
  AppLocalizationsEn([String locale = 'en']) : super(locale);

  @override
  String get appTitle => 'NeoBank Client';

  @override
  String get authLoginLabel => 'Username';

  @override
  String get authPasswordLabel => 'Password';

  @override
  String get authButton => 'Sign In';

  @override
  String homeGreeting(String name) {
    return 'Hello, $name';
  }

  @override
  String get totalBalanceTitle => 'Total Balance';

  @override
  String get sectionAccounts => 'My Accounts';

  @override
  String get sectionCards => 'Cards';

  @override
  String get actionTopUp => 'Top Up';

  @override
  String get actionTransfer => 'Transfer';

  @override
  String get actionHistory => 'History';

  @override
  String get actionSettings => 'Settings';

  @override
  String get tabHome => 'Home';

  @override
  String get tabPayments => 'Payments';

  @override
  String get tabAtm => 'ATMs';

  @override
  String get tabProfile => 'Profile';

  @override
  String get statusActive => 'Active';

  @override
  String get statusBlocked => 'Blocked';

  @override
  String get settingsLanguage => 'App Language';

  @override
  String get settingsTheme => 'Dark Mode';
}
