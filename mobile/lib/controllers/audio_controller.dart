import 'dart:ffi';
import 'dart:io';
import 'package:ffi/ffi.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';
import 'audio_engine_bindings.dart';

class AudioController {
  late final AudioEngineBindings _bindings;
  bool _isInitialized = false;

  AudioController() {
    final dylib = Platform.isAndroid
        ? DynamicLibrary.open('libaudio_engine.so')
        : DynamicLibrary.process();
    _bindings = AudioEngineBindings(dylib);
  }

  Future<void> initialize() async {
    if (_isInitialized) return;

    // Request microphone permission
    final status = await Permission.microphone.request();
    if (status != PermissionStatus.granted) {
      print('Microphone permission denied');
      return;
    }

    // Get storage path for chunks
    final directory = await getApplicationDocumentsDirectory();
    final storagePath = directory.path;

    // Pass path to native engine
    final pathPtr = storagePath.toNativeUtf8();
    _bindings.init_audio_engine(pathPtr.cast<Char>());
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
