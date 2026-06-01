import 'package:flutter_dotenv/flutter_dotenv.dart';

class EnvConfig {
  static Future<void> load() async {
    await dotenv.load(fileName: "assets/.env.prod");
  }

  static String get apiUrl => dotenv.get('API_URL');
  static String get apiProtocol => dotenv.get('API_PROTOCOL');
  static String get wsProtocol => dotenv.get('WS_PROTOCOL');

  static String get baseUrl => '$apiProtocol://$apiUrl';
  static String get wsUrl => '$wsProtocol://$apiUrl/ws';
}
