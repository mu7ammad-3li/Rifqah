import 'dart:convert';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:flutter/foundation.dart';

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

class RoomController extends ChangeNotifier {
  late WebSocketChannel _channel;
  RoomState? _state;
  final String userId;

  RoomController({required String roomID, required this.userId}) {
    // Note: URL needs to be configured based on env
    final url = 'ws://localhost:8080/ws/$roomID?userID=$userId';
    _channel = WebSocketChannel.connect(Uri.parse(url));
    _listen();
  }

  RoomState? get state => _state;

  void _listen() {
    _channel.stream.listen((message) {
      final jsonMsg = jsonDecode(message);
      if (jsonMsg['type'] == 'BALL_STATE') {
        _state = RoomState.fromJson(jsonMsg);
        notifyListeners();
      }
    });
  }

  void requestBall() {
    _channel.sink.add(jsonEncode({'type': 'REQUEST_BALL'}));
  }

  void dispose() {
    _channel.sink.close();
  }
}
