abstract class ISettingsRepository {
  Future<void> setLanguage(String languageCode); // 'ru' or 'en'
  Future<String> getLanguage();
  Future<void> setTheme(bool isDark);
}
