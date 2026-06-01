import 'package:flutter/material.dart';
import '../../controllers/audio_controller.dart';
import '../../controllers/room_controller.dart';
import '../meeting/tap_to_talk_button.dart';

class MeetingRoomScreen extends StatefulWidget {
  final String roomID;
  final String userID;

  const MeetingRoomScreen({
    super.key,
    required this.roomID,
    required this.userID,
  });

  @override
  State<MeetingRoomScreen> createState() => _MeetingRoomScreenState();
}

class _MeetingRoomScreenState extends State<MeetingRoomScreen> {
  late RoomController _roomController;
  late AudioController _audioController;

  @override
  void initState() {
    super.initState();
    _roomController = RoomController(
      roomID: widget.roomID,
      userId: widget.userID,
    );
    _audioController = AudioController();
    _initAudio();
    _roomController.addListener(_onStateChanged);
  }

  Future<void> _initAudio() async {
    await _audioController.initialize();
  }

  void _onStateChanged() {
    setState(() {});
  }

  @override
  void dispose() {
    _roomController.removeListener(_onStateChanged);
    _roomController.dispose();
    _audioController.destroy();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final state = _roomController.state;
    final isMyTurn = state?.activeSpeaker == widget.userID;

    return Scaffold(
      appBar: AppBar(title: Text('Room: ${widget.roomID}')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(isMyTurn ? 'Your Turn' : 'Waiting...'),
            const SizedBox(height: 32),
            TapToTalkButton(
              audioController: _audioController,
              isMyTurn: isMyTurn,
            ),
            const SizedBox(height: 16),
            if (!isMyTurn)
              ElevatedButton(
                onPressed: () => _roomController.requestBall(),
                child: const Text('Request Ball'),
              ),
          ],
        ),
      ),
    );
  }
}
