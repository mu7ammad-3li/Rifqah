import 'package:flutter/material.dart';

class RifqahColors {
  static const Color surface = Color(0xFFFBF9F5);
  static const Color surfaceContainerLow = Color(0xFFF5F3EF);
  static const Color surfaceContainerHigh = Color(0xFFEAE8E4);
  static const Color surfaceContainerHighest = Color(0xFFE4E2DE);
  static const Color onSurface = Color(0xFF1B1C1A);
  static const Color onSurfaceVariant = Color(0xFF414848);

  static const Color primary = Color(0xFF062425);
  static const Color onPrimary = Color(0xFFFFFFFF);
  static const Color primaryContainer = Color(0xFF1E3A3A);

  static const Color secondary = Color(0xFF7B563F);
  static const Color onSecondary = Color(0xFFFFFFFF);
  static const Color secondaryContainer = Color(0xFECDB0);
  static const Color onSecondaryContainer = Color(0xFF79553E);

  static const Color tertiary = Color(0xFF24200F);
  static const Color tertiaryFixedDim = Color(0xFFCEC6AD);

  static const Color outline = Color(0xFF717878);
  static const Color outlineVariant = Color(0xFFC1C8C7);
}

class RifqahSpacing {
  static const double base = 8.0;
  static const double stackGap = 16.0;
  static const double sectionGap = 48.0;
  static const double containerPaddingMobile = 24.0;
}

class RifqahShapes {
  static const BorderRadius sm = BorderRadius.all(Radius.circular(8.0));
  static const BorderRadius md = BorderRadius.all(Radius.circular(24.0));
  static const BorderRadius lg = BorderRadius.all(Radius.circular(48.0));
  static const BorderRadius full = BorderRadius.all(Radius.circular(9999.0));
}

class AppTheme {
  static ThemeData get lightTheme {
    return ThemeData(
      useMaterial3: true,
      scaffoldBackgroundColor: RifqahColors.surface,
      colorScheme: const ColorScheme.light(
        primary: RifqahColors.primary,
        onPrimary: RifqahColors.onPrimary,
        primaryContainer: RifqahColors.primaryContainer,
        secondary: RifqahColors.secondary,
        onSecondary: RifqahColors.onSecondary,
        secondaryContainer: RifqahColors.secondaryContainer,
        onSecondaryContainer: RifqahColors.onSecondaryContainer,
        surface: RifqahColors.surface,
        onSurface: RifqahColors.onSurface,
        outline: RifqahColors.outline,
      ),
      textTheme: const TextTheme(
        headlineLarge: TextStyle(
          fontFamily: 'Plus Jakarta Sans',
          fontSize: 40,
          fontWeight: FontWeight.w700,
          letterSpacing: -0.02 * 40,
          color: RifqahColors.primary,
        ),
        headlineMedium: TextStyle(
          fontFamily: 'Plus Jakarta Sans',
          fontSize: 28,
          fontWeight: FontWeight.w600,
          letterSpacing: -0.01 * 28,
          color: RifqahColors.primary,
        ),
        bodyLarge: TextStyle(
          fontFamily: 'Plus Jakarta Sans',
          fontSize: 18,
          fontWeight: FontWeight.w400,
          color: RifqahColors.onSurface,
        ),
        bodyMedium: TextStyle(
          fontFamily: 'Plus Jakarta Sans',
          fontSize: 16,
          fontWeight: FontWeight.w400,
          color: RifqahColors.onSurface,
        ),
        labelMedium: TextStyle(
          fontFamily: 'Plus Jakarta Sans',
          fontSize: 14,
          fontWeight: FontWeight.w600,
          letterSpacing: 0.05 * 14,
          color: RifqahColors.secondary,
        ),
        labelSmall: TextStyle(
          fontFamily: 'Plus Jakarta Sans',
          fontSize: 12,
          fontWeight: FontWeight.w500,
          letterSpacing: 0.03 * 12,
          color: RifqahColors.onSurfaceVariant,
        ),
      ),
    );
  }
}
