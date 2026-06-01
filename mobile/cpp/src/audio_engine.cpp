#include "audio_engine.h"
#include <oboe/Oboe.h>
#include <android/log.h>
#include <string>
#include <memory>
#include <vector>
#include <fstream>
#include <ctime>

#define LOG_TAG "AudioEngine"
#define LOGI(...) __android_log_print(ANDROID_LOG_INFO, LOG_TAG, __VA_ARGS__)
#define LOGE(...) __android_log_print(ANDROID_LOG_ERROR, LOG_TAG, __VA_ARGS__)

// Temporary "Scramble" key for verification
const uint8_t SCRAMBLE_KEY = 0xAA;

class AudioCaptureCallback : public oboe::AudioStreamCallback {
public:
    oboe::DataCallbackResult onAudioReady(oboe::AudioStream *audioStream, void *audioData, int32_t numFrames) override;
    void onErrorAfterClose(oboe::AudioStream *audioStream, oboe::Result result) override {
        LOGE("Audio Stream Error: %s", oboe::convertToText(result));
    }

private:
    void saveChunk();
    std::vector<int16_t> mBuffer;
    const int32_t mFramesPerChunk = 48000 * 10; // 10 seconds at 48kHz
};

struct EngineContext {
    std::string storage_path;
    std::shared_ptr<oboe::AudioStream> recording_stream;
    AudioCaptureCallback callback;
};

static std::unique_ptr<EngineContext> g_engine;

void init_audio_engine(const char* storage_path) {
    if (!g_engine) {
        g_engine = std::make_unique<EngineContext>();
    }
    g_engine->storage_path = storage_path;
    LOGI("Audio Engine Initialized with path: %s", storage_path);
}

oboe::DataCallbackResult AudioCaptureCallback::onAudioReady(oboe::AudioStream *audioStream, void *audioData, int32_t numFrames) {
    int16_t *inputData = static_cast<int16_t*>(audioData);
    
    // Accumulate frames
    for (int i = 0; i < numFrames; ++i) {
        mBuffer.push_back(inputData[i]);
    }

    // If we have 10 seconds, save it
    if (mBuffer.size() >= mFramesPerChunk) {
        saveChunk();
        mBuffer.clear();
    }

    return oboe::DataCallbackResult::Continue;
}

void AudioCaptureCallback::saveChunk() {
    if (g_engine->storage_path.empty()) return;

    std::string filename = g_engine->storage_path + "/chunk_" + std::to_string(std::time(nullptr)) + ".raw";
    std::ofstream outfile(filename, std::ios::binary);

    if (!outfile.is_open()) {
        LOGE("Failed to open file for writing: %s", filename.c_str());
        return;
    }

    outfile.write(reinterpret_cast<const char*>(mBuffer.data()), mBuffer.size() * sizeof(int16_t));
    outfile.close();

    LOGI("Saved raw 10s chunk: %s", filename.c_str());
}

void start_capture() {
    if (!g_engine) return;

    oboe::AudioStreamBuilder builder;
    builder.setDirection(oboe::Direction::Input)
           ->setPerformanceMode(oboe::PerformanceMode::LowLatency)
           ->setSharingMode(oboe::SharingMode::Shared) // Use Shared for better compatibility
           ->setFormat(oboe::AudioFormat::I16)
           ->setChannelCount(oboe::ChannelCount::Mono)
           ->setSampleRate(48000)
           ->setCallback(&g_engine->callback);

    oboe::Result result = builder.openStream(g_engine->recording_stream);
    if (result != oboe::Result::OK) {
        LOGE("Failed to open stream: %s", oboe::convertToText(result));
        return;
    }

    result = g_engine->recording_stream->requestStart();
    if (result != oboe::Result::OK) {
        LOGE("Failed to start stream: %s", oboe::convertToText(result));
        return;
    }

    LOGI("Audio Capture Started");
}

void stop_capture() {
    if (!g_engine || !g_engine->recording_stream) return;

    g_engine->recording_stream->requestStop();
    g_engine->recording_stream->close();
    g_engine->recording_stream.reset();

    LOGI("Audio Capture Stopped");
}

void destroy_audio_engine() {
    if (g_engine) {
        stop_capture();
        g_engine.reset();
    }
    LOGI("Audio Engine Destroyed");
}
