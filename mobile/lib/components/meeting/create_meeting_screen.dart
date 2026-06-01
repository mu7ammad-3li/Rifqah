import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:uuid/uuid.dart';

class CreateMeetingScreen extends StatefulWidget {
  final String creatorID;
  const CreateMeetingScreen({super.key, required this.creatorID});

  @override
  State<CreateMeetingScreen> createState() => _CreateMeetingScreenState();
}

class _CreateMeetingScreenState extends State<CreateMeetingScreen> {
  final _titleController = TextEditingController();
  DateTime _selectedDate = DateTime.now().add(const Duration(hours: 1));

  Future<void> _createMeeting() async {
    final response = await http.post(
      Uri.parse('http://192.168.1.10:8080/meetings'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'title': _titleController.text,
        'creator_id': widget.creatorID,
        'meeting_type': 'standard',
        'scheduled_at': _selectedDate.toIso8601String(),
      }),
    );

    if (response.statusCode == 200) {
      if (mounted) Navigator.pop(context);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Schedule Meeting')),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            TextField(
              controller: _titleController,
              decoration: const InputDecoration(labelText: 'Title'),
            ),
            TextButton(
              onPressed: () async {
                final date = await showDatePicker(
                  context: context,
                  initialDate: _selectedDate,
                  firstDate: DateTime.now(),
                  lastDate: DateTime.now().add(const Duration(days: 30)),
                );
                if (date != null) setState(() => _selectedDate = date);
              },
              child: Text('Scheduled for: ${_selectedDate.toLocal()}'),
            ),
            ElevatedButton(
              onPressed: _createMeeting,
              child: const Text('Create'),
            ),
          ],
        ),
      ),
    );
  }
}
