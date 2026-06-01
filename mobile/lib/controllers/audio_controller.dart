import 'dart:ffi';
import 'dart:io';
import 'dart:typed_data';
import 'package:ffi/ffi.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:lazysodium/lazysodium.dart';
import 'audio_engine_bindings.dart';

class AudioController {
  late final AudioEngineBindings _bindings;
  bool _isInitialized = false;

  // Symmetric key for encryption - In production, this should be securely managed
  static final Uint8List _key = Uint8List.fromList(
    List.generate(32, (index) => index),
  );
  static final Lazysodium _sodium = Lazysodium.instance();

  AudioController() {
    final dylib = Platform.isAndroid
        ? DynamicLibrary.open('libaudio_engine.so')
        : DynamicLibrary.process();
    _bindings = AudioEngineBindings(dylib);
  }

  // Handle incoming WebSocket messages
  void handleMessage(String type) {
    if (type == 'FORCE_MUTE') {
      print('Received FORCE_MUTE signal');
      stopCapture();
    }
  }

  // Callback for native layer - must be static
  static void _onChunkReady(Pointer<Int8> filenamePtr) {
    final filename = filenamePtr.cast<Utf8>().toDartString();
    print('Chunk ready for encryption: $filename');
    _processChunk(filename);
  }

  static Future<void> _processChunk(String filename) async {
    try {
      final file = File(filename);
      if (!await file.exists()) return;

      final rawData = await file.readAsBytes();

      // Encrypt
      final encryptedData = _sodium.cryptoSecretBoxEasy(rawData, _key);

      // Write encrypted file and delete raw
      final encryptedFile = File('$filename.enc');
      await encryptedFile.writeAsBytes(encryptedData);
      await file.delete();

      print('Encrypted chunk saved: ${encryptedFile.path}');
    } catch (e) {
      print('Encryption error: $e');
    }
  }

  Future<void> initialize() async {
    if (_isInitialized) return;

    // Ensure Sodium is ready
    await _sodium.ready;

    // Request microphone permission
    final status = await Permission.microphone.request();
    if (status != PermissionStatus.granted) {
      print('Microphone permission denied');
      return;
    }

    // Get storage path for chunks
    final directory = await getApplicationDocumentsDirectory();
    final storagePath = directory.path;

    // Create a pointer to the static Dart callback
    final callback = Pointer.fromFunction<Void Function(Pointer<Int8>)>(
      _onChunkReady,
    );

    // Pass path and callback to native engine
    final pathPtr = storagePath.toNativeUtf8();
    _bindings.init_audio_engine(pathPtr.cast<Char>(), callback);
    malloc.free(pathPtr);

    _isInitialized = true;
    print('Audio Engine Initialized at: $storagePath');
  }

  void startCapture() {
    if (_isInitialized) {
      _bindings.start_capture();
    }
  }

  void stopCapture() {
    if (_isInitialized) {
      _bindings.stop_capture();
    }
  }

  void destroy() {
    if (_isInitialized) {
      _bindings.destroy_audio_engine();
      _isInitialized = false;
    }
  }
}
