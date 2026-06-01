#ifndef AUDIO_ENGINE_H
#define AUDIO_ENGINE_H

#ifdef __cplusplus
extern "C" {
#endif

// Initialize the native audio engine with a storage path for chunks
void init_audio_engine(const char* storage_path);

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
