import 'package:flutter/material.dart';
import 'package:uuid/uuid.dart';
import '../meeting/meeting_room_screen.dart';

class GuestDashboardScreen extends StatelessWidget {
  const GuestDashboardScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Browse Groups')),
      body: ListView(
        children: [
          ListTile(
            title: const Text('Group A'),
            subtitle: const Text('Tap to test connection'),
            onTap: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => MeetingRoomScreen(
                    roomID: 'testroom',
                    userID: const Uuid().v4(),
                  ),
                ),
              );
            },
          ),
        ],
      ),
    );
  }
}
