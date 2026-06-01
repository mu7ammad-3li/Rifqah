import 'package:flutter/material.dart';

class AppColors {
  static const Color primary = Color(0xFF2D5A85);
  static const Color secondary = Color(0xFF4C8E99);
  static const Color actionActive = Color(0xFFD9534F);
  static const Color actionQueue = Color(0xFFF0AD4E);
  static const Color background = Color(0xFFF4F7F6);
  static const Color text = Color(0xFF333333);
}

class AppTheme {
  static ThemeData get lightTheme {
    return ThemeData(
      primaryColor: AppColors.primary,
      scaffoldBackgroundColor: AppColors.background,
      colorScheme: const ColorScheme.light(
        primary: AppColors.primary,
        secondary: AppColors.secondary,
        surface: AppColors.background,
      ),
      textTheme: const TextTheme(
        bodyMedium: TextStyle(color: AppColors.text),
      ),
      elevatedButtonTheme: ElevatedButtonThemeData(
        style: ElevatedButton.styleFrom(
          backgroundColor: AppColors.primary,
          foregroundColor: Colors.white,
        ),
      ),
    );
  }
}
