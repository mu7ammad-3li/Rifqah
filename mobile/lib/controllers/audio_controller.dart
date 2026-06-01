import 'dart:ffi';
import 'dart:io';
import 'package:ffi/ffi.dart';
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

  void initialize() {
    if (!_isInitialized) {
      _bindings.init_audio_engine();
      _isInitialized = true;
    }
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
