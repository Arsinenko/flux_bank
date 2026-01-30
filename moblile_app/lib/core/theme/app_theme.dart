
import 'package:flutter/material.dart';
import 'package:moblile_app/core/theme/app_colors.dart';

abstract class AppTheme {
  static final ThemeData lightTheme = ThemeData(
    brightness: Brightness.light,
    primaryColor: AppColors.lightPrimary,
    scaffoldBackgroundColor: AppColors.lightBackground,
    cardColor: AppColors.lightCardBackground,
    dividerColor: AppColors.lightDivider,
    textTheme: const TextTheme(
      bodyLarge: TextStyle(color: AppColors.lightPrimaryText),
      bodyMedium: TextStyle(color: AppColors.lightSecondaryText),
    ),
    colorScheme: const ColorScheme.light(
      primary: AppColors.lightPrimary,
      secondary: AppColors.lightPrimary,
      surface: AppColors.lightCardBackground,
      background: AppColors.lightBackground,
      error: AppColors.lightError,
      onPrimary: Colors.white,
      onSecondary: Colors.white,
      onSurface: AppColors.lightPrimaryText,
      onBackground: AppColors.lightPrimaryText,
      onError: Colors.white,
      brightness: Brightness.light,
    ),
  );

  static final ThemeData darkTheme = ThemeData(
    brightness: Brightness.dark,
    primaryColor: AppColors.darkPrimary,
    scaffoldBackgroundColor: AppColors.darkBackground,
    cardColor: AppColors.darkCardBackground,
    dividerColor: AppColors.darkDivider,
    textTheme: const TextTheme(
      bodyLarge: TextStyle(color: AppColors.darkPrimaryText),
      bodyMedium: TextStyle(color: AppColors.darkSecondaryText),
    ),
    colorScheme: const ColorScheme.dark(
      primary: AppColors.darkPrimary,
      secondary: AppColors.darkPrimary,
      surface: AppColors.darkCardBackground,
      background: AppColors.darkBackground,
      error: AppColors.darkError,
      onPrimary: Colors.black,
      onSecondary: Colors.black,
      onSurface: AppColors.darkPrimaryText,
      onBackground: AppColors.darkPrimaryText,
      onError: Colors.black,
      brightness: Brightness.dark,
    ),
  );
}
