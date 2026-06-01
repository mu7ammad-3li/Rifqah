import 'package:flutter/material.dart';
import '../../components/shared/app_theme.dart';
import 'unauthenticated_discovery_screen.dart';

class InitialWelcomeScreen extends StatelessWidget {
  const InitialWelcomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.all(RifqahSpacing.containerPaddingMobile),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Icon(
                Icons.shield_outlined,
                size: 64,
                color: RifqahColors.primary,
              ),
              const SizedBox(height: RifqahSpacing.sectionGap),
              Text(
                'Welcome to Rifqah',
                style: Theme.of(context).textTheme.headlineLarge,
              ),
              const SizedBox(height: RifqahSpacing.stackGap),
              Text(
                'A private sanctuary for community connection.',
                style: Theme.of(context).textTheme.bodyLarge,
                textAlign: TextAlign.center,
              ),
              const SizedBox(height: RifqahSpacing.sectionGap),
              _buildButton(
                context,
                'Sign In',
                RifqahColors.primaryContainer,
                RifqahColors.onPrimary,
                () {},
              ),
              const SizedBox(height: RifqahSpacing.stackGap),
              _buildButton(
                context,
                'Sign Up',
                RifqahColors.secondaryContainer,
                RifqahColors.onSecondaryContainer,
                () {},
              ),
              const SizedBox(height: RifqahSpacing.stackGap),
              TextButton(
                onPressed: () {
                  Navigator.of(context).push(
                    MaterialPageRoute(
                      builder: (context) =>
                          const UnauthenticatedDiscoveryScreen(),
                    ),
                  );
                },
                child: Text(
                  'Discover Circles',
                  style: Theme.of(context).textTheme.labelMedium,
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildButton(
    BuildContext context,
    String text,
    Color bgColor,
    Color textColor,
    VoidCallback onPressed,
  ) {
    return SizedBox(
      width: double.infinity,
      height: 56,
      child: ElevatedButton(
        style: ElevatedButton.styleFrom(
          backgroundColor: bgColor,
          foregroundColor: textColor,
          shape: const RoundedRectangleBorder(
            borderRadius: BorderRadius.all(Radius.circular(28)),
          ),
        ),
        onPressed: onPressed,
        child: Text(
          text,
          style: Theme.of(
            context,
          ).textTheme.labelMedium?.copyWith(color: textColor),
        ),
      ),
    );
  }
}
