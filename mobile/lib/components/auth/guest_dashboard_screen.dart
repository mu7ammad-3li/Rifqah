import 'package:flutter/material.dart';

class GuestDashboardScreen extends StatelessWidget {
  const GuestDashboardScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Browse Groups')),
      body: ListView(
        children: const [
          ListTile(title: Text('Group A'), subtitle: Text('Locked')),
          ListTile(title: Text('Group B'), subtitle: Text('Locked')),
        ],
      ),
    );
  }
}
