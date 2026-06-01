import 'dart:convert';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:flutter/foundation.dart';
import '../../config/env_config.dart';

class RoomState {
  final String activeSpeaker;
  final List<String> queue;

  RoomState({required this.activeSpeaker, required this.queue});

  factory RoomState.fromJson(Map<String, dynamic> json) {
    return RoomState(
      activeSpeaker: json['active'] ?? '',
      queue: List<String>.from(json['queue'] ?? []),
    );
  }
}

class PollState {
  final String targetUserID;
  final int roundIndex;
  PollState({required this.targetUserID, required this.roundIndex});
}

class RoomController extends ChangeNotifier {
  late WebSocketChannel _channel;
  RoomState? _state;
  PollState? _currentPoll;
  bool _isGracePeriod = false;
  final String userId;

  RoomController({required String roomID, required this.userId}) {
    final url = '${EnvConfig.wsUrl}/$roomID?userID=$userId';
    _channel = WebSocketChannel.connect(Uri.parse(url));
    _listen();
  }

  RoomState? get state => _state;
  PollState? get currentPoll => _currentPoll;
  bool get isGracePeriod => _isGracePeriod;

  void _listen() {
    _channel.stream.listen((message) {
      final jsonMsg = jsonDecode(message);
      if (jsonMsg['type'] == 'BALL_STATE') {
        _state = RoomState.fromJson(jsonMsg);
        notifyListeners();
      } else if (jsonMsg['type'] == 'REPORT_POLL') {
        _currentPoll = PollState(
          targetUserID: jsonMsg['target_user_id'],
          roundIndex: jsonMsg['round_index'],
        );
        notifyListeners();
      } else if (jsonMsg['type'] == 'GRACE_PERIOD_START') {
        _isGracePeriod = true;
        notifyListeners();
      }
    });
  }

  void submitVote(String targetUserID, int roundIndex, bool vote) {
    _channel.sink.add(
      jsonEncode({
        'type': 'SUBMIT_VOTE',
        'payload': {
          'target_user_id': targetUserID,
          'round_index': roundIndex,
          'vote': vote,
        },
      }),
    );
    _currentPoll = null;
    notifyListeners();
  }

  void requestBall() {
    _channel.sink.add(jsonEncode({'type': 'REQUEST_BALL'}));
  }

  void reportSegment(String targetUserID, int roundIndex) {
    _channel.sink.add(
      jsonEncode({
        'type': 'REPORT_SEGMENT',
        'payload': {'target_user_id': targetUserID, 'round_index': roundIndex},
      }),
    );
  }

  void dispose() {
    _channel.sink.close();
  }
}
