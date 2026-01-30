// ignore: unused_import
import 'package:intl/intl.dart' as intl;
import 'app_localizations.dart';

// ignore_for_file: type=lint

/// The translations for Russian (`ru`).
class AppLocalizationsRu extends AppLocalizations {
  AppLocalizationsRu([String locale = 'ru']) : super(locale);

  @override
  String get appTitle => 'Клиент НеоБанка';

  @override
  String get authLoginLabel => 'Логин';

  @override
  String get authPasswordLabel => 'Пароль';

  @override
  String get authButton => 'Войти';

  @override
  String homeGreeting(String name) {
    return 'Привет, $name';
  }

  @override
  String get totalBalanceTitle => 'Общий баланс';

  @override
  String get sectionAccounts => 'Мои счета';

  @override
  String get sectionCards => 'Карты';

  @override
  String get actionTopUp => 'Пополнить';

  @override
  String get actionTransfer => 'Перевести';

  @override
  String get actionHistory => 'История';

  @override
  String get actionSettings => 'Настройки';

  @override
  String get tabHome => 'Главная';

  @override
  String get tabPayments => 'Платежи';

  @override
  String get tabAtm => 'Банкоматы';

  @override
  String get tabProfile => 'Профиль';

  @override
  String get statusActive => 'Активен';

  @override
  String get statusBlocked => 'Заблокирован';

  @override
  String get settingsLanguage => 'Язык приложения';

  @override
  String get settingsTheme => 'Темная тема';
}
