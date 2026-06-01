import 'package:flutter/material.dart';
import 'components/shared/app_theme.dart';
import 'components/auth/initial_welcome_screen.dart';
import 'config/env_config.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await EnvConfig.load();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Rifqah',
      theme: AppTheme.lightTheme,
      home: const InitialWelcomeScreen(),
    );
  }
}
