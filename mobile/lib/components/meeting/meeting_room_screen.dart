import 'package:flutter/material.dart';
import '../../controllers/audio_controller.dart';
import '../../controllers/room_controller.dart';
import '../meeting/tap_to_talk_button.dart';
import 'participant_orb.dart';
import '../shared/app_theme.dart';

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
  bool _isAnonymized = true;

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
    if (_roomController.currentPoll != null) {
      _showPollDialog(_roomController.currentPoll!);
    }
    setState(() {});
  }

  void _showPollDialog(PollState poll) {
    showDialog(
      context: context,
      barrierDismissible: false,
      builder: (context) => AlertDialog(
        title: const Text('Safety Report'),
        content: Text(
          'Participant ${poll.targetUserID} reported for segment ${poll.roundIndex}. Do you confirm the violation?',
        ),
        actions: [
          TextButton(
            onPressed: () {
              _roomController.submitVote(
                poll.targetUserID,
                poll.roundIndex,
                false,
              );
              Navigator.pop(context);
            },
            child: const Text('No'),
          ),
          ElevatedButton(
            onPressed: () {
              _roomController.submitVote(
                poll.targetUserID,
                poll.roundIndex,
                true,
              );
              Navigator.pop(context);
            },
            child: const Text('Yes'),
          ),
        ],
      ),
    );
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
    final activeSpeaker = state?.activeSpeaker ?? '';

    return Scaffold(
      backgroundColor: RifqahColors.surface,
      appBar: AppBar(
        backgroundColor: RifqahColors.surface,
        title: Text(
          'Room: ${widget.roomID}',
          style: Theme.of(context).textTheme.labelMedium,
        ),
        actions: [
          IconButton(
            icon: Icon(_isAnonymized ? Icons.visibility_off : Icons.visibility),
            onPressed: () {
              setState(() {
                _isAnonymized = !_isAnonymized;
              });
            },
          ),
        ],
      ),
      body: Stack(
        children: [
          // Background Atmosphere
          Center(
            child: ParticipantOrb(
              isActive: activeSpeaker == widget.userID,
              participantName: _isAnonymized ? 'Masked' : 'You',
              userID: widget.userID,
              roomController: _roomController,
              size: 200,
            ),
          ),
          // Controls
          Positioned(
            bottom: 40,
            left: 0,
            right: 0,
            child: Column(
              children: [
                TapToTalkButton(
                  audioController: _audioController,
                  isMyTurn: activeSpeaker == widget.userID,
                ),
                const SizedBox(height: 16),
                if (activeSpeaker != widget.userID)
                  ElevatedButton(
                    onPressed: () => _roomController.requestBall(),
                    style: ElevatedButton.styleFrom(
                      backgroundColor: RifqahColors.secondaryContainer,
                      shape: const RoundedRectangleBorder(
                        borderRadius: RifqahShapes.full,
                      ),
                    ),
                    child: const Text('Request Ball'),
                  ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
