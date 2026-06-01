#ifndef AUDIO_ENGINE_H
#define AUDIO_ENGINE_H

#ifdef __cplusplus
extern "C" {
#endif

// Callback type for notifying Dart when a chunk is ready
typedef void (*ChunkReadyCallback)(const char* filename);

// Initialize the native audio engine with a storage path for chunks and a callback for chunk readiness
void init_audio_engine(const char* storage_path, ChunkReadyCallback callback);

// Start capturing audio (triggered by UI)
void start_capture();

// Stop capturing audio
void stop_capture();

// Deinitialize the engine
void destroy_audio_engine();

#ifdef __cplusplus
}
#endif

#endif // AUDIO_ENGINE_H
