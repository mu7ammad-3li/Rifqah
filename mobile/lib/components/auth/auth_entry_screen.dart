import 'package:flutter/material.dart';
import 'guest_dashboard_screen.dart';

class AuthEntryScreen extends StatelessWidget {
  const AuthEntryScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Rifqah', style: Theme.of(context).textTheme.headlineLarge),
            const SizedBox(height: 48),
            ElevatedButton(
              onPressed: () {
                // TODO: Navigate to Sign In
              },
              child: const Text('Sign In'),
            ),
            const SizedBox(height: 16),
            TextButton(
              onPressed: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (context) => const GuestDashboardScreen()),
                );
              },
              child: const Text('Browse as Guest'),
            ),
          ],
        ),
      ),
    );
  }
}
